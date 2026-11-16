# MQ tools

The MQ tools consume the protobuf block records produced by Injective's MQ
publisher. The stream server provides a gRPC view of a Kafka topic; the detector
requests locally retained blocks when the `latest` pipeline stops progressing.

## Runtime image

`Dockerfile.mq` builds a dedicated image with two static Go binaries:

- `/usr/local/bin/mq-stream-server`
- `/usr/local/bin/mq-detector`

The server and detector build with `CGO_ENABLED=0` and need no wasmvm libraries.
The interactive stream client imports the full chain decoder and is intentionally
excluded from the runtime image. Run it from this repository with `go run`.
The runtime contains `/bin/sh` for a Kubernetes sidecar supervisor and runs as
UID/GID 10001. Mount a writable state directory with suitable ownership or a pod
`fsGroup`; `/var/lib/mq` is writable in the image.

The repository's existing CI workflows remain unchanged. There is no additional
MQ workflow or automatic image publication. Build and verify the image with the
commands under Verification below. Every Go dependency is public, so the Docker
build fetches modules from the default Go proxy and needs no credentials. The
final image contains the compiled tools, including dependency code, but no
module source archives.

Image publication is a separate release step through the approved release
process. The proposed next image is
`ghcr.io/injectivelabs/sdk-go-mq:v1.64.0-mq.8`; publish and verify its amd64 and
arm64 variants before a deployment uses that tag. Committing these files does
not create a release tag or publish an image.

## Detector

```sh
mq-detector \
  -kafka-brokers=broker-0:9092,broker-1:9092,broker-2:9092 \
  -consumer-id=quants-publisher-0-detector \
  -raw-topic=unrefined_block_results \
  -latest-topic=latest \
  -full-nodes=http://127.0.0.1:8080 \
  -state-file=/var/lib/mq/detector.json \
  -message-timeout=30s \
  -retry-backoff=5s
```

Provide a **stable, unique consumer ID for every detector**. Sharing an ID splits
Kafka partitions among detectors, leaving each with an incomplete view. Keep the
state file across container and pod restarts. Put it outside the chain `data/`
directory if it must survive `unsafe-reset-all`. Do not share a file between
detectors or reuse it for another consumer/topic pair. The detector rejects
mismatched state instead of guessing.
Use a single partition for `latest`, preserving global block-height order.

The publisher control API binds to loopback, so a detector belongs in the same
pod as its publishing node. `-control-token` supplies the optional bearer token
for both `GET /status` and `POST /request`. The publisher control API uses its own authentication. Configure the detector's
Kafka credentials using the environment variables below.

Recovery requires a publisher whose `/status` includes `earliest_height` and
`last_height` (the original `v1.20.4-bl-data.1` lacks `earliest_height`). Deploy the
core recovery follow-up release before enabling the detector. With no saved
watermark, the detector bootstraps by replaying the node's earliest retained
height, never by iterating from height zero. With saved state it requests
`latest_height + 1`. A node behind that height is retried after it catches up.
A request older than retained history is reported explicitly: recovery requires
another publisher with the missing history or operator intervention; the detector
does not reset Redis cursors or silently skip a gap.

Latest progress is saved atomically and synced before its Kafka record is marked
for commit. Restarts can therefore recover a preexisting gap even when no new
latest record arrives. A real timer checks progress when Kafka is idle. Kafka
fetch errors and unavailable control endpoints are retried with bounded delay;
only one control request worker runs at a time. Raw traffic does not create an
unbounded height map. Replays may be repeated until the latest cursor advances;
publisher and connector deduplication remain necessary. Local state write errors
stop the process without committing beyond the saved watermark.

Configure sidecar supervision for fatal configuration/storage errors separately
from chain readiness. A broker outage must not become a chain health probe.

## Kafka TLS and SCRAM authentication

The detector and stream server accept the same three environment variables.

| Variable | Value |
|---|---|
| `MQ_KAFKA_USERNAME` | The Kafka principal assigned to this process |
| `MQ_KAFKA_PASSWORD` | Its SCRAM password |
| `MQ_KAFKA_CA_FILE` | A mounted PEM CA bundle for this Kafka cluster |

