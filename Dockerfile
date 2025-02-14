# Build stage: Use an official Golang image to compile the binary.
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application. The binary will be named "main"
RUN go build -o main .

# Final stage: Use a lightweight image to run the binary
FROM alpine:latest

# Install CA certificates (in case your app makes HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 9808

# Command to run the executable
CMD ["./main"]
