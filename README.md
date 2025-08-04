# 🚀 Go API - Nominatif Kredit

A secure, high-performance REST API built with Go and Gin framework for managing Nominatif Kredit data with advanced filtering capabilities and API Key authentication.

## ✨ Features

- 🔐 **API Key Authentication** - Secure access without complex login systems
- 🔍 **Advanced Filtering** - Filter by CAB, AO, KET_KD_PRD, TEMPAT_BEKERJA
- 📄 **Pagination Support** - Efficient handling of large datasets (38K+ records)  
- 🌐 **CORS Enabled** - Ready for frontend integration
- 🚀 **High Performance** - Optimized database queries
- 📱 **Mobile Ready** - RESTful API design
- 🔧 **Laravel Integration** - Built for Laravel frontend consumption

## 🛠️ Tech Stack

- **Go 1.22.2** - Core language
- **Gin Framework** - HTTP web framework  
- **MySQL** - Database
- **godotenv** - Environment configuration
- **CORS Middleware** - Cross-origin support

## Struktur Project

```
goapi/
├── main.go                    # Entry point aplikasi
├── config.go                  # Konfigurasi environment
├── database.go                # Database connection
├── models/
│   └── nominatif_kredit.go    # Model dan repository
├── handlers/
│   └── nominatif_kredit.go    # HTTP handlers/controllers
├── middleware/
│   └── cors.go                # CORS middleware
├── .env                       # Environment variables
├── .env.example               # Environment template
└── API_DOCUMENTATION.md       # API documentation
```

## Setup

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Setup Database
Pastikan MySQL server sudah berjalan dan database `monitoring` sudah dibuat.

### 3. Environment Configuration
1. Copy file `.env.example` ke `.env`:
   ```bash
   cp .env.example .env
   ```

2. Edit file `.env` dan sesuaikan dengan konfigurasi database Anda:
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=secret
   DB_NAME=monitoring
   ```

### 4. Jalankan Aplikasi
```bash
go run .
```

Aplikasi akan berjalan di `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /ping` - Basic health check
- `GET /health` - Database connection health check

### Nominatif Kredit API
- `GET /api/nominatif-kredit` - Get all data dengan pagination
- `GET /api/nominatif-kredit/:id` - Get data by ID
- `GET /api/nominatif-kredit/rekening/:nomor_rekening` - Get data by nomor rekening

## Fitur

✅ **CORS Support** - Frontend Laravel dapat mengakses API  
✅ **Pagination** - Data besar dapat dibagi per halaman  
✅ **Error Handling** - Response error yang konsisten  
✅ **Environment Config** - Konfigurasi fleksibel via file .env  
✅ **Database Connection** - Auto-reconnect MySQL  
✅ **Standard Response** - Format response JSON yang konsisten  

## Format Response

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [...],
  "meta": {
    "current_page": 1,
    "per_page": 10,
    "total": 100,
    "last_page": 10
  }
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | MySQL host | localhost |
| DB_PORT | MySQL port | 3306 |
| DB_USER | MySQL username | root |
| DB_PASSWORD | MySQL password | - |
| DB_NAME | MySQL database name | - |
| PORT | Server port | 8080 |
| APP_ENV | Application environment | development |

## Testing API

### 1. Test Health Check
```bash
curl http://localhost:8080/health
```

### 2. Test Get All Nominatif Kredit
```bash
curl "http://localhost:8080/api/nominatif-kredit?page=1&per_page=5"
```

### 3. Test Get by ID
```bash
curl http://localhost:8080/api/nominatif-kredit/1
```

### 4. Test Get by Nomor Rekening
```bash
curl http://localhost:8080/api/nominatif-kredit/rekening/123456789
```

## Integrasi dengan Laravel Frontend

Di Laravel, Anda dapat menggunakan HTTP facade untuk consume API:

```php
use Illuminate\Support\Facades\Http;

// Controller method
public function getNominatifKredit()
{
    $response = Http::get('http://localhost:8080/api/nominatif-kredit', [
        'page' => request('page', 1),
        'per_page' => 20
    ]);

    if ($response->successful()) {
        return $response->json();
    }

    return response()->json(['error' => 'Failed to fetch data'], 500);
}
```

## Troubleshooting

### 1. CORS Error
Pastikan middleware CORS sudah aktif. API sudah include CORS header untuk semua origin.

### 2. Database Connection Error
- Pastikan MySQL service berjalan
- Verify kredensial database di file .env
- Test koneksi dengan endpoint `/health`

### 3. Port Already in Use
Ubah port di file .env jika port 8080 sudah digunakan:
```
PORT=8081
```

## Development

### Menambah Endpoint Baru
1. Tambahkan method di repository (`models/nominatif_kredit.go`)
2. Tambahkan handler di `handlers/nominatif_kredit.go`
3. Register route di `main.go`

### Database Migration
Pastikan struktur tabel `nominatif_kredit` sesuai dengan field di model Go.

Untuk dokumentasi API lengkap, lihat file `API_DOCUMENTATION.md`.
