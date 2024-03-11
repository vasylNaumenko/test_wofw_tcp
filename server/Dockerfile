FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum ./cmd/config.yaml ./
RUN go mod download
COPY . .
RUN go build -o ./main ./cmd/.

FROM ubuntu:latest
LABEL Name=server
COPY --from=builder /app/main /app/config.yaml ./

EXPOSE 9090
ENTRYPOINT ["./main"]