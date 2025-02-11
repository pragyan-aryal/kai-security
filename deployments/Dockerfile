# Build stage
FROM golang:1.23.3 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/kai-security ./cmd/server

# Final stage
FROM alpine:latest

# ENV PORT=$HTTP_PORT

# Set working directory
WORKDIR /app/
# Copy the binary from the builder stage
COPY --from=builder /app/bin/kai-security .

# Copy the .env file
COPY .env .

# Expose the HTTP port
EXPOSE 8080
# EXPOSE $PORT

# Run the application
CMD ["./kai-security"]