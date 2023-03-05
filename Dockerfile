# Use a Golang base image
FROM golang:1.19.5

# Set the working directory to /app
WORKDIR /app

# Copy the Go module files into the container
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the application files into the container
COPY . .

# Build the Go application
RUN go build -o app

# Expose port 8080 for the application to listen on
EXPOSE 8080

# Run the Go application
CMD ["./app"]docker build -t