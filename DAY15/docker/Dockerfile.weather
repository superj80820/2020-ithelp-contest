FROM golang:1.14.6-alpine

COPY ./weather /app

WORKDIR /app

ENTRYPOINT ["go", "run", "cmd/main.go"]