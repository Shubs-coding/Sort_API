# Dockerfile

# Build Stage
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

# Final Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
