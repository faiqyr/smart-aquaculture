# 🌊 Smart Aquaculture IoT System

![Project Status](https://img.shields.io/badge/Status-Completed-success)
![Version](https://img.shields.io/badge/Version-1.0.0-blue)
![Architecture](https://img.shields.io/badge/Architecture-Clean_Architecture-orange)

Sistem Otomasi dan Monitoring Kolam Ikan Pintar (*Smart Aquaculture*) berskala penuh (*End-to-End*). Proyek ini dirancang untuk membaca data sensor lingkungan kolam, memprosesnya melalui infrastruktur *backend* berkinerja tinggi, dan menyajikannya secara *real-time* ke dasbor web maupun aplikasi *mobile* native. Sistem ini juga dilengkapi dengan kontrol aktuator otomatis dan notifikasi darurat.

---

## 🏗️ Arsitektur Proyek

Sistem ini terbagi menjadi 4 komponen utama yang saling terhubung menggunakan protokol **HTTP REST** dan **MQTT**:

1. **`backend/` (Golang)**
   - Menggunakan *Clean Architecture* (Domain, Usecase, Repository, Delivery).
   - **Framework API**: Fiber v2.
   - **Database**: PostgreSQL (Supabase) via GORM.
   - **Fitur Spesial**: *MQTT Background Worker*, Integrasi Telegram Bot API, dan Firebase Cloud Messaging (FCM) Admin SDK untuk *Push Notifications*.
   
2. **`frontend/` (Web Dashboard)**
   - **Teknologi**: React, TypeScript, Vite.
   - **Desain**: Tailwind CSS v4 (*Premium Aqua Dark-mode*), Lucide Icons.
   - **Grafik**: Recharts (Pembaruan data asinkron/Polling setiap 5 detik).

3. **`mobile/` (Aplikasi Android/iOS)**
   - **Teknologi**: Flutter.
   - **Desain**: Kustomisasi Tema *Dark Navy* menggunakan Google Fonts.
   - **Fitur Spesial**: Notifikasi Push Latar Belakang (Firebase Messaging), dan Grafik Interaktif (`fl_chart`).

4. **`firmware/` (Simulator Perangkat Keras)**
   - **Target**: Mikrokontroler ESP32 (Disimulasikan via Wokwi).
   - **Perangkat**: Sensor Suhu DHT22, Sensor Cahaya LDR, Motor Servo, LED.
   - **Konektivitas**: WiFi Wokwi-GUEST, MQTT Client (`PubSubClient`).

---

## ✨ Fitur Utama
* 🌡️ **Monitoring Real-Time**: Pemantauan suhu air, kelembaban udara, dan intensitas cahaya kolam.
* 🐟 **Remote Feeding Control**: Kendali pakan otomatis dengan menggerakkan motor Servo langsung dari tombol di Web/Mobile.
* 💡 **Otomasi Pencahayaan LDR**: Lampu LED kolam akan otomatis menyala di malam hari (ketika sensor LDR mendeteksi kegelapan).
* 🚨 **Multi-Channel Alert System**: Jika suhu kolam melebihi batas (35°C), sistem akan mengirimkan pesan peringatan via **Telegram** dan memancarkan **Push Notification** ke HP via Firebase secara bersamaan.

---

## ⚙️ Prasyarat (*Prerequisites*)
Pastikan komputer Anda sudah terinstal:
* [Golang](https://go.dev/) (Minimal versi 1.20)
* [Node.js](https://nodejs.org/) & npm
* [Flutter SDK](https://flutter.dev/) & Android Studio
* Akun [Supabase](https://supabase.com/) (Untuk Database)
* Akun [Firebase](https://firebase.google.com/) (Untuk Push Notification)

---

## 🚀 Panduan Instalasi & Menjalankan (*How to Run*)

### 1. Menyiapkan Backend (Golang)
Buka terminal dan masuk ke folder `backend`:
```bash
cd backend
```
Buat file konfigurasi rahasia:
- Ubah nama `backend/.env.example` menjadi `backend/.env`.
- Isi kredensial `DB_PASSWORD` (Supabase) dan `TELEGRAM_BOT_TOKEN`.
- Letakkan file `firebase-admin.json` Anda di dalam folder ini (Didapat dari Firebase Console -> Service Accounts).

Jalankan server:
```bash
go mod tidy
go run main.go
```
*Server akan menyala di `http://localhost:3000`.*

---

### 2. Menyiapkan Frontend (Web React)
Buka terminal baru dan masuk ke folder `frontend`:
```bash
cd frontend
npm install
npm run dev
```
*Dashboard Web bisa langsung diakses di `http://localhost:5173`.*

---

### 3. Menyiapkan Aplikasi Mobile (Flutter)
Buka terminal baru (sangat disarankan via Android Studio) dan masuk ke folder `mobile`:
```bash
cd mobile
```
- Pastikan Anda meletakkan file `google-services.json` ke dalam `mobile/android/app/`.
- Nyalakan Emulator Android (AVD) di Android Studio.

Jalankan aplikasi ke Emulator:
```bash
flutter run
```

---

### 4. Menyalakan Simulator Fisik (Wokwi)
1. Buka [Wokwi ESP32](https://wokwi.com/projects/new/esp32).
2. Salin isi `firmware/sketch.ino` ke tab kode C++ di Wokwi.
3. Salin isi `firmware/diagram.json` ke tab skema Wokwi.
4. Buat file baru di Wokwi bernama `libraries.txt` dan isi dengan data dari `firmware/libraries.txt`.
5. Tekan tombol **Play (▶)** di Wokwi.

---

## 🛠️ Alur Komunikasi Sistem (Cara Kerja Singkat)
1. Wokwi membaca sensor dan mem-*publish* data berformat JSON ke broker MQTT publik (`broker.emqx.io`) pada topik `kolam/aquarium/telemetry`.
2. *Background Worker* Golang menerima pesan MQTT tersebut dan menyimpannya secara permanen ke Supabase PostgreSQL.
3. Aplikasi Web dan Flutter melakukan *Polling* (memanggil API Golang secara berkala) untuk merender grafik.
4. Saat pengguna menekan "Beri Pakan" di UI, aplikasi menembak `POST /api/v1/control/feed` ke Golang.
5. Golang lalu mem-*publish* pesan berformat `{"command": "feed"}` ke topik MQTT `kolam/aquarium/control`.
6. ESP32 di Wokwi menerimanya dan menggerakkan Servo.

---
*Didesain dan dikembangkan dengan ❤️ oleh Muhammad Faiq Yusuf Raharja.*

