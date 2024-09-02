FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go mod tidy

EXPOSE 8081