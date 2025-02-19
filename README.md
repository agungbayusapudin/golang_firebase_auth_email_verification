markdown
Copy
# Golang Firebase Auth Email Verification

Repository ini berisi contoh implementasi verifikasi email menggunakan Firebase Authentication dengan bahasa pemrograman Go (Golang). Projek ini dapat digunakan sebagai referensi untuk mengintegrasikan Firebase Auth ke dalam aplikasi backend Golang, khususnya untuk fitur verifikasi email.

## Fitur

- Registrasi pengguna dengan email dan password.
- Mengirim email verifikasi ke pengguna yang baru terdaftar.
- Memverifikasi email pengguna menggunakan link verifikasi yang dikirimkan.
- Contoh endpoint API untuk registrasi dan verifikasi email.

## Prasyarat

Sebelum menjalankan projek ini, pastikan Anda telah memenuhi persyaratan berikut:

1. **Go** - Pastikan Go sudah terinstal di sistem Anda. Anda dapat mengunduhnya dari [situs resmi Go](https://golang.org/dl/).
2. **Firebase Project** - Buat projek Firebase di [Firebase Console](https://console.firebase.google.com/).
3. **Firebase Admin SDK** - Unduh file kredensial Firebase Admin SDK (`serviceAccountKey.json`) dari Firebase Console.
4. **Environment Variables** - Simpan file kredensial Firebase dan konfigurasikan environment variable.

## Instalasi

1. Clone repository ini ke lokal Anda:
   ```bash
   git clone https://github.com/username/golang_firebase_auth_email_verification.git
   cd golang_firebase_auth_email_verification
Install dependencies menggunakan Go mod:
```
Copy
go mod download
```
Simpan file serviceAccountKey.json dari Firebase Console ke direktori projek.
Buat file .env di root direktori dan tambahkan konfigurasi berikut:
```
env
Copy
GOOGLE_APPLICATION_CREDENTIALS=./serviceAccountKey.json
```
Jalankan aplikasi:
```
bash
Copy
go run main.go
```
Struktur Projek

Copy
golang_firebase_auth_email_verification/
├── handlers/           # Berisi handler untuk endpoint API
├── models/             # Berisi model data
├── utils/              # Berisi utility functions
├── main.go             # Entry point aplikasi
├── go.mod              # File dependensi Go
├── go.sum              # File checksum dependensi Go
├── README.md           # Dokumentasi projek
└── .env                # File environment variable
Endpoint API

1. Registrasi Pengguna

Endpoint: POST /register

Request Body:
```
json
Copy
{
  "email": "user@example.com",
  "password": "password123"
}
```
Response:
```
json
Copy
{
  "message": "User registered successfully. Please check your email for verification.",
  "user_id": "abc123"
}
```
2. Verifikasi Email
```
Endpoint: POST /verify-email
```
Request Body:
```
json
Copy
{
  "token": "email-verification-token"
}
```
Response:
```
json
Copy
{
  "message": "Email verified successfully."
}
```
Contoh Penggunaan

Registrasi Pengguna:
```
Copy
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "password123"}'
```

Jika Anda ingin berkontribusi pada projek ini, silakan buka issue atau pull request. Semua kontribusi sangat diterima!
