# Use the official Golang image as a base image
FROM golang:1.24

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server/main.go

# Expose the gRPC and Prometheus metrics ports
EXPOSE 50051 2112

# Command to run the application
CMD ["/app/main"]