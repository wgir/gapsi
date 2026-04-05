# Build stage
FROM golang:1.26 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# We use -ldflags="-s -w" to reduce the binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o todo-api ./cmd/api/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates and tzdata for SSL/TLS and timezones
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/todo-api .
# Copy migrations if they are needed at runtime
COPY --from=builder /app/migrations ./migrations
# Copy .env file if needed (optional, docker-compose uses its own)
# COPY --from=builder /app/.env .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./todo-api"]
