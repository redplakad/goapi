# Integrasi Backend Go API dengan Frontend Laravel

## 🎯 Overview

Backend API yang telah dibuat menggunakan **Go (Golang)** dengan framework **Gin** dan database **MySQL** untuk melayani data dari tabel `nominatif_kredit`. API ini dirancang khusus untuk dikonsumsi oleh frontend **Laravel**.

## 📁 Struktur Project Go API

```
goapi/
├── main.go                          # Entry point & routing
├── config.go                        # Environment configuration  
├── database.go                      # Database connection
├── models/
│   └── nominatif_kredit.go         # Model & repository layer
├── handlers/
│   └── nominatif_kredit.go         # HTTP handlers (controllers)
├── middleware/
│   └── cors.go                     # CORS middleware untuk Laravel
├── examples/
│   ├── LaravelNominatifKreditController.php
│   └── laravel_routes_example.php
├── .env                            # Environment variables
├── .env.example                    # Environment template
├── README.md                       # Documentation
└── API_DOCUMENTATION.md            # API endpoints documentation
```

## 🚀 Cara Menjalankan

### 1. Backend Go API
```bash
cd /root/project/goapi
go run .
```
Server akan berjalan di: `http://localhost:8080`

### 2. Test API
```bash
# Health check
curl http://localhost:8080/health

# Get all data dengan pagination
curl "http://localhost:8080/api/nominatif-kredit?page=1&per_page=5"

# Get by ID
curl http://localhost:8080/api/nominatif-kredit/1

# Get by Nomor Rekening
curl http://localhost:8080/api/nominatif-kredit/rekening/123456789
```

## 🔗 API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | `/api/nominatif-kredit` | Get all data dengan pagination |
| GET | `/api/nominatif-kredit/:id` | Get data by ID |
| GET | `/api/nominatif-kredit/rekening/:nomor_rekening` | Get data by nomor rekening |
| GET | `/health` | Database health check |
| GET | `/ping` | Basic health check |

## 📝 Response Format

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

## 🎨 Integrasi dengan Laravel Frontend

### A. Buat Controller di Laravel

```php
// app/Http/Controllers/Api/NominatifKreditController.php
<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class NominatifKreditController extends Controller
{
    private $apiBaseUrl = 'http://localhost:8080/api';

    public function index(Request $request)
    {
        $response = Http::get($this->apiBaseUrl . '/nominatif-kredit', [
            'page' => $request->get('page', 1),
            'per_page' => $request->get('per_page', 10)
        ]);

        return response()->json($response->json());
    }

    public function show($id)
    {
        $response = Http::get($this->apiBaseUrl . '/nominatif-kredit/' . $id);
        return response()->json($response->json());
    }
}
```

### B. Setup Routes di Laravel

```php
// routes/api.php
use App\Http\Controllers\Api\NominatifKreditController;

Route::prefix('v1')->group(function () {
    Route::get('/nominatif-kredit', [NominatifKreditController::class, 'index']);
    Route::get('/nominatif-kredit/{id}', [NominatifKreditController::class, 'show']);
});
```

### C. Consume API via JavaScript/AJAX

```javascript
// Frontend JavaScript
fetch('/api/v1/nominatif-kredit?page=1&per_page=10')
.then(response => response.json())
.then(data => {
    if (data.success) {
        console.log('Data:', data.data);
        console.log('Pagination:', data.meta);
        
        // Render data ke table atau component
        renderTable(data.data);
        renderPagination(data.meta);
    }
})
.catch(error => console.error('Error:', error));
```

### D. Blade Template Example

```html
<!-- resources/views/nominatif-kredit/index.blade.php -->
<div class="table-responsive">
    <table class="table table-striped">
        <thead>
            <tr>
                <th>No. Rekening</th>
                <th>Nama Nasabah</th>
                <th>Baki Debet</th>
                <th>Status</th>
            </tr>
        </thead>
        <tbody>
            @foreach($kredits as $kredit)
            <tr>
                <td>{{ $kredit['NOMOR_REKENING'] }}</td>
                <td>{{ $kredit['NAMA_NASABAH'] }}</td>
                <td>{{ number_format($kredit['BAKI_DEBET'], 2) }}</td>
                <td>{{ $kredit['KODE_KOLEK'] }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
</div>

<!-- Pagination -->
@if(isset($meta['last_page']) && $meta['last_page'] > 1)
<nav>
    <ul class="pagination">
        @for($i = 1; $i <= $meta['last_page']; $i++)
        <li class="page-item {{ $i == $meta['current_page'] ? 'active' : '' }}">
            <a class="page-link" href="?page={{ $i }}">{{ $i }}</a>
        </li>
        @endfor
    </ul>
</nav>
@endif
```

## ⚙️ Konfigurasi Environment

### Go API (.env)
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=garuda1*
DB_NAME=monitoring
PORT=8080
APP_ENV=development
```

### Laravel (.env)
```env
# Tambahkan konfigurasi untuk Go API
GO_API_BASE_URL=http://localhost:8080/api
```

## 🔧 Fitur Yang Sudah Diimplementasi

✅ **Database Connection** - Koneksi otomatis ke MySQL  
✅ **CORS Support** - Laravel dapat akses API tanpa masalah  
✅ **Pagination** - Data besar dapat dibagi per halaman  
✅ **Error Handling** - Response error yang konsisten  
✅ **Environment Config** - Konfigurasi fleksibel  
✅ **Model Conversion** - Konversi lengkap dari Laravel model ke Go struct  
✅ **Repository Pattern** - Clean architecture pattern  
✅ **Standard Response** - Format JSON yang konsisten  

## 📊 Data Fields

Model Go include semua field dari Laravel model:
- ID, DATADATE, CAB, NOMOR_REKENING, NO_CIF
- NAMA_NASABAH, ALAMAT, KODE_KOLEK
- Financial fields: PLAFOND_AWAL, BAKI_DEBET, TUNGGAKAN_POKOK, dll.
- Date fields: TGL_PK, TGL_AWAL_FAS, TGL_AKHIR_FAS, dll.
- Dan 80+ field lainnya sesuai model Laravel

## 🚦 Next Steps

1. **Test dengan data real** - Pastikan tabel `nominatif_kredit` memiliki data
2. **Performance tuning** - Add indexing untuk query yang sering digunakan
3. **Add authentication** - JWT atau API key jika diperlukan
4. **Add filtering** - Filter berdasarkan CAB, KODE_KOLEK, dll.
5. **Add caching** - Redis untuk performance yang lebih baik

## 🔍 Troubleshooting

1. **CORS Error**: Middleware CORS sudah disetup untuk semua origin
2. **Database Connection**: Check kredensial di `.env` dan test dengan `/health`
3. **Port Conflict**: Ubah PORT di `.env` jika 8080 sudah digunakan
4. **Laravel HTTP Client**: Pastikan `guzzlehttp/guzzle` terinstall

Backend Go API Anda sudah siap untuk dikonsumsi oleh frontend Laravel! 🎉
