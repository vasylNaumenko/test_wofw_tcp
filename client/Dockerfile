FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM ubuntu:latest
LABEL Name=client
COPY --from=builder /app/main ./
ENTRYPOINT ["./main"]