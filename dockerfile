# Gunakan image Go sebagai basis
FROM golang:1.16

# Set direktori kerja
WORKDIR /app

# Salin file go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependensi
RUN go mod download

# Salin semua file sumber ke dalam container
COPY . .

# Bangun aplikasi
RUN go build -o main .

# Tentukan perintah untuk menjalankan aplikasi
CMD ["./main"]