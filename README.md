# Backend Jual Beli Barang Bekas

## **Deskripsi Proyek**

Proyek ini adalah backend untuk sistem "Jual dan Beli Barang Bekas" menggunakan **Golang**, **Gin Framework**, dan **MySQL**. Backend ini mendukung operasi CRUD untuk pengguna (users), barang (items), dan pesanan (orders) dengan pengamanan endpoint menggunakan JWT (JSON Web Token).

---

## **Fitur Utama**

1. **Autentikasi JWT**:
   - Registrasi dan login pengguna.
   - Middleware untuk pengamanan endpoint.
2. **CRUD Entities**:
   - Pengguna (Users): Tambah, lihat, ubah, hapus pengguna.
   - Barang (Items): Tambah, lihat, ubah, hapus barang.
   - Pesanan (Orders): Tambah, lihat, ubah, hapus pesanan.

---

## **Kebutuhan Sistem**

- **Go** versi 1.19 atau lebih baru.
- **MySQL** versi 8.0 atau lebih baru.
- **Postman** (untuk pengujian API).
- **XAMPP** (opsional, untuk menjalankan MySQL lokal).

---

## **Langkah Instalasi**

### **1. Clone Repository**

```bash
git clone <repository-url>
cd backend-jual-beli
```

### **2. Install Dependency**

```bash
go mod tidy
```

### **3. Buat File .env**

Buat file .env di root proyek dengan isi seperti berikut:
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=jual_beli_db
JWT_SECRET=your_secret_key

### **4. Setup Database**

Jika menggunakan MySQL lokal atau XAMPP, buat database dengan nama jual_beli_db:
```bash
CREATE DATABASE jual_beli_db;
```

### **5. Jalankan Aplikasi**

```bash
go run main.go
```

## **Endpoint API**

### **Autentikasi**

| HTTP Method | Endpoint       | Deskripsi             |
|-------------|----------------|-----------------------|
| POST        | `/api/register`| Registrasi pengguna   |
| POST        | `/api/login`   | Login pengguna        |

### **Users**

| HTTP Method | Endpoint         | Deskripsi             |
|-------------|------------------|-----------------------|
| GET         | `/api/users`     | Lihat semua pengguna  |
| GET         | `/api/users/:id` | Lihat detail pengguna |
| PUT         | `/api/users/:id` | Update pengguna       |
| DELETE      | `/api/users/:id` | Hapus pengguna        |

### **Items**

| HTTP Method | Endpoint         | Deskripsi             |
|-------------|------------------|-----------------------|
| POST        | `/api/items`     | Tambah barang         |
| GET         | `/api/items`     | Lihat semua barang    |
| GET         | `/api/items/:id` | Lihat detail barang   |
| PUT         | `/api/items/:id` | Update barang         |
| DELETE      | `/api/items/:id` | Hapus barang          |

### **Orders**

| HTTP Method | Endpoint          | Deskripsi             |
|-------------|-------------------|-----------------------|
| POST        | `/api/orders`     | Tambah pesanan        |
| GET         | `/api/orders`     | Lihat semua pesanan   |
| GET         | `/api/orders/:id` | Lihat detail pesanan  |
| PUT         | `/api/orders/:id` | Update pesanan        |
| DELETE      | `/api/orders/:id` | Hapus pesanan         |
