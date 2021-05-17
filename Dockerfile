FROM golang:1.16.4-alpine as builder

RUN apk update && \
    apk add --no-cache git ca-certificates && \
    update-ca-certificates

WORKDIR /go/src/gitlab.com/whizus/cert-manager-webhook-pinto

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -mod=readonly -a -o cert-manager-webhook-pinto main.go

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/gitlab.com/whizus/cert-manager-webhook-pinto/cert-manager-webhook-pinto .
ENTRYPOINT ["/cert-manager-webhook-pinto"]
