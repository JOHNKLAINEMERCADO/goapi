# Stage 1: Build the Go executable binary
FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    golang-go \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
# Set working directory inside the container
WORKDIR /app

# Copy dependency tracking files first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project source code
COPY . .

# Build a static binary optimized for minimal linux containers
RUN go build -o goapi ./cmd/api

# Stage 2: Create a secure, lightweight runtime environment

# Expose your application port
EXPOSE 8000

# Fire up the compiled Go application executable
CMD ["./goapi"]
