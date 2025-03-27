FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the app
RUN go build -o main ./cmd/harbor

# Expose the port (Railway injects this via PORT env var)
EXPOSE 8080

# Start the app (read PORT from env)
CMD ["sh", "-c", "./main"]
