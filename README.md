# 🚗 Fleetify Backend - Fachransandi

Sebuah perusahaan Multinasional memiliki jumlah karyawan diatas 50 karyawan, dan memiliki berbagai macam Divisi atau departemen didalamnya. Karena banyaknya karyawan untuk dikelola, perusahaan membutuhkan Sistem untuk Absensi guna mencatat serta mengevaluasi kedisiplinan karyawan secara sistematis.

---

## 📊 Dokumentasi API (Sederhana)

Berikut adalah endpoint-endpoint utama dari aplikasi ini.

### 💼 Karyawan

| Method | Endpoint        | Deskripsi                            |
| ------ | --------------- | ------------------------------------ |
| GET    | `/employee`     | List semua karyawan                  |
| GET    | `/employee/:id` | Ambil detail karyawan berdasarkan ID |
| POST   | `/employee`     | Tambah karyawan baru                 |
| PUT    | `/employee/:id` | Update data karyawan                 |
| DELETE | `/employee/:id` | Hapus karyawan                       |

### 🏢 Departemen

| Method | Endpoint          | Deskripsi                              |
| ------ | ----------------- | -------------------------------------- |
| GET    | `/department`     | List semua departemen                  |
| GET    | `/department/:id` | Ambil detail departemen berdasarkan ID |
| POST   | `/department`     | Tambah departemen baru                 |
| PUT    | `/department/:id` | Update data departemen                 |
| DELETE | `/department/:id` | Hapus departemen                       |

### ⏰ Absensi

| Method | Endpoint               | Deskripsi                                                              |
| ------ | ---------------------- | ---------------------------------------------------------------------- |
| POST   | `/attendance/checkin`  | Absen masuk                                                            |
| PUT    | `/attendance/checkout` | Absen keluar                                                           |
| GET    | `/attendance/logs`     | List log absensi lengkap + ketepatan, bisa filter tanggal & departemen |

**Catatan:** Endpoint `/attendance/logs` akan menampilkan informasi ketepatan waktu (tepat, telat, pulang cepat) berdasarkan `max_clock_in_time` dan `max_clock_out_time` dari masing-masing departemen.

---

## 🎓 Studi Kasus

- CRUD Karyawan
- CRUD Departemen
- POST Absen Masuk
- PUT Absen Keluar
- GET Log Absensi (dengan ketepatan & filter tanggal/departemen)

---

## 🚀 Cara Menjalankan Project

### 1. Clone Repository

```bash
git clone https://github.com/tiedsandi/fleetify-backend-fachransandi.git
cd fleetify-backend-fachransandi
```

### 2. Jalankan MySQL & Buat Database

Pastikan MySQL sudah berjalan di `127.0.0.1:3306`. Kemudian buat database baru dengan nama `fleetify`:

```sql
CREATE DATABASE fleetify;
```

### 3. Install Dependency

```bash
go mod tidy
```

### 4. Jalankan

```bash
go run .
```

---

## 🧪 (Optional) Seeder & Reset

```bash
 #main.go
 seeds.AddSeeder()

 config.ResetDB()
```
