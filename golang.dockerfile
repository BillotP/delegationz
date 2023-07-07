FROM docker.io/golang:1.20-alpine as builder
ARG cmd
RUN apk add git upx

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY pkg pkg
COPY cmd/$cmd/main.go .

ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X 'delegationz/pkg/api.VERSION=$version'" \ 
    -o /build/app -mod=mod \
    /src/main.go
RUN upx --best --lzma /build/app

FROM scratch
WORKDIR /run
COPY --from=builder /build/app /run/app
# Copy ca certificates for external https calls to work
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/run/app"]
