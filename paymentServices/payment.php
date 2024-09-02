<?php
header("Content-Type: application/json");
header("Access-Control-Allow-Methods: GET");
include 'consumer.php';

$method = $_SERVER['REQUEST_METHOD'];
$route = $_SERVER['REQUEST_URI'];
echo consumer();

if ($method == 'GET' && $route == '/payment.php') {
    $objec = '{"id":"123","itemName": "Condones","price": 200, "quantity": 2, "total": 400}';
    $response = json_decode($objec, true);
    echo json_encode($response);
} 

?>
