FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o translator ./cmd/AI-API

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/translator .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
CMD ["./translator"]