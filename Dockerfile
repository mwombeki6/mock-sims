# Multi-stage build for smaller final image
FROM golang:alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mock-sims cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mock-sims-seed cmd/seed/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /build/mock-sims .
COPY --from=builder /build/mock-sims-seed .

# Copy .env file (optional, can be overridden by environment variables)
COPY .env.example .env

# Expose port
EXPOSE 8000

# Run the application
CMD ["./mock-sims"]
