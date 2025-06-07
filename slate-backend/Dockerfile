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
    sqlite-dev \
    curl

# Install Air with explicit path and make it executable
RUN curl -sSf https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh && \
    mv ./bin/air /usr/local/bin/ && \
    chmod +x /usr/local/bin/air

# Copy air config first
COPY air.toml ./

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Create tmp directory for Air
RUN mkdir -p tmp

EXPOSE 8080

# Use absolute path to Air
CMD ["/usr/local/bin/air", "-c", "air.toml"]