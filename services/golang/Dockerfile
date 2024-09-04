FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY otel.go .
COPY main.go .

RUN go mod tidy
RUN go build -o main

EXPOSE 8081