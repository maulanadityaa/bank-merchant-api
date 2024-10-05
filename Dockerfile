# Use the official Golang image to build the Go application
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache tzdata

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal base image to reduce the size
FROM alpine:latest

# Set environment variables for the final stage
ENV APP_ENV=production \
    GIN_MODE=release

# Set working directory in the minimal image
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Debugging: Check if the 'main' file exists
RUN ls -l ./main

# Ensure the static directory exists (if needed)
RUN mkdir -p /app/static

# Copy the .env file
COPY .env /app

# Copy static files
COPY --from=builder /app/static ./static

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
