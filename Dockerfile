# syntax=docker/dockerfile:1
FROM golang:1.25.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
      -ldflags="-X github.com/y-yu/kindle-clock-go/domain/build.gitCommitHash=$GIT_COMMIT_HASH -s -w" \
      -trimpath \
      -o /app/server main.go

FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* &&  \
    update-ca-certificates

ENV DOSIS_FONT_PATH=/etc/Dosis.ttf
ENV ROBOTO_SLAB_FONT_PATH=/etc/RobotoSlab.ttf
RUN curl -sSL -o "$ROBOTO_SLAB_FONT_PATH" "https://raw.githubusercontent.com/google/fonts/refs/heads/main/apache/robotoslab/RobotoSlab%5Bwght%5D.ttf" && \
    curl -sSL -o "$DOSIS_FONT_PATH" "https://raw.githubusercontent.com/google/fonts/refs/heads/main/ofl/dosis/Dosis%5Bwght%5D.ttf"

COPY --from=builder /app/server /bin/server
COPY --from=builder /app/etc/weather_icon /etc/weather_icon

EXPOSE 8080
ENTRYPOINT ["/bin/server"]