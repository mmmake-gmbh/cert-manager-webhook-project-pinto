FROM golang:1.15.2-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /go/src/gitlab.com/whizus/cert-manager-webhook-pinto

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -mod=readonly -a -o cert-manager-webhook-pinto main.go

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/gitlab.com/whizus/cert-manager-webhook-pinto/cert-manager-webhook-pinto .
ENTRYPOINT ["/cert-manager-webhook-pinto"]
