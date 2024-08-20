# Gunakan base image Golang
FROM golang:1.22

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependensi
RUN go mod download

# Compile aplikasi Go
RUN go build -o clothes-inventory-api ./cmd/http

# Perintah yang dijalankan ketika container dijalankan
CMD ["./clothes-inventory-api"]
