# -----------------------------
# Build Stage
# -----------------------------
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# -----------------------------
# Runtime Stage
# -----------------------------
FROM debian:stable-slim

WORKDIR /app

# Install certs for HTTPS requests if needed later
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy built binary
COPY --from=builder /app/server .

# Expose service port
EXPOSE 8080

# Run binary
CMD ["./server"]