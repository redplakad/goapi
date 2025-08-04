<?php

// routes/api.php
use App\Http\Controllers\Api\NominatifKreditController;

Route::prefix('v1')->group(function () {
    // API routes untuk consume Go API
    Route::get('/nominatif-kredit', [NominatifKreditController::class, 'index']);
    Route::get('/nominatif-kredit/{id}', [NominatifKreditController::class, 'show']);
    Route::get('/nominatif-kredit/rekening/{nomor_rekening}', [NominatifKreditController::class, 'getByRekening']);
});

// routes/web.php
use App\Http\Controllers\Web\NominatifKreditController as WebNominatifKreditController;

Route::prefix('nominatif-kredit')->group(function () {
    // Web routes untuk menampilkan view
    Route::get('/', [WebNominatifKreditController::class, 'index'])->name('nominatif-kredit.index');
    Route::get('/rekening/{nomor_rekening}', [WebNominatifKreditController::class, 'getByRekening'])->name('nominatif-kredit.rekening');
});

/*
Usage examples in Laravel:

1. API Endpoint:
   GET /api/v1/nominatif-kredit?page=1&per_page=20
   GET /api/v1/nominatif-kredit?cab=007&ao=JAENUDIN&page=1&per_page=10
   GET /api/v1/nominatif-kredit?ket_kd_prd=HONORER&tempat_bekerja=SDN
   GET /api/v1/nominatif-kredit/1
   GET /api/v1/nominatif-kredit/rekening/123456789

2. Web Route:
   GET /nominatif-kredit (akan menampilkan view dengan data dari Go API)

3. JavaScript/AJAX call with filters:
   
   // Basic request
   fetch('/api/v1/nominatif-kredit?page=1&per_page=10')
   .then(response => response.json())
   .then(data => {
       if (data.success) {
           console.log('Data:', data.data);
           console.log('Meta:', data.meta);
       }
   });

   // With filters
   const filters = {
       cab: '007',
       ao: 'JAENUDIN',
       ket_kd_prd: 'HONORER',
       tempat_bekerja: 'SDN',
       page: 1,
       per_page: 20
   };
   
   const queryString = new URLSearchParams(filters).toString();
   fetch(`/api/v1/nominatif-kredit?${queryString}`)
   .then(response => response.json())
   .then(data => {
       console.log('Filtered data:', data.data);
   });

4. Dalam Blade template dengan form filter:
   
   <form method="GET" action="{{ route('nominatif-kredit.index') }}">
       <div class="row">
           <div class="col-md-3">
               <input type="text" name="cab" placeholder="CAB" value="{{ request('cab') }}" class="form-control">
           </div>
           <div class="col-md-3">
               <input type="text" name="ao" placeholder="AO" value="{{ request('ao') }}" class="form-control">
           </div>
           <div class="col-md-3">
               <input type="text" name="ket_kd_prd" placeholder="Product Description" value="{{ request('ket_kd_prd') }}" class="form-control">
           </div>
           <div class="col-md-3">
               <input type="text" name="tempat_bekerja" placeholder="Workplace" value="{{ request('tempat_bekerja') }}" class="form-control">
           </div>
       </div>
       <button type="submit" class="btn btn-primary mt-2">Filter</button>
       <a href="{{ route('nominatif-kredit.index') }}" class="btn btn-secondary mt-2">Reset</a>
   </form>
   
   @if($success)
       @foreach($kredits as $kredit)
       <tr>
           <td>{{ $kredit['NOMOR_REKENING'] }}</td>
           <td>{{ $kredit['NAMA_NASABAH'] }}</td>
           <td>{{ number_format($kredit['BAKI_DEBET'], 2) }}</td>
       </tr>
       @endforeach
       
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
   @else
       <p>{{ $message }}</p>
   @endif
*/
