import json
import pika
import sys
import os
from flask import Flask
from waitress import serve

def main():
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()

    exchange_name = 'paymentExchange'
    exchange_type = 'topic'
    queue = 'product'
    pattern = 'log.userAndProduct.#'

    channel.exchange_declare(exchange=exchange_name, exchange_type=exchange_type)
    channel.queue_declare(queue=queue)
    channel.queue_bind(exchange=exchange_name, queue=queue, routing_key=pattern)

    def callback(ch, method, properties, body):
        print(f" [x] Received {body}")

    channel.basic_consume(queue=queue, on_message_callback=callback, auto_ack=True)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()

app = Flask(__name__)

@app.route('/products')
def products():
    value = {
        "products": [
            {"id": 1, "name": "Lubricante", "price": 100},
            {"id": 2, "name": "Troyano", "price": 150},
            {"id": 3, "name": "Azulita", "price": 125},
            {"id": 4, "name": "Condones", "price": 200}
        ]
    }
    return json.dumps(value)

def create_app():
    return app

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
    
    serve(app, host='127.0.0.1', port=5095)
