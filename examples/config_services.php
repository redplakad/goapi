<?php

// config/services.php - Add this to your existing services config

'go_api' => [
    'base_url' => env('GO_API_BASE_URL', 'http://localhost:8080/api'),
    'timeout' => env('GO_API_TIMEOUT', 30),
],
