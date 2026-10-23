# MQ container releases

The `Release MQ container` workflow builds both `mq-detector` and
`mq-stream-server` from an SDK MQ release tag and publishes one `linux/amd64`
image to `public.ecr.aws/l9h3g6c6/sdk-go:<release-tag>`. It uses the tagged
[`Dockerfile.mq`](https://github.com/InjectiveLabs/sdk-go/blob/v1.64.0-mq.8/Dockerfile.mq),
Go version, dependency versions, and module-preparation script.
It does not publish a `latest` tag or change the SDK's existing test/lint jobs.

## Prerequisites

1. Apply [injective-iac PR #516](https://github.com/InjectiveLabs/injective-iac/pull/516),
   which creates the public `sdk-go` ECR
   repository and the `sdk-go` IAM publishing role in AWS account `981432137740`.
   ECR Public authentication uses `us-east-1`.
2. Configure the release-only `MQ_RELEASE_GH_TOKEN` repository secret with a
   read-only credential for the private InjectiveLabs Go modules used by the
   release. The public SDK repository's automatic `GITHUB_TOKEN` cannot read
   those other repositories. For `v1.64.0-mq.8`, the private dependencies are
   `InjectiveLabs/block-sdk`, `cometbft`, `cosmos-sdk`, `go-ethereum`, `ibc-apps`,
   `ibc-go`, and `wasmd`. A fine-grained token should select those repositories
   with **Contents: read-only** (and the default read-only metadata permission).
   Recheck the tagged module graph when dependencies change. Do not add this
   credential to the ordinary test or pull request workflows.
3. Merge this workflow onto `master` before manually releasing an existing tag.
   Include the workflow in the branch used to create future MQ tags if automatic
   tag-push releases are wanted.

The AWS role uses GitHub OIDC, so no AWS access-key secret is required. Its trust
is limited to this SDK repository's `master` branch and MQ version tags; pull
request and fork subjects cannot assume it. Maintainers who can dispatch this
workflow or create MQ release tags are trusted to run tagged build code. Release
tag protections should follow the repository's normal release controls.

## Publish an existing release

Existing tags, including `v1.64.0-mq.8`, predate this workflow. Adding a workflow
to `master` does not retroactively run workflows for those tags. Once the
prerequisites are complete, run:

```sh
gh workflow run release-mq.yml \
  --repo InjectiveLabs/sdk-go \
  --ref master \
  -f release_tag=v1.64.0-mq.8
```

The workflow only accepts `v<major>.<minor>.<patch>-mq.<number>` tags. It resolves
the requested tag to an exact commit before checking out the release source.
Manual runs from branches other than `master` are skipped. Future matching tag
pushes run automatically if the tagged commit contains this workflow.

Publication uses the supplied version tag. Do not move published release tags or
rerun publication to replace a released image; publish a new MQ version instead.
For deployment, record the image digest from the push output and pin it where
immutable image references are required.

## Build and credential boundaries

The build job has read-only GitHub permissions and no AWS role. Only dependency
preparation receives `MQ_RELEASE_GH_TOKEN`. Go caching is disabled, and the
Docker build does not export build caches or intermediate layers. The token is
not a Docker build argument and is not copied into the build context.

The tagged multi-stage Dockerfile copies only the two binaries into its final
runtime stage. The workflow runs the MQ package tests, builds the image, checks
that it runs as UID/GID `10001:10001`, and invokes both binaries with `--help`
without network access. It also checks that the runtime filesystem does not
contain the source or Go module cache directories.

Only that final runtime image and its SHA-256 checksum are uploaded as a
short-lived workflow artifact. Treat the image and workflow logs as public:
private module archives, Git credentials, and intermediate layers must never be
added to artifacts or external build caches. The separate publishing job checks
the artifact checksum, assumes the narrowly scoped AWS role, and pushes the
versioned image. It does not check out or execute the release source.
