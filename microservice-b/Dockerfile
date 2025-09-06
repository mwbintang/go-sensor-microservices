FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Create final minimal image
FROM alpine:3.18

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy built binary
COPY --from=builder /app/main .

EXPOSE 50051 8080

CMD ["./main"]