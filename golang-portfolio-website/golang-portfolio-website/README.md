
README.md - Project Website Portfolio Golang


# 🌐 Project Website Portfolio (Golang)

Aplikasi **Website Portfolio** berbasis **Golang (Go)** menggunakan **Server Side Rendering (SSR)**.

Website ini digunakan untuk menampilkan **profil**, **daftar project**, serta **form kontak**, dengan penyimpanan data menggunakan **PostgreSQL**.

---

## 👤 Informasi Umum

- **Nama Project**: project-app-portfolio-golang-rafli
- **Bahasa**: Go (Golang)
- **Database**: PostgreSQL
- **Web Framework**: chi
- **Template Engine**: html/template
- **Interface**: Website
- **Jenis Project**: Individu

---

## 📂 Struktur Folder

Struktur folder dirancang mengikuti konsep **Clean Architecture**  
(pemisahan handler, service, repository, dan konfigurasi).

 ```text 
project-app-portfolio-golang-rafli/
│
├── cmd/
│   └── server/
│
├── internal/
│   ├── config/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── router/
│   └── util/
│
├── templates/
│
├── static/
│   ├── css/
│   ├── js/
│   └── images/
│
├── migrations/
│
├── tests/
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
 ``` 

---

## 🧩 Fitur Aplikasi

### ✅ Halaman Website
- Halaman Home
- Halaman Daftar Project
- Halaman Detail Project
- Halaman Kontak

### ✅ Manajemen Project
- Menampilkan daftar project
- Menampilkan detail project
- Penyimpanan data project ke PostgreSQL

### ✅ Form Kontak
- Form kontak pengunjung
- Validasi input form
- Penyimpanan pesan ke database

### ✅ Fitur Teknis
- Clean Architecture (handler → service → repository)
- Server Side Rendering (SSR)
- Routing menggunakan chi
- Validasi data
- Error handling konsisten

---

## 🛠️ Persiapan & Instalasi

Pastikan **Go 1.20+** dan **PostgreSQL** sudah terinstal.

### 1️⃣ Clone Repository & Install Dependency

 ```bash 
git clone https://github.com/username/project-app-portfolio-golang-rafli.git
cd project-app-portfolio-golang-rafli
go mod tidy
 ``` 

---

### 2️⃣ Setup Database

Buat database baru:

 ```sql 
CREATE DATABASE portfolio;
 ``` 

Import struktur tabel:

 ```bash 
psql -d portfolio -f migrations/init.sql
 ``` 

---

### 3️⃣ Konfigurasi Database

Sesuaikan konfigurasi koneksi database di folder `internal/config`.

Contoh connection string:

 ```go 
postgres://postgres:password@localhost:5432/portfolio
 ``` 

---

## ▶️ Cara Menjalankan Aplikasi

### Menjalankan Langsung

 ```bash 
go run ./cmd/server
 ``` 

### Build Binary

 ```bash 
go build -o server ./cmd/server
./server
 ``` 

---

## 🌍 Akses Aplikasi

Buka browser dan akses:

http://localhost:8080

---

## 🧪 Testing

Menjalankan seluruh unit test:

 ```bash 
go test ./...
 ``` 

Menjalankan test dengan coverage:

 ```bash 
go test ./... -cover
 ``` 

---

## 📌 Catatan

- Project ini menggunakan **Server Side Rendering**
- Cocok digunakan sebagai **portfolio backend Golang**
- Struktur folder disiapkan untuk mudah dikembangkan ke REST API atau Admin Panel

---

🔥 **STATUS: SIAP LANGSUNG COMMIT KE GITHUB**  
😎 **PORTOFOLIO READY LEVEL MAHASISWA → JUNIOR BACKEND**

