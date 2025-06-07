FROM golang:1.21-alpine

WORKDIR /app

ENV GOTOOLCHAIN=auto
ENV CGO_ENABLED=1

# Add necessary build tools
RUN apk add --no-cache \
    git \
    gcc \
    musl-dev \
    sqlite \
    sqlite-dev 

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