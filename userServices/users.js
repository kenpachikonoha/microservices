const http = require('http');
const amqp = require('amqplib');


const exchangeName = 'paymentExchange';
const exchangeType = 'topic';
const queue = 'user';
const pattern = 'log.userAndProduct.#';

async function rabbitMqTopic(exchangeNam, exchangeTyp, queu, patter) {
    try {
        const connection = await amqp.connect('amqp://localhost//');
        const channel = await connection.createChannel();

        await channel.assertQueue(queu,{durable: true});
        await channel.assertExchange(exchangeNam, exchangeTyp, { durable: false });
        await channel.bindQueue(queu, exchangeNam, patter);

        console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", queu);
        channel.consume(queu, (msg) => {
            console.log(" [x] Received %s", msg.content);

            channel.ack(msg);
            // return content;
        });
    } catch (error) {
        console.log(error); 
    }
}



const user = {
    name: 'Toji Senn`in',
    age: 26,
    nickname: 'El chambeador',
    hobbies: ['Hacer pasar verguenza al personaje mas fuerte del anime', 'chambear', 'ser el mejor'],
}

const server = http.createServer((req, res) => {

    if (req.url === '/user') {
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify(user));
    }

});

rabbitMqTopic(exchangeName, exchangeType, queue, pattern);
server.listen(3030); 