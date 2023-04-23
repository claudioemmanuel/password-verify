# Use an official Golang image as the base image
FROM golang:latest

# Set the working directory in the container to /src/app
WORKDIR /src/app

# Copy the Go source code
COPY . .

# Download and install any Go dependencies
RUN go get -d -v ./...

# Download and install the Gin framework
RUN go get github.com/gin-gonic/gin

# Download and install the Testify for testing
RUN go get  github.com/stretchr/testify/assert

# Build the Go binary
RUN go build -o main .

# Expose port 8080 for the Go application
EXPOSE 8080

# Start the Go application
CMD ["./main"]