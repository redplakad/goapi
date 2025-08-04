<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Http\JsonResponse;

class NominatifKreditController extends Controller
{
    private string $goApiBaseUrl;

    public function __construct()
    {
        // URL Go API server
        $this->goApiBaseUrl = config('services.go_api.base_url', 'http://localhost:8080/api');
    }

    /**
     * Get all nominatif kredit data with pagination and filtering
     */
    public function index(Request $request): JsonResponse
    {
        try {
            // Build query parameters
            $queryParams = [];
            
            // Pagination
            if ($request->has('page')) {
                $queryParams['page'] = $request->get('page', 1);
            }
            if ($request->has('per_page')) {
                $queryParams['per_page'] = min($request->get('per_page', 10), 100);
            }
            
            // Filters
            if ($request->filled('cab')) {
                $queryParams['cab'] = $request->get('cab');
            }
            if ($request->filled('ao')) {
                $queryParams['ao'] = $request->get('ao');
            }
            if ($request->filled('ket_kd_prd')) {
                $queryParams['ket_kd_prd'] = $request->get('ket_kd_prd');
            }
            if ($request->filled('tempat_bekerja')) {
                $queryParams['tempat_bekerja'] = $request->get('tempat_bekerja');
            }

            // Make request to Go API
            $response = Http::timeout(30)->get("{$this->goApiBaseUrl}/nominatif-kredit", $queryParams);

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Failed to fetch data from Go API',
                'error' => $response->body()
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API',
                'error' => $e->getMessage()
            ], 500);
        }
    }

    /**
     * Get nominatif kredit by ID
     */
    public function show(int $id): JsonResponse
    {
        try {
            $response = Http::timeout(30)->get("{$this->goApiBaseUrl}/nominatif-kredit/{$id}");

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Failed to fetch data from Go API',
                'error' => $response->body()
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API',
                'error' => $e->getMessage()
            ], 500);
        }
    }

    /**
     * Get nominatif kredit by nomor rekening
     */
    public function getByRekening(string $nomorRekening): JsonResponse
    {
        try {
            $response = Http::timeout(30)->get("{$this->goApiBaseUrl}/nominatif-kredit/rekening/{$nomorRekening}");

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Failed to fetch data from Go API',
                'error' => $response->body()
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API',
                'error' => $e->getMessage()
            ], 500);
        }
    }
}

// Web Controller untuk menampilkan view
namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\View\View;

class NominatifKreditController extends Controller
{
    private string $goApiBaseUrl;

    public function __construct()
    {
        $this->goApiBaseUrl = config('services.go_api.base_url', 'http://localhost:8080/api');
    }

    /**
     * Display list of nominatif kredit with filters
     */
    public function index(Request $request): View
    {
        try {
            // Build query parameters dari request
            $queryParams = [
                'page' => $request->get('page', 1),
                'per_page' => $request->get('per_page', 20)
            ];
            
            // Add filters if provided
            $filters = ['cab', 'ao', 'ket_kd_prd', 'tempat_bekerja'];
            foreach ($filters as $filter) {
                if ($request->filled($filter)) {
                    $queryParams[$filter] = $request->get($filter);
                }
            }

            $response = Http::timeout(30)->get("{$this->goApiBaseUrl}/nominatif-kredit", $queryParams);

            if ($response->successful()) {
                $data = $response->json();
                
                return view('nominatif-kredit.index', [
                    'success' => $data['success'],
                    'message' => $data['message'],
                    'kredits' => $data['data'] ?? [],
                    'meta' => $data['meta'] ?? null,
                    'filters' => $request->only($filters)
                ]);
            }

            return view('nominatif-kredit.index', [
                'success' => false,
                'message' => 'Failed to fetch data from Go API',
                'kredits' => [],
                'meta' => null,
                'filters' => $request->only($filters)
            ]);

        } catch (\Exception $e) {
            return view('nominatif-kredit.index', [
                'success' => false,
                'message' => 'Error connecting to Go API: ' . $e->getMessage(),
                'kredits' => [],
                'meta' => null,
                'filters' => $request->only(['cab', 'ao', 'ket_kd_prd', 'tempat_bekerja'])
            ]);
        }
    }

    /**
     * Display specific nominatif kredit by rekening
     */
    public function getByRekening(string $nomorRekening): View
    {
        try {
            $response = Http::timeout(30)->get("{$this->goApiBaseUrl}/nominatif-kredit/rekening/{$nomorRekening}");

            if ($response->successful()) {
                $data = $response->json();
                
                return view('nominatif-kredit.detail', [
                    'success' => $data['success'],
                    'message' => $data['message'],
                    'kredit' => $data['data'] ?? null
                ]);
            }

            return view('nominatif-kredit.detail', [
                'success' => false,
                'message' => 'Failed to fetch data from Go API',
                'kredit' => null
            ]);

        } catch (\Exception $e) {
            return view('nominatif-kredit.detail', [
                'success' => false,
                'message' => 'Error connecting to Go API: ' . $e->getMessage(),
                'kredit' => null
            ]);
        }
    }
}
