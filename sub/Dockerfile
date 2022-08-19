
# syntax = docker/dockerfile:experimental
FROM golang:1.18-buster as builder

RUN apt update && \
    apt -y install \
        apt-transport-https \
        ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux go build -v -o bin/sub main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/sub /usr/local/bin/sub

ENTRYPOINT [ "/usr/local/bin/sub" ]
