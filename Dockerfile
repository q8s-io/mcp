## Build the mcp-server binary
FROM golang:1.13-alpine as builder
WORKDIR /workspace

# Run this with docker build --build_arg $(go env GOPROXY) to override the goproxy
ARG goproxy=https://goproxy.cn
ENV GOPROXY=$goproxy

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# Cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the sources
COPY ./ ./

# Build
ARG ARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} \
    go build -a -ldflags '-extldflags "-static"' \
    -o mcp-server ./cmd/mcp-server

## Copy the mcp-server into a thin image
FROM alpine:3.11
WORKDIR /
COPY --from=builder /workspace/mcp-server .
CMD ["/mcp-server"]
