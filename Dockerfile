# Use an official Go runtime as a parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port (if your Go application listens on a specific port)
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]
