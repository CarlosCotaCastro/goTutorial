# Dockerfile for Go code execution sandbox
FROM golang:1.21-alpine

# Install necessary packages
RUN apk add --no-cache \
    gcc \
    musl-dev \
    git \
    sqlite

# Create a non-root user for security
RUN adduser -D -s /bin/sh gouser

# Set working directory
WORKDIR /app

# Copy the Go code execution script
COPY execute.go /app/

# Build the execution helper
RUN go build -o execute execute.go

# Switch to non-root user
USER gouser

# Set up the execution environment
RUN mkdir -p /app/code && \
    mkdir -p /app/data && \
    chown -R gouser:gouser /app

# Default command
CMD ["./execute"]
