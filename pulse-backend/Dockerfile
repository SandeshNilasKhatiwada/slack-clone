# STAGE 1: Build the Go binary
FROM golang:alpine AS builder
WORKDIR /app
# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code and build the application
COPY . .
RUN go build -o main .

# STAGE 2: Create the minimal production image
FROM alpine:latest
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Expose the port Fiber runs on
EXPOSE 8080
# Run the binary
CMD ["./main"]