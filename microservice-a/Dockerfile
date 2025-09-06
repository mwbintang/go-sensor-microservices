FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy proto files first
COPY ../proto/ ./proto/

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
EXPOSE 50051 8080
CMD ["./main"]