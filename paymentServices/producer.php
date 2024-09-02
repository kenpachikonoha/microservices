<?php
require_once __DIR__ . '/vendor/autoload.php';
use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;


function producer() {

    $exchangeName = 'paymentExchange';
    $exchangeType = 'topic';
    $pattern = 'log.userAndProduct.data';
    
    $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
    $channel = $connection->channel();
    $channel->exchange_declare($exchangeName, $exchangeType, false, false, false);
    
    
    $msg = new AMQPMessage('Hello World!');
    $channel->basic_publish($msg, $exchangeName, $pattern);
    
    echo " [x] Sent 'Hello World!'\n";
    
    $channel->close();
    $connection->close();
}


?>