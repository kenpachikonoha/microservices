// const amqp = require('amqplib/callback_api');


// amqp.connect('amqp://localhost:5555', (err, connection) => {
//     if (err) {
//         throw err;
//     }
//     connection.createChannel((err, channel) => {
//         if (err) {
//             throw err;
//         }
//         const queue = 'user_queue';
//         channel.assertQueue(queue, {
//             durable: false
//         });
//         console.log(`Waiting for messages in ${queue}`);
//         channel.consume(queue, (msg) => {
//             console.log(`Received ${msg.content.toString()}`);
//         });
//     });
// });