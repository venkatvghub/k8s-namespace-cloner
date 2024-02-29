 # Use the official golang image to create a build environment
FROM golang:1.21 as builder

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Use a minimal alpine image as the base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main /app/main

# Define the binary as the entrypoint of the container
ENTRYPOINT ["/app/main"]

# Provide a default argument
CMD ["--incluster=true"]
