# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY main.go .

# Initialize Go module automatically
RUN go mod init backend || true
RUN go build -o backend .

# Final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8000

CMD ["./backend"]

