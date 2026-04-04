# --- Build Stage ---
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies to leverage Docker's cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application binary.
# CGO_ENABLED=0 is often used to statically link the binary, making it portable.
RUN CGO_ENABLED=0 go build -o /go-api-app .

# --- Run Stage ---
# Use a minimal base image, like scratch or alpine, for the final image
FROM alpine:latest

# Set the working directory to root
WORKDIR /root/

# Copy the built binary from the 'builder' stage into the new image
COPY --from=builder /go-api-app .

# Expose the port your Go API listens on (e.g., 8080)
EXPOSE 3001

# Command to run the executable when the container starts
CMD ["./go-api-app"]
