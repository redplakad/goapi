<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Http\JsonResponse;

class NominatifKreditController extends Controller
{
    private $apiBaseUrl = 'http://localhost:8080/api';

    /**
     * Get all nominatif kredit data with pagination
     */
    public function index(Request $request): JsonResponse
    {
        try {
            $response = Http::get($this->apiBaseUrl . '/nominatif-kredit', [
                'page' => $request->get('page', 1),
                'per_page' => $request->get('per_page', 10)
            ]);

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Failed to fetch data from Go API'
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API: ' . $e->getMessage()
            ], 500);
        }
    }

    /**
     * Get nominatif kredit by ID
     */
    public function show($id): JsonResponse
    {
        try {
            $response = Http::get($this->apiBaseUrl . '/nominatif-kredit/' . $id);

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Data not found'
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API: ' . $e->getMessage()
            ], 500);
        }
    }

    /**
     * Get nominatif kredit by nomor rekening
     */
    public function getByRekening($nomorRekening): JsonResponse
    {
        try {
            $response = Http::get($this->apiBaseUrl . '/nominatif-kredit/rekening/' . $nomorRekening);

            if ($response->successful()) {
                return response()->json($response->json());
            }

            return response()->json([
                'success' => false,
                'message' => 'Data not found for rekening: ' . $nomorRekening
            ], $response->status());

        } catch (\Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Error connecting to Go API: ' . $e->getMessage()
            ], 500);
        }
    }

    /**
     * Get nominatif kredit data for frontend view (example)
     */
    public function getForView(Request $request)
    {
        try {
            $response = Http::get($this->apiBaseUrl . '/nominatif-kredit', [
                'page' => $request->get('page', 1),
                'per_page' => $request->get('per_page', 20)
            ]);

            if ($response->successful()) {
                $data = $response->json();
                
                return view('nominatif-kredit.index', [
                    'kredits' => $data['data'] ?? [],
                    'meta' => $data['meta'] ?? [],
                    'success' => $data['success'] ?? false,
                    'message' => $data['message'] ?? ''
                ]);
            }

            return view('nominatif-kredit.index', [
                'kredits' => [],
                'meta' => [],
                'success' => false,
                'message' => 'Failed to fetch data from API'
            ]);

        } catch (\Exception $e) {
            return view('nominatif-kredit.index', [
                'kredits' => [],
                'meta' => [],
                'success' => false,
                'message' => 'Error: ' . $e->getMessage()
            ]);
        }
    }
}
