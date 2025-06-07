FROM golang:1.21-alpine

WORKDIR /app

ENV GOTOOLCHAIN=auto

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /app/main ./cmd/server

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]