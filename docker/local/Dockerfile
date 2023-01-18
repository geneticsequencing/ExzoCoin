FROM golang:1.18-alpine AS builder

WORKDIR /exzocoin

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o exzocoin main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /exzocoin

COPY --from=builder /exzocoin/exzocoin ./
COPY ./docker/local/exzocoin.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./exzocoin.sh"]
