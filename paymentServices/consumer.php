<?php
include 'producer.php';
require_once __DIR__ . '/vendor/autoload.php';
use PhpAmqpLib\Connection\AMQPStreamConnection;

function consumer() {
    
    $exchangeName = 'cartExchange';
    $exchangeType = 'direct';
    $queue = 'paymentQueue';
    $pattern = 'fromCart';
    
    
    $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
    $channel = $connection->channel();
    
    $channel->queue_declare($queue, false, false, false, false);
    $channel->exchange_declare($exchangeName, $exchangeType, false, false, false);
    $channel->queue_bind($queue, $exchangeName, $pattern);
    
    echo " [*] Waiting for messages. To exit press CTRL+C\n";
    
    $callback = function ($msg) {
        echo ' [x] Received ', $msg->body, "\n";
        producer();
      };
      
      $channel->basic_consume($queue, $pattern, false, true, false, false, $callback);
      
      try {
          $channel->consume();
      } catch (\Throwable $exception) {
          echo $exception->getMessage();
      }

}
?>