# syntax=docker/dockerfile:1
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server main.go

FROM ubuntu:24.04

# 必要な環境変数を設定
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    software-properties-common apt-transport-https ca-certificates curl && \
    add-apt-repository ppa:inkscape.dev/stable && \
    apt-get update && \
    apt-get install -y --no-install-recommends inkscape && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /bin/server

EXPOSE 8080
ENTRYPOINT ["/bin/server"]