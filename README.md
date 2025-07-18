# 🚗 Fleetify Backend - Fachransandi

Fleetify adalah sistem manajemen kehadiran dan karyawan berbasis REST API, dibangun menggunakan Golang (`gin`) dan MySQL. Project ini merupakan bagian dari **studi kasus rekrutmen** berdasarkan ERD, flowchart, dan data tabel yang telah ditentukan.

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

### 3. Buat file `.env`

```env
DB_USER=root
DB_PASS=password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=fleetify
```

### 4. Install Dependency

```bash
go mod tidy
```

### 5. Jalankan

```bash
go run .
```

---

## 🧪 (Optional) Seeder & Testing

```bash
go run seed/seeder.go
go test ./...
```

---

## 👨‍💻 Author

Fachransandi ([@tiedsandi](https://github.com/tiedsandi))

---

## 📜 License

This project is open-source under the [MIT License](LICENSE).