Set all three to enable TLS with SCRAM-SHA-512. The clients verify the broker
certificate and hostname. Inject credentials through the deployment's Secret
references and mount the public CA bundle. A partial configuration, unreadable CA
file, or invalid PEM stops the process before it starts serving. Every stream
request uses the server's configured Kafka principal, with the requested topic
and consumer group checked by the broker's ACLs. Restart the process after
rotating the credentials or CA bundle.

With all three variables empty, both tools keep their existing plaintext Kafka
connection behavior. Kafka authentication does not authenticate callers of the
gRPC stream listener: any caller that can reach it can use the server principal's
permitted topics and consumer groups. Restrict listener ingress to trusted
workloads or an authenticated/allowlisted gateway, and grant that principal only
the intended topics and groups. The publisher control API has separate access
controls.

### Kafka integration test

`TestKafkaStreamTLSAndACLs` exercises the real stream handler over gRPC against a
disposable Kafka broker with TLS, SCRAM-SHA-512 and enforced ACLs. Create the
`latest` and `forbidden` topics. Give the reader principal Read/Describe only on
the literal `latest` topic and literal `mm-rfq-inj1jvrnhahd` group, and give a
separate writer principal Write/Describe on `latest`. Do not grant either test
principal superuser or wildcard access.

Set `MQ_TEST_KAFKA_BROKERS`, the three `MQ_KAFKA_*` variables above for the reader,
and `MQ_TEST_WRITER_USERNAME` / `MQ_TEST_WRITER_PASSWORD` for the writer. Supply
passwords through the test environment's secret mechanism, not command arguments.
Then run:

```sh
go test ./client/mq/stream/server -run '^TestKafkaStreamTLSAndACLs$' -count=1 -v
```

The test writes a marked record and advances the allowed consumer group. Use an
isolated fixture, never a production broker. It verifies successful delivery and
broker rejection of a forbidden topic, forbidden consumer group, and incorrect
SCRAM password. Without `MQ_TEST_KAFKA_BROKERS`, this test is skipped. These checks
prove the broker ACL boundary; gateway allowlisting and NetworkPolicy enforcement
must also be verified in the deployment before admitting external clients.

To test the detector's authenticated client and durable watermark, also create
`unrefined_block_results`. Set `MQ_KAFKA_USERNAME` / `MQ_KAFKA_PASSWORD` to a
detector principal with Read/Describe on both input topics and the literal group
`quants-gap-integration`, grant the writer Write/Describe on both topics, and run:

```sh
go test ./client/mq/detector -run '^TestKafkaDetectorTLSWatermark$' -count=1 -v
```

This writes marked raw and latest records and verifies that the detector's real
startup path consumes both over TLS/SCRAM, persists the height in a temporary
state file, and commits offsets past both records.
It does not contact a live publisher control API or request block replay.

## Stream server and client

```sh
mq-stream-server -listen-address=0.0.0.0:9988 -kafka-brokers=broker-0:9092

go run ./client/mq/stream/client \
  -address=localhost:9988 -consumer-id=my-stable-consumer -topic=guaranteed \
  -max-recv-msg-bytes=34603008 -format=minimal
```

Consumers must use an explicit stable `-consumer-id`: it names a Kafka consumer
group. Reusing the ID resumes committed offsets; sharing it distributes records
across consumers. The server starts new groups at the earliest retained offset.
The CLI accepts 33 MiB responses by default, enough for a 32 MiB block plus
framing headroom; custom clients must set an equivalent gRPC receive limit.
The server has TCP transport and reflection, but no gRPC health service; use a TCP
probe. It forwards any requested topic and has no built-in client authentication.
The `guaranteed` topic describes chain ordering and hash selection, not
exactly-once client processing or an application acknowledgement protocol.

## Verification

```sh
go test -race ./client/mq/...
docker build -f Dockerfile.mq -t sdk-go-mq:test .
docker run --rm sdk-go-mq:test /usr/local/bin/mq-detector -help
docker run --rm sdk-go-mq:test /usr/local/bin/mq-stream-server -help
```

Before deployment, also validate broker and Redis outages, publisher retention
exhaustion, container restarts during an existing gap, and block payloads larger
than 4 MiB against the real pipeline. Unit tests do not replace that acceptance
test with Kafka, the publisher, and Connect running together.
