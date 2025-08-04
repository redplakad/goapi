# API Documentation - Nominatif Kredit

## Base URL
```
http://localhost:8080/api
```

## Endpoints

### 1. Get All Nominatif Kredit (with Pagination)
```
GET /api/nominatif-kredit
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `per_page` (optional): Items per page (default: 10, max: 100)

**Example Request:**
```
GET /api/nominatif-kredit?page=1&per_page=20
```

**Response:**
```json
{
  "success": true,
  "message": "Nominatif kredit data retrieved successfully",
  "data": [
    {
      "id": 1,
      "DATADATE": "2024-01-01T00:00:00Z",
      "CAB": "001",
      "NOMOR_REKENING": "123456789",
      "NO_CIF": "CIF001",
      "NAMA_NASABAH": "John Doe",
      // ... other fields
    }
  ],
  "meta": {
    "current_page": 1,
    "per_page": 20,
    "total": 100,
    "last_page": 5
  }
}
```

### 2. Get Nominatif Kredit by ID
```
GET /api/nominatif-kredit/{id}
```

**Example Request:**
```
GET /api/nominatif-kredit/1
```

**Response:**
```json
{
  "success": true,
  "message": "Nominatif kredit data retrieved successfully",
  "data": {
    "id": 1,
    "DATADATE": "2024-01-01T00:00:00Z",
    "CAB": "001",
    "NOMOR_REKENING": "123456789",
    // ... other fields
  }
}
```

### 3. Get Nominatif Kredit by Nomor Rekening
```
GET /api/nominatif-kredit/rekening/{nomor_rekening}
```

**Example Request:**
```
GET /api/nominatif-kredit/rekening/123456789
```

**Response:**
```json
{
  "success": true,
  "message": "Nominatif kredit data retrieved successfully",
  "data": [
    {
      "id": 1,
      "DATADATE": "2024-01-01T00:00:00Z",
      "CAB": "001",
      "NOMOR_REKENING": "123456789",
      // ... other fields
    }
  ]
}
```

## Error Responses

**Bad Request (400):**
```json
{
  "success": false,
  "message": "Invalid ID format"
}
```

**Not Found (404):**
```json
{
  "success": false,
  "message": "Nominatif kredit not found"
}
```

**Internal Server Error (500):**
```json
{
  "success": false,
  "message": "Failed to fetch nominatif kredit data"
}
```

## Fields Description

The API returns all fields from the `nominatif_kredit` table:

- `id`: Primary key
- `DATADATE`: Data date
- `CAB`: Branch code  
- `NOMOR_REKENING`: Account number
- `NO_CIF`: Customer identification number
- `NAMA_NASABAH`: Customer name
- `ALAMAT`: Address
- `KODE_KOLEK`: Collection code
- `PLAFOND_AWAL`: Initial ceiling (decimal)
- `BAKI_DEBET`: Debit balance (decimal)
- `TUNGGAKAN_POKOK`: Principal arrears (decimal)
- `TUNGGAKAN_BUNGA`: Interest arrears (decimal)
- And many more fields as defined in the Laravel model...

## Usage in Laravel Frontend

To consume this API in your Laravel frontend, you can use HTTP client like Guzzle or Laravel's HTTP facade:

```php
use Illuminate\Support\Facades\Http;

// Get all nominatif kredit with pagination
$response = Http::get('http://localhost:8080/api/nominatif-kredit', [
    'page' => 1,
    'per_page' => 20
]);

$data = $response->json();

// Get by ID
$response = Http::get('http://localhost:8080/api/nominatif-kredit/1');
$kredit = $response->json();

// Get by nomor rekening
$response = Http::get('http://localhost:8080/api/nominatif-kredit/rekening/123456789');
$kredits = $response->json();
```
