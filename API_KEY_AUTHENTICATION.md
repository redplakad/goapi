# 🔐 API Key Authentication - Go API Nominatif Kredit

## ✅ Implementasi Berhasil!

API sekarang telah diamankan menggunakan **API Key Authentication** tanpa perlu user login. Sistem ini memungkinkan frontend Laravel (atau aplikasi lain) untuk mengakses API dengan mudah menggunakan API key yang telah ditentukan.

## 🔑 API Keys yang Tersedia

| API Key | Aplikasi | Deskripsi | Status |
|---------|----------|-----------|---------|
| `api-key-laravel-frontend-2025` | Laravel Frontend | Main Laravel frontend application | ✅ Active |
| `api-key-mobile-app-2025` | Mobile App | Mobile application | ✅ Active |
| `api-key-dashboard-admin-2025` | Admin Dashboard | Administrative dashboard | ✅ Active |

## 📝 Cara Penggunaan

### **1. Menggunakan X-API-Key Header (Recommended)**
```bash
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/nominatif-kredit"
```

### **2. Menggunakan Authorization Header**
```bash
curl -H "Authorization: ApiKey api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/nominatif-kredit"
```

## 🚀 Contoh Penggunaan Lengkap

### **Basic Request dengan Pagination**
```bash
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/nominatif-kredit?page=1&per_page=20"
```

### **Request dengan Filter**
```bash
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/nominatif-kredit?cab=007&ao=JAENUDIN&page=1&per_page=10"
```

### **Validasi API Key**
```bash
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/validate"
```

## 🔧 Integrasi Laravel Frontend

### **1. Konfigurasi di Laravel (.env)**
```env
GO_API_BASE_URL=http://localhost:8081/api
GO_API_KEY=api-key-laravel-frontend-2025
```

### **2. Service Class Laravel**
```php
<?php

namespace App\Services;

use Illuminate\Support\Facades\Http;

class GoAPIService
{
    private string $baseUrl;
    private string $apiKey;

    public function __construct()
    {
        $this->baseUrl = config('services.go_api.base_url');
        $this->apiKey = config('services.go_api.api_key');
    }

    public function getNominatifKredit(array $params = [])
    {
        return Http::withHeaders([
            'X-API-Key' => $this->apiKey,
            'Accept' => 'application/json',
        ])->get("{$this->baseUrl}/nominatif-kredit", $params);
    }

    public function getNominatifKreditById(int $id)
    {
        return Http::withHeaders([
            'X-API-Key' => $this->apiKey,
            'Accept' => 'application/json',
        ])->get("{$this->baseUrl}/nominatif-kredit/{$id}");
    }

    public function getNominatifKreditByRekening(string $nomorRekening)
    {
        return Http::withHeaders([
            'X-API-Key' => $this->apiKey,
            'Accept' => 'application/json',
        ])->get("{$this->baseUrl}/nominatif-kredit/rekening/{$nomorRekening}");
    }
}
```

### **3. Controller Laravel**
```php
<?php

namespace App\Http\Controllers;

use App\Services\GoAPIService;
use Illuminate\Http\Request;

class NominatifKreditController extends Controller
{
    private GoAPIService $goApi;

    public function __construct(GoAPIService $goApi)
    {
        $this->goApi = $goApi;
    }

    public function index(Request $request)
    {
        $params = $request->only(['page', 'per_page', 'cab', 'ao', 'ket_kd_prd', 'tempat_bekerja']);
        
        $response = $this->goApi->getNominatifKredit($params);
        
        if ($response->successful()) {
            $data = $response->json();
            return view('nominatif-kredit.index', compact('data'));
        }
        
        return view('nominatif-kredit.index', [
            'data' => ['success' => false, 'message' => 'Failed to fetch data']
        ]);
    }
}
```

