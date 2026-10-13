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

The `MQ tools` workflow tests the tools, builds the image, and smoke-tests the
runtime on pull requests. Publishing uses `GITHUB_TOKEN` with `packages: write`.
A published tag matching `vX.Y.Z-mq.N`, or an explicit manual publishing run for
that existing tag, publishes `ghcr.io/injectivelabs/sdk-go-mq:<tag>` for amd64 and
arm64. Manual runs against branches build and test without publishing. No tag or
image is created by committing these files. The proposed next release is
`v1.64.0-mq.8`; it must be released before a deployment uses that image tag.

The workflow also needs a repository or organization Actions secret named
`GH_TOKEN` with read access to the private InjectiveLabs Go dependency repositories,
including `InjectiveLabs/cosmos-sdk`. This is a source-read credential, separate
from the workflow's `GITHUB_TOKEN` used to publish GHCR packages. Configure a
fine-grained token with repository contents read permission for those repositories
(or the organization's equivalent approved credential). CI fails explicitly if
this secret is unavailable; fork pull requests do not receive Actions secrets.
The existing general test and lint workflows need their own private-module
authentication if they do not already have it.

`scripts/with-private-git.sh` applies authentication only to its child command;
it does not write credentials into Git config. Before building an image,
`scripts/prepare-mq-go-modules.sh` stages the SDK's private module versions in
gitignored `.modcache/`. Docker reads that file proxy and receives no token.
Locally, the scripts can use an existing Git credential helper without `GH_TOKEN`.

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
for both `GET /status` and `POST /request`. Neither the Kafka connection nor the
stream server has native TLS/SASL configuration; use controlled internal access.

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
./scripts/with-private-git.sh go test -race ./client/mq/...
./scripts/prepare-mq-go-modules.sh
docker build -f Dockerfile.mq -t sdk-go-mq:test .
docker run --rm sdk-go-mq:test /usr/local/bin/mq-detector -help
docker run --rm sdk-go-mq:test /usr/local/bin/mq-stream-server -help
```

Before deployment, also validate broker and Redis outages, publisher retention
exhaustion, container restarts during an existing gap, and block payloads larger
than 4 MiB against the real pipeline. Unit tests do not replace that acceptance
test with Kafka, the publisher, and Connect running together.
