# Use the official Golang image to create a build artifact.
FROM golang:1.19 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./api/go.mod ./api/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY ./api .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o rest_api ./rest_api

# Start a new stage from scratch
FROM alpine:latest

# Install CA certificates for Alpine
RUN apk --no-cache add ca-certificates

# Set the working directory in the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/rest_api .

# Expose port 3001 to the outside world
EXPOSE 3001

# Command to run the executable
CMD ["./rest_api"]
