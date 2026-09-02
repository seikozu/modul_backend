# Dokumentasi REST API Students

| Metode | Endpoint | Parameter | Contoh Body | Status | Contoh Respons |
|---|---|---|---|---|---|
| **GET** | `/api/v1/students` | `page`, `limit`, `search`, `sort`, `order`, `is_active` | - | `200 OK` | `{"success":true, "data":[...], "meta":{...}}` |
| **GET** | `/api/v1/students/:id` | `id` (path) | - | `200 OK`, `404 Not Found` | `{"success":true, "data":{"id":1, "nim":"123", ...}}` |
| **POST** | `/api/v1/students` | - | `{"nim":"123", "name":"Jinhsi", "grade":90}` | `201 Created`, `422 Unprocessable`| `{"success":true, "message":"student berhasil dibuat", "data":{...}}` |
| **PUT** | `/api/v1/students/:id` | `id` (path) | `{"nim":"123", "name":"Baru", "grade":95, "is_active":true}` | `200 OK`, `422 Unprocessable` | `{"success":true, "message":"student berhasil diganti seluruhnya", "data":{...}}` |
| **PATCH** | `/api/v1/students/:id` | `id` (path) | `{"is_active":false}` | `200 OK`, `400 Bad Request` | `{"success":true, "message":"student berhasil diperbarui sebagian", "data":{...}}` |
| **DELETE**| `/api/v1/students/:id` | `id` (path) | - | `204 No Content`, `404 Not Found` | *(Tanpa body)* |

---

# Student Management API 

Repositori ini berisi RESTful API untuk mengelola data mahasiswa (*students*) menggunakan **Go (Fiber)**, **PostgreSQL** (`pgx`), serta menerapkan **Repository Pattern**.

## 📄 Daftar Variabel Environment

Menggunakan berkas `.env` untuk menyimpan konfigurasi koneksi dan port. 

Sebelum menjalankan, buat berkas bernama `.env` di root folder proyek (setara dengan `main.go`), lalu isi dengan variabel berikut:

APP_PORT=3000
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
DB_MAX_CONNS=

> **Catatan:**
> * Sesuaikan `DB_PASSWORD` dengan password user `postgres` di laptop masing-masing.
> * `DB_NAME` diisi sesuai nama database yang dibuat di PostgreSQL 

## 🗄️ Skema Tabel Basis Data

Tabel yang digunakan bernama `students` dengan tambahan *Index* untuk mengoptimalkan kinerja pencarian serta menjaga keunikan NIM:

CREATE TABLE IF NOT EXISTS students (
    id         SERIAL       PRIMARY KEY,
    nim        VARCHAR(50)  NOT NULL,
    name       VARCHAR(255) NOT NULL,
    grade      NUMERIC(5,2) NOT NULL DEFAULT 0,
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Indeks unik untuk NIM (case-insensitive)
CREATE UNIQUE INDEX IF NOT EXISTS students_nim_lower_key
    ON students (LOWER(nim));

-- Indeks untuk mempercepat pencarian berdasarkan nama
CREATE INDEX IF NOT EXISTS students_name_lower_idx
    ON students (LOWER(name));

## Cara Menyiapkan & Menjalankan dari Nol

Ikuti langkah-langkah berikut untuk menjalankan proyek ini di komputer lokal:

### 1. Prasyarat
Pastikan sudah menginstal:
* [Go](https://go.dev/) (versi 1.20+)
* [PostgreSQL](https://www.postgresql.org/) & pgAdmin
* Git

### 2. Persiapan Basis Data (PostgreSQL)
1. Buka **pgAdmin** atau terminal `psql`.
2. Buat database baru bernama `prak_backend`:

   CREATE DATABASE prak_backend;

3. Buka **Query Tool** di dalam database `prak_backend`.
4. Salin dan jalankan script SQL yang ada pada bagian **Skema Tabel Basis Data** di atas (atau ambil dari file `migrations/001_create_students.sql`).

### 3. Setup Repositori & Dependencies
1. Buka terminal, masuk ke direktori proyek:

   cd "modul 3"

2. Buat berkas `.env` dan atur variabel sesuai konfigurasi PostgreSQL lokalmu.
3. Unduh seluruh *dependencies* yang dibutuhkan:

   go mod tidy

### 4. Menjalankan Aplikasi
Jalankan aplikasi dengan perintah:

go run .

Jika berhasil, terminal akan menampilkan pesan:
Server berjalan di http://localhost:3000

## Endpoint API Utama

| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/api/v1/health` | Cek status server & koneksi database |
| `GET` | `/api/v1/students` | Mengambil daftar student (mendukung `search`, `page`, `limit`, `sort`, `order`) |
| `GET` | `/api/v1/students/:id` | Mengambil detail student berdasarkan ID |
| `POST` | `/api/v1/students` | Menambah student baru |
| `PUT` | `/api/v1/students/:id` | Mengganti seluruh data student |
| `PATCH` | `/api/v1/students/:id` | Mengubah sebagian data student |
| `DELETE` | `/api/v1/students/:id` | Menghapus data student |

---

## atatan Pengujian (Postman / Thunder Client)
* Untuk request ber-body (`POST`, `PUT`, `PATCH`), sertakan header `Content-Type: application/json`.
* Jika mencoba memasukkan NIM yang sudah ada, API akan mengembalikan respons `409 Conflict`.