### **4. Blade Template dengan Filter**
```html
<!-- resources/views/nominatif-kredit/index.blade.php -->
<div class="container">
    <h1>Nominatif Kredit</h1>
    
    <!-- Filter Form -->
    <form method="GET" class="mb-4">
        <div class="row">
            <div class="col-md-3">
                <input type="text" name="cab" placeholder="CAB" 
                       value="{{ request('cab') }}" class="form-control">
            </div>
            <div class="col-md-3">
                <input type="text" name="ao" placeholder="AO" 
                       value="{{ request('ao') }}" class="form-control">
            </div>
            <div class="col-md-3">
                <input type="text" name="ket_kd_prd" placeholder="Product" 
                       value="{{ request('ket_kd_prd') }}" class="form-control">
            </div>
            <div class="col-md-3">
                <input type="text" name="tempat_bekerja" placeholder="Workplace" 
                       value="{{ request('tempat_bekerja') }}" class="form-control">
            </div>
        </div>
        <div class="mt-2">
            <button type="submit" class="btn btn-primary">Filter</button>
            <a href="{{ route('nominatif-kredit.index') }}" class="btn btn-secondary">Reset</a>
        </div>
    </form>
    
    <!-- Data Table -->
    @if($data['success'])
        <table class="table table-striped">
            <thead>
                <tr>
                    <th>Nomor Rekening</th>
                    <th>Nama Nasabah</th>
                    <th>CAB</th>
                    <th>AO</th>
                    <th>Baki Debet</th>
                    <th>Tempat Bekerja</th>
                </tr>
            </thead>
            <tbody>
                @foreach($data['data'] as $kredit)
                <tr>
                    <td>{{ $kredit['NOMOR_REKENING'] }}</td>
                    <td>{{ $kredit['NAMA_NASABAH'] }}</td>
                    <td>{{ $kredit['CAB'] }}</td>
                    <td>{{ $kredit['AO'] }}</td>
                    <td>{{ number_format($kredit['BAKI_DEBET'], 2) }}</td>
                    <td>{{ $kredit['TEMPAT_BEKERJA'] }}</td>
                </tr>
                @endforeach
            </tbody>
        </table>
        
        <!-- Pagination -->
        @if($data['meta']['last_page'] > 1)
        <nav>
            <ul class="pagination">
                @for($i = 1; $i <= $data['meta']['last_page']; $i++)
                <li class="page-item {{ $i == $data['meta']['current_page'] ? 'active' : '' }}">
                    <a class="page-link" href="?page={{ $i }}">{{ $i }}</a>
                </li>
                @endfor
            </ul>
        </nav>
        @endif
    @else
        <div class="alert alert-danger">
            {{ $data['message'] }}
        </div>
    @endif
</div>
```

## 🛡️ Security Features

### **✅ Fitur Keamanan yang Diimplementasikan**
- ✅ **API Key Validation** - Setiap request harus menyertakan API key yang valid
- ✅ **Multiple Header Support** - Support X-API-Key dan Authorization header
- ✅ **Key Status Management** - API key bisa diaktifkan/dinonaktifkan
- ✅ **Request Tracking** - API key name dan description tersimpan di context
- ✅ **CORS Protection** - CORS header yang aman untuk frontend
- ✅ **Error Handling** - Error response yang informatif

### **🔒 Error Responses**
```json
// Missing API Key
{
  "success": false,
  "message": "API key is required. Please provide X-API-Key header or Authorization: ApiKey <key>",
  "error": "missing_api_key"
}

// Invalid API Key
{
  "success": false,
  "message": "Invalid API key",
  "error": "invalid_api_key"
}

// Disabled API Key
{
  "success": false,
  "message": "API key is disabled",
  "error": "disabled_api_key"
}
```

## 📊 Available Endpoints

| Method | Endpoint | Description | Requires API Key |
|--------|----------|-------------|------------------|
| GET | `/ping` | Health check | ❌ No |
| GET | `/health` | Database health | ❌ No |
| GET | `/api/validate` | Validate API key | ✅ Yes |
| GET | `/api/keys` | List API keys (masked) | ✅ Yes |
| GET | `/api/nominatif-kredit` | List data with filters | ✅ Yes |
| GET | `/api/nominatif-kredit/:id` | Get by ID | ✅ Yes |
| GET | `/api/nominatif-kredit/rekening/:nomor_rekening` | Get by rekening | ✅ Yes |

## 🧪 Testing Commands

```bash
# ✅ Test tanpa API key (should fail)
curl "http://localhost:8081/api/nominatif-kredit"

# ✅ Test dengan API key valid
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/nominatif-kredit?page=1&per_page=2"

# ✅ Test dengan Authorization header
curl -H "Authorization: ApiKey api-key-mobile-app-2025" \
     "http://localhost:8081/api/nominatif-kredit?cab=007"

# ✅ Test validasi API key
curl -H "X-API-Key: api-key-laravel-frontend-2025" \
     "http://localhost:8081/api/validate"

# ❌ Test dengan API key invalid
curl -H "X-API-Key: invalid-key-123" \
     "http://localhost:8081/api/nominatif-kredit"
```

## 🎯 Production Considerations

### **Environment Variables (.env)**
```env
# API Key Configuration
API_KEY_LARAVEL=api-key-laravel-frontend-prod-2025
API_KEY_MOBILE=api-key-mobile-app-prod-2025
API_KEY_ADMIN=api-key-dashboard-admin-prod-2025

# Server Configuration
PORT=8080
GIN_MODE=release
```

### **Deployment Tips**
1. **Change API Keys** - Generate strong, unique API keys untuk production
2. **Environment Variables** - Store API keys di environment variables
3. **Rate Limiting** - Pertimbangkan menambahkan rate limiting
4. **Logging** - Log semua API requests untuk monitoring
5. **HTTPS** - Selalu gunakan HTTPS di production

## 🎉 Summary

**✅ API Key Authentication berhasil diimplementasikan dengan fitur:**
- ✅ Simple API key validation tanpa perlu login
- ✅ Multiple API keys untuk different applications
- ✅ Support 2 format header (X-API-Key dan Authorization)
- ✅ API key management (activate/deactivate)
- ✅ Comprehensive error handling
- ✅ Laravel integration ready
- ✅ Production-ready security features

**🚀 Your API is now secured and ready for production use!**
