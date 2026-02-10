# Contact Management API

Mini project **Contact Management** ini adalah RESTful API sederhana untuk mengelola data kontak (CRUD) yang dibangun menggunakan **Golang**. Project ini dirancang untuk latihan backend engineering dengan struktur yang rapi, dependency injection, validasi, autentikasi JWT, serta dokumentasi API menggunakan OpenAPI/Swagger.

---

## 🚀 Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Web Router**: `github.com/julienschmidt/httprouter`
- **ORM**: `gorm.io/gorm`
- **PostgreSQL Driver**: `gorm.io/driver/postgres`
- **Dependency Injection**: `github.com/google/wire`
- **Authentication**: JWT (`github.com/golang-jwt/jwt/v5`)
- **Validation**: `github.com/go-playground/validator`
- **Configuration**: `github.com/spf13/viper`
- **Database Migration**: Golang Migrate

---

## ⚙️ Prerequisites

Pastikan tools berikut sudah terinstall di komputer kamu:

- **Go** ≥ 1.21
- **PostgreSQL**
- **Golang Migrate** (CLI)

### Install Golang Migrate

MacOS / Linux:

```
brew install golang-migrate
```

Linux (binary):

```
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

Windows:

- Download binary dari: [https://github.com/golang-migrate/migrate/releases](https://github.com/golang-migrate/migrate/releases)
- Tambahkan ke PATH

---

## 🔧 Configuration

Project ini menggunakan **Viper** untuk manajemen konfigurasi.

Buat file `config-local.env` berdasarkan contoh:

```
cp .env.example .env
```

Contoh isi `config-local.env`:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=contact_db
DB_SSLMODE=disable

JWT_SECRET=71ab959062e27c61f06d54eecdbc133d
```

---

## 🗄️ Database Migration

Sebelum menjalankan aplikasi, jalankan migrasi database:

```
migrate -database "postgres://user:password@localhost:5432/contact_db?sslmode=disable" \
  -path migrations up
```

Untuk rollback:

```
migrate -database "postgres://user:password@localhost:5432/contact_db?sslmode=disable" \
  -path migrations down
```

---

## ▶️ Running the Application

Install dependency:

```
go mod tidy
```

Generate Wire (Dependency Injection):

```
wire ./...
```

Jalankan aplikasi:

```
go run .
```

Aplikasi akan berjalan di:

```
http://localhost:8080
```

---

## 🔐 Authentication (JWT)

API ini menggunakan **JWT Bearer Token** untuk endpoint yang membutuhkan autentikasi.

Header yang dikirim:

```
Authorization: Bearer <your_jwt_token>
```

JWT digunakan untuk:

- Login
- Proteksi endpoint contact

---

## 📑 OpenAPI / Swagger Documentation

File **OpenAPI/Swagger** tersedia di **root project**:

```
apispec.yaml
```

File ini berisi:

- Daftar endpoint API
- Request & response schema
- Authentication (Bearer Token)
- Error response

### Cara menggunakan

1. Buka **Swagger Editor**: [https://editor.swagger.io/](https://editor.swagger.io/)
2. Upload file `apispec.yaml`
3. Atau gunakan Swagger UI lokal

Jika menggunakan Swagger UI, endpoint akan otomatis menampilkan:

- Input Bearer Token
- Contoh request/response

---

## ✅ Features

- CRUD Contact
- JWT Authentication
- Request validation
- Clean architecture (handler, service, repository)
- Dependency Injection dengan Wire
- Database migration
- OpenAPI documentation

---

## 🧪 Possible Improvements

- Unit & integration test
- Refresh token
- Pagination & search contact
- Docker support
- CI/CD pipeline

---

## 👨‍💻 Author

**Iwan Plamboyan**  
Mini project ini dibuat sebagai bagian dari pembelajaran dan eksplorasi backend development menggunakan Golang.

---

Happy coding 🚀
