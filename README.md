# Aplikasi Pembayaran User

Aplikasi ini adalah API sederhana yang dibangun menggunakan Golang dan Gin Framework untuk mengelola pengguna dan transaksi pembayaran antar pengguna.

## Fitur

- Registrasi dan login pengguna
- Manajemen pengguna (CRUD)
- Riwayat pembayaran
- Pembayaran antar pengguna
- Logout dan manajemen sesi dengan JWT

## Menjalankan Aplikasi

1. **Clone repositori ini**:

```bash
  git clone https://github.com/username/repo-name.git
  cd repo-name
```

2. **Membuat file `.env`**:

Buat file `.env` di root proyek Anda dan tambahkan variabel berikut:

```plaintext
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gin_project_db
JWT_SECRET=akudankamuadalahmustahil
```

Sesuaikan password dan username Anda.
Jangan lupa membuat database MySQL yang sesuai.

3.  **Jalankan aplikasi**:

```bash
go run main.go
```

Aplikasi akan tersedia di `http://localhost:8080`

## Dokumentasi Endpoint

Berikut adalah daftar endpoint yang tersedia dalam aplikasi:

### Public Routes

- **POST** `/login`

  - **Deskripsi**: Melakukan login pengguna.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    - Status 200: Login berhasil dan mengembalikan token JWT.
    - Status 401: Username atau password tidak valid.

- **POST** `/register`
  - **Deskripsi**: Mendaftar pengguna baru.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    - Status 201: Pengguna berhasil dibuat.
    - Status 400: Permintaan tidak valid (misalnya, username sudah ada).

### Protected Routes (Memerlukan Token JWT)

- **GET** `/users`

  - **Deskripsi**: Mengambil daftar semua pengguna.
  - **Response**:
    - Status 200: Berhasil mengembalikan daftar pengguna.
    - Status 401: Token tidak valid atau tidak ada.

- **POST** `/users`

  - **Deskripsi**: Membuat pengguna baru (admin hanya).
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    - Status 201: Pengguna berhasil dibuat.
    - Status 401: Token tidak valid atau tidak ada.

- **GET** `/history`

  - **Deskripsi**: Mengambil riwayat tindakan pengguna.
  - **Response**:
    - Status 200: Berhasil mengembalikan riwayat.
    - Status 401: Token tidak valid atau tidak ada.

- **POST** `/logout`
  - **Deskripsi**: Mengeluarkan pengguna dari sesi.
  - **Response**:
    - Status 200: Logout berhasil.
    - Status 401: Token tidak valid atau tidak ada.

### Payment Routes

- **POST** `/payments`

  - **Deskripsi**: Membuat pembayaran antar pengguna.
  - **Request Body**:
    ```json
    {
      "receiver_id": "uint", // ID pengguna penerima
      "amount": "float64" // Jumlah yang akan dibayarkan
    }
    ```
  - **Response**:
    - Status 201: Pembayaran berhasil dibuat.
    - Status 400: Permintaan tidak valid (misalnya, saldo tidak cukup).
    - Status 401: Token tidak valid atau tidak ada.

- **GET** `/payments`
  - **Deskripsi**: Mengambil riwayat pembayaran pengguna.
  - **Response**:
    - Status 200: Berhasil mengembalikan riwayat pembayaran.
    - Status 401: Token tidak valid atau tidak ada.
