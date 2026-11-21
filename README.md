# 🌊 Smart Aquaculture IoT System

<div align="center">

![Project Status](https://img.shields.io/badge/Status-Completed-success?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Clean_Architecture-orange?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**Sistem Otomasi dan Monitoring Kolam Ikan Pintar berbasis IoT — End-to-End**

[Fitur Utama](#-fitur-utama) · [Arsitektur](#-arsitektur-sistem) · [Tech Stack](#-tech-stack) · [Cara Install](#-panduan-instalasi) · [Alur Sistem](#-alur-komunikasi-sistem)

</div>

---

## 📖 Tentang Proyek

**Smart Aquaculture** adalah sistem IoT lengkap untuk monitoring dan otomasi kolam ikan secara real-time. Sistem ini menghubungkan mikrokontroler ESP32 dengan infrastruktur cloud modern — memungkinkan pembudidaya ikan memantau kondisi kolam dan mengontrol pemberian pakan dari mana saja melalui web atau aplikasi mobile.

Proyek ini dirancang sebagai sistem end-to-end yang mencakup seluruh lapisan: dari firmware perangkat keras, backend API, dashboard web, hingga aplikasi mobile dengan notifikasi push.

---

## ✨ Fitur Utama

| Fitur | Deskripsi |
|---|---|
| 🌡️ **Monitoring Real-Time** | Suhu air, kelembaban udara, dan intensitas cahaya kolam diperbarui setiap 5 detik |
| 🐟 **Remote Feeding Control** | Perintah pakan ke servo motor langsung dari Web atau Mobile via MQTT |
| 💡 **Otomasi Pencahayaan** | LED kolam menyala otomatis saat sensor LDR mendeteksi kegelapan (nilai > 2000) |
| 🚨 **Multi-Channel Alert** | Notifikasi darurat via **Telegram Bot** + **Firebase Push Notification** saat suhu > 35°C |
| 📊 **Historical Chart** | Grafik riwayat data sensor interaktif di Web (Recharts) dan Mobile (fl_chart) |
| 📱 **Cross-Platform Mobile** | Aplikasi Flutter native untuk Android & iOS |

---

## 🏗️ Arsitektur Sistem

```
┌─────────────────────────────────────────────────────────────────┐
│                        LAYER PRESENTASI                         │
│                                                                 │
│   ┌──────────────────────┐      ┌──────────────────────────┐   │
│   │   Web Dashboard      │      │   Mobile App (Flutter)   │   │
│   │   React + TypeScript │      │   Android / iOS          │   │
│   │   Vite + Tailwind    │      │   Firebase Messaging     │   │
│   └──────────┬───────────┘      └───────────┬──────────────┘   │
└──────────────┼───────────────────────────────┼──────────────────┘
               │  HTTP REST (Polling 5s)        │
               ▼                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                     BACKEND (Golang)                            │
│                   Clean Architecture                            │
│                                                                 │
│   ┌──────────┐  ┌──────────┐  ┌────────────┐  ┌───────────┐   │
│   │ Delivery │  │ Usecase  │  │ Repository │  │  Domain   │   │
│   │  (Fiber) │→│ (Logic)  │→│  (GORM)    │→│  (Models) │   │
│   └──────────┘  └──────────┘  └────────────┘  └───────────┘   │
│                                                                 │
│   ┌─────────────────┐   ┌────────────┐   ┌─────────────────┐  │
│   │ MQTT Worker     │   │ PostgreSQL │   │ Firebase Admin  │  │
│   │ (Background)    │   │ (Supabase) │   │ SDK (FCM)       │  │
│   └────────┬────────┘   └────────────┘   └─────────────────┘  │
└────────────┼────────────────────────────────────────────────────┘
             │  MQTT (broker.emqx.io)
             │  Topic: kolam/aquarium/telemetry
             │  Topic: kolam/aquarium/control
             ▼
┌─────────────────────────────────────────────────────────────────┐
│                   FIRMWARE (ESP32)                              │
│                   Arduino Framework                             │
│                                                                 │
│   DHT22 (Suhu+Kelembaban) │ LDR (Cahaya) │ Servo │ LED         │
│   Simulasi via Wokwi      │              │       │             │
└─────────────────────────────────────────────────────────────────┘
```

---

## 🛠️ Tech Stack

### Backend (`/backend`)
| Teknologi | Kegunaan |
|---|---|
| **Go 1.25** | Bahasa pemrograman utama |
| **Fiber v2** | HTTP Web Framework (Express-like for Go) |
| **GORM** | ORM untuk PostgreSQL |
| **Supabase (PostgreSQL)** | Database cloud |
| **Eclipse Paho MQTT** | MQTT client untuk komunikasi IoT |
| **Firebase Admin SDK** | Push notification via FCM |
| **Telegram Bot API** | Alert notifikasi darurat |
| **godotenv** | Manajemen environment variables |

### Frontend (`/frontend`)
| Teknologi | Kegunaan |
|---|---|
| **React 19** | UI Library |
| **TypeScript** | Type-safe JavaScript |
| **Vite 8** | Build tool & dev server |
| **Tailwind CSS v4** | Utility-first CSS framework |
| **Recharts** | Library grafik data sensor |
| **Axios** | HTTP client untuk API calls |
| **Lucide React** | Icon library |

### Mobile (`/mobile`)
| Teknologi | Kegunaan |
|---|---|
| **Flutter 3** | Cross-platform mobile framework |
| **Dart 3** | Bahasa pemrograman |
| **Firebase Messaging** | Push notification latar belakang |
| **fl_chart** | Grafik interaktif |
| **Google Fonts** | Custom typography |
| **http** | HTTP client untuk API calls |

### Firmware (`/firmware`)
| Teknologi | Kegunaan |
|---|---|
| **ESP32** | Mikrokontroler utama |
| **Arduino Framework** | Framework firmware |
| **DHT22** | Sensor suhu & kelembaban |
| **LDR** | Sensor intensitas cahaya |
| **Servo Motor** | Aktuator pemberian pakan |
| **PubSubClient** | MQTT client library |
| **ArduinoJson** | Parsing/serialisasi JSON |
| **Wokwi** | Simulator ESP32 online |

---

## 📁 Struktur Direktori

```
smart-aquaculture/
│
├── backend/                        # Go REST API + MQTT Worker
│   ├── delivery/http/              # HTTP handlers (Fiber routes)
│   │   ├── control_handler.go      # POST /control/feed
│   │   ├── monitoring_handler.go   # GET /monitoring/latest & /history
│   │   └── notification_handler.go # POST /fcm/token
│   ├── domain/
│   │   └── model.go                # Struct domain: PoolMonitoring, DeviceControl, DeviceToken
│   ├── infrastructure/
│   │   ├── database.go             # Koneksi Supabase + AutoMigrate
│   │   ├── firebase.go             # Inisialisasi Firebase Admin SDK
│   │   └── mqtt.go                 # MQTT broker connection + subscriber
│   ├── repository/
│   │   ├── monitoring_repository.go
│   │   └── token_repository.go
│   ├── usecase/
│   │   └── monitoring_usecase.go   # Business logic + Telegram/FCM alert
│   ├── main.go                     # Entry point
│   ├── go.mod
│   ├── .env.example                # ← Template konfigurasi (SALIN ke .env)
│   └── firebase-admin.json         # ← TIDAK di-commit (unduh dari Firebase Console)
│
├── frontend/                       # React Web Dashboard
│   ├── src/
│   │   ├── components/
│   │   │   ├── MetricCard.tsx      # Kartu metrik sensor
│   │   │   └── HistoryChart.tsx    # Grafik riwayat data
│   │   ├── App.tsx                 # Komponen utama + API calls
│   │   └── main.tsx
│   ├── package.json
│   └── vite.config.ts
│
├── mobile/                         # Flutter Mobile App
│   ├── lib/
│   │   ├── screens/
│   │   │   └── dashboard_screen.dart
│   │   ├── services/
│   │   │   ├── api_service.dart    # HTTP calls ke backend
│   │   │   └── fcm_service.dart    # Firebase push notification handler
│   │   ├── widgets/
│   │   │   ├── metric_card.dart
│   │   │   └── history_chart.dart
│   │   ├── theme.dart              # Custom dark theme
│   │   └── main.dart
│   └── pubspec.yaml
│
├── firmware/                       # ESP32 Simulator (Wokwi)
│   ├── sketch.ino                  # Kode firmware utama
│   ├── diagram.json                # Skema rangkaian Wokwi
│   └── libraries.txt               # Daftar library yang dibutuhkan
│
├── src/main.cpp                    # PlatformIO entry (placeholder)
├── platformio.ini                  # Konfigurasi PlatformIO
└── README.md
```

---

## ⚙️ Prasyarat

Pastikan sudah terinstall:

- [Go](https://go.dev/) ≥ 1.20
- [Node.js](https://nodejs.org/) ≥ 18 dan npm
- [Flutter SDK](https://flutter.dev/) ≥ 3.0 dan Android Studio
- Akun [Supabase](https://supabase.com/) — untuk database PostgreSQL
- Akun [Firebase](https://firebase.google.com/) — untuk push notification
- Akun [Telegram](https://telegram.org/) dan Bot Token dari [@BotFather](https://t.me/BotFather) — untuk alert
- Browser untuk [Wokwi](https://wokwi.com/) — simulator ESP32

---

## 🚀 Panduan Instalasi

### 1️⃣ Clone Repository

```bash
git clone https://github.com/faiqyr/smart-aquaculture.git
cd smart-aquaculture
```

---

### 2️⃣ Setup Backend (Golang)

**a. Masuk ke folder backend:**
```bash
cd backend
```

**b. Buat file konfigurasi dari template:**
```bash
# Windows
copy .env.example .env

# Linux/Mac
cp .env.example .env
```

**c. Edit `.env` dengan kredensial Anda:**
```env
DB_HOST=your-host.pooler.supabase.com
DB_PORT=6543
DB_USER=postgres.your_project_ref
DB_PASSWORD=your_password
DB_NAME=postgres
PORT=3000
MQTT_HOST=broker.emqx.io
MQTT_PORT=1883
MQTT_CLIENT_ID=smart_aqua_backend
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHAT_ID=your_chat_id
```

**d. Siapkan Firebase Admin SDK:**
1. Buka [Firebase Console](https://console.firebase.google.com/) → Project Anda
2. Settings → Service Accounts → **Generate new private key**
3. Simpan file JSON yang diunduh sebagai `backend/firebase-admin.json`

**e. Jalankan backend:**
```bash
go mod tidy
go run main.go
```

Backend berjalan di `http://localhost:3000`

**API Endpoints:**
| Method | Endpoint | Deskripsi |
|---|---|---|
| `GET` | `/api/v1/monitoring/latest` | Data sensor terbaru |
| `GET` | `/api/v1/monitoring/history` | Riwayat data sensor |
| `POST` | `/api/v1/control/feed` | Kirim perintah pakan ke ESP32 |
| `POST` | `/api/v1/fcm/token` | Daftarkan FCM token device |

---

### 3️⃣ Setup Frontend (Web Dashboard)

```bash
cd frontend
npm install
npm run dev
```

Dashboard web berjalan di `http://localhost:5173`

> **Catatan:** Pastikan backend sudah berjalan di port 3000. URL API dikonfigurasi di `frontend/src/App.tsx` pada variabel `API_BASE_URL`.

---

### 4️⃣ Setup Mobile (Flutter)

**a. Masuk ke folder mobile:**
```bash
cd mobile
```

**b. Siapkan Firebase untuk Android:**
1. Buka Firebase Console → Project → Project Settings → Your Apps
2. Tambahkan Android app dengan package name `com.example.mobile`
3. Unduh `google-services.json`
4. Letakkan file tersebut di `mobile/android/app/google-services.json`

**c. Install dependencies dan jalankan:**
```bash
flutter pub get
flutter run
```

> Pastikan emulator Android sudah berjalan atau perangkat fisik terhubung via USB.

---

### 5️⃣ Jalankan Simulator Firmware (Wokwi)

1. Buka [Wokwi ESP32 Simulator](https://wokwi.com/projects/new/esp32)
2. Salin isi `firmware/sketch.ino` ke editor kode Wokwi
3. Salin isi `firmware/diagram.json` ke tab diagram Wokwi
4. Buat file baru bernama `libraries.txt` di Wokwi, isi dengan konten dari `firmware/libraries.txt`
5. Tekan tombol **▶ Play** untuk mulai simulasi

Firmware akan otomatis terhubung ke WiFi Wokwi dan mulai mengirim data sensor ke broker MQTT.

---

## 🔄 Alur Komunikasi Sistem

```
[Wokwi ESP32]
    │
    │  1. Baca sensor (DHT22, LDR) setiap 5 detik
    │  2. Publish JSON ke MQTT topic: kolam/aquarium/telemetry
    │     {"temperature": 28.5, "humidity": 65.2, "light": 1500}
    ▼
[MQTT Broker: broker.emqx.io]
    │
    │  3. Backend Go subscribe ke topic telemetry
    ▼
[Backend Go]
    │
    │  4. Parse payload JSON
    │  5. Simpan ke Supabase PostgreSQL
    │  6. Cek threshold: suhu > 35°C?
    │     → Ya: kirim alert ke Telegram + FCM Push Notification
    ▼
[Supabase PostgreSQL]

[Web / Mobile] ──── Polling GET /monitoring/latest & /history setiap 5 detik ────▶ [Backend Go]

[User klik "Beri Pakan" di Web/Mobile]
    │
    │  POST /api/v1/control/feed
    ▼
[Backend Go]
    │
    │  Publish ke MQTT topic: kolam/aquarium/control
    │  {"command": "feed"}
    ▼
[Wokwi ESP32]
    │
    │  Terima command → Gerakkan Servo 0° → 90° (2 detik) → 0°
    ▼
[Servo Motor bergerak = Pakan dilepas]
```

---

## 🔐 Keamanan & Konfigurasi

> ⚠️ **PENTING**: File-file berikut **tidak disertakan** di repository karena berisi data rahasia. Anda harus membuatnya sendiri:

| File | Lokasi | Cara Mendapatkan |
|---|---|---|
| `.env` | `backend/.env` | Salin dari `backend/.env.example`, isi dengan kredensial Anda |
| `firebase-admin.json` | `backend/firebase-admin.json` | Unduh dari Firebase Console → Service Accounts |
| `google-services.json` | `mobile/android/app/` | Unduh dari Firebase Console → Project Settings |

Semua file sensitif tersebut sudah terdaftar di `.gitignore` dan **tidak akan pernah masuk ke git**.

---

## 🗄️ Skema Database

Tabel dikelola otomatis oleh GORM `AutoMigrate` saat backend pertama kali dijalankan.

### `pool_monitorings`
| Kolom | Tipe | Deskripsi |
|---|---|---|
| `id` | BIGINT (PK) | Auto-increment primary key |
| `temperature` | FLOAT | Suhu air (°C) |
| `humidity` | FLOAT | Kelembaban udara (%) |
| `light_intensity` | INT | Intensitas cahaya (nilai ADC 0-4095) |
| `created_at` | TIMESTAMP | Waktu data diterima |

### `device_controls`
| Kolom | Tipe | Deskripsi |
|---|---|---|
| `id` | BIGINT (PK) | Auto-increment primary key |
| `command` | VARCHAR | Perintah dikirim (contoh: "feed") |
| `created_at` | TIMESTAMP | Waktu perintah dikirim |

### `device_tokens`
| Kolom | Tipe | Deskripsi |
|---|---|---|
| `id` | BIGINT (PK) | Auto-increment primary key |
| `token` | VARCHAR | FCM device token untuk push notification |
| `created_at` | TIMESTAMP | Waktu token didaftarkan |

---

## 🔧 Konfigurasi Pin ESP32

| Pin | Komponen | Deskripsi |
|---|---|---|
| `GPIO 15` | DHT22 | Sensor suhu & kelembaban |
| `GPIO 34` | LDR | Sensor intensitas cahaya (ADC input) |
| `GPIO 13` | Servo Motor | Aktuator pemberian pakan |
| `GPIO 2` | LED Built-in | Indikator otomasi cahaya |

---

## 📊 Threshold & Parameter Sistem

| Parameter | Nilai | Keterangan |
|---|---|---|
| Interval telemetri | 5 detik | Frekuensi pengiriman data sensor |
| Feed duration | 2 detik | Durasi servo di posisi 90° saat memberi pakan |
| Alert suhu | > 35°C | Trigger notifikasi Telegram + FCM |
| LDR threshold | > 2000 | Nilai ADC trigger LED menyala (gelap) |
| MQTT Broker | broker.emqx.io | Public broker (port 1883) |
| Polling frontend | 5 detik | Interval fetch data di Web & Mobile |

---

## 🤝 Kontribusi

1. Fork repository ini
2. Buat branch fitur: `git checkout -b feature/nama-fitur`
3. Commit perubahan: `git commit -m "feat: tambahkan fitur X"`
4. Push ke branch: `git push origin feature/nama-fitur`
5. Buat Pull Request

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).

---

<div align="center">
Dikembangkan oleh <strong>Muhammad Faiq Yusuf Raharja</strong>
</div>
