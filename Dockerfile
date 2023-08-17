FROM --platform=$BUILDPLATFORM golang:1.21 AS builder

WORKDIR /workspace
COPY . /workspace

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o annict-epgstation-connector ./cmd/annict-epgstation-connector

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/annict-epgstation-connector .
USER 65532:65532

ENTRYPOINT ["/annict-epgstation-connector"]
