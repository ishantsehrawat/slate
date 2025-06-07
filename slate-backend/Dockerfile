FROM golang:1.21-alpine

WORKDIR /app

ENV GOTOOLCHAIN=auto

# Add necessary build tools
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build with explicit output path and verify
RUN go build -o main ./cmd/server && \
    ls -la main && \
    chmod +x main

# Expose port
EXPOSE 8080

# Use relative path for CMD
CMD ["./main"]