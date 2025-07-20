# --- Build Stage ---
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

# Copy go.mod and download dependencies first for caching
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Enable static build to reduce dependencies
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app/cmd/api

# Build the binary
RUN go build -ldflags="-s -w" -o url-service .

# --- Final Stage ---
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy the static binary from the builder stage
COPY --from=builder /app/cmd/api/url-service .

# Expose ports for 10 auth-service connections (starting at 30500)
EXPOSE 8080

# Run the auth-service binary
CMD ["./url-service"]