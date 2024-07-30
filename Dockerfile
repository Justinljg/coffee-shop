# Use the official Go image as a build environment
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (if needed, adjust to your application's port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
