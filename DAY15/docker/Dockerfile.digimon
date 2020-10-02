FROM golang:1.14.6-alpine

COPY ./go-server /app

WORKDIR /app

ENTRYPOINT ["go", "run", "cmd/main.go"]