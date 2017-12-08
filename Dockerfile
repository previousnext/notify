FROM golang:1.9
ADD workspace /go
RUN go get github.com/mitchellh/gox
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/notify_linux_amd64 /usr/local/bin/notify
CMD ["notify"]