# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /go/src/github.com/previousnext/notify

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor/ vendor/
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

# Copy the go source
COPY main.go main.go

# Build
RUN go build -o bin/notify main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/previousnext/notify/bin/notify /usr/local/bin/notify
CMD ["notify"]
