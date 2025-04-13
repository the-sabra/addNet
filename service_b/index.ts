import amqp from 'amqplib';
import { readFile, updateNumber } from './save-file';
import express from 'express';

const runConsumer = async (): Promise<void> => {
  const connection = await amqp.connect('amqp://localhost');
  const channel = await connection.createChannel();

  const handleMessage = (queue: string) => async (message: amqp.ConsumeMessage | null): Promise<void> => {
    if (message) {
      console.log(`Received message from ${queue}:`, message.content.toString());
      
      const intValue = parseInt(message.content.toString(), 10);

      if (queue === 'addition_queue') {
        await updateNumber(intValue)
      }

      channel.ack(message);
    }
  };


  // Subscribing to sms-queue
  await channel.assertQueue('addition_queue', { durable: false });
  await channel.consume('addition_queue', handleMessage('addition_queue'));

  console.log('Consumer is subscribed to queues: addition_queue');
};

// Express server setup
const app = express();

app.get("/latest", async (req, res) => {
  try {
    const latestValue = await readFile();
    res.status(200).json({ latestValue });
  }
  catch (error) {
    console.error('Error reading file:', error);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

runConsumer()
  .then(() => {
    console.log('Consumer is running...');
  })
  .catch((error) => {
    console.error('Failed to run RabbitMQ consumer', error);
  });

app.listen(3000, () => {
  console.log('Server is running on http://localhost:3000/latest');
});