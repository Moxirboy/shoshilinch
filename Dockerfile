FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /app


COPY . .
ENTRYPOINT [ "go","run","cmd/main.go" ]