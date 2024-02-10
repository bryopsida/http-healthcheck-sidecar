FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/http-healthcheck-sidecar/
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/http-healthcheck-sidecar
FROM scratch
USER 10001
WORKDIR /app
COPY --from=builder /go/bin/http-healthcheck-sidecar /app/http-healthcheck-sidecar
EXPOSE 3000
ENTRYPOINT ["/app/http-healthcheck-sidecar"]