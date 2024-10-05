# Gunakan image Golang sebagai base image untuk build
FROM golang:1.22.5 AS builder

# Set environment variable untuk disable CGO (agar binary statik dan bisa dijalankan di Alpine)
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Install dependensi
RUN go mod download

# Copy seluruh kode sumber
COPY . .

# Build aplikasi
RUN go build -o main .

# Gunakan image alpine sebagai base image untuk menjalankan aplikasi
FROM alpine:latest

# Install ca-certificates untuk HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app/

# Copy binary dari tahap builder
COPY --from=builder /app/main .

# Beri izin eksekusi untuk binary
RUN chmod +x ./main

# Copy .env file
COPY .env .env

# Expose port
EXPOSE 8080

# Command untuk menjalankan aplikasi
CMD ["./main"]
