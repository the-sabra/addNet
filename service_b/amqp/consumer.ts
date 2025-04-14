import amqp from 'amqplib';
import { updateNumber } from '../save-file';

export const runConsumer = async (): Promise<void> => {
  const connection = await amqp.connect('amqp://localhost');
  const channel = await connection.createChannel();

  const handleMessage = (queue: string) => async (message: amqp.ConsumeMessage | null): Promise<void> => {
    if (message) {
      console.log(`Received message from ${queue}:`, message.content.toString());
      const intValue = parseInt(message.content.toString(), 10);

      if (queue === 'addition_queue') {
        await updateNumber(intValue);
      }

      channel.ack(message);
    }
  };

  await channel.assertQueue('addition_queue', { durable: false });
  await channel.consume('addition_queue', handleMessage('addition_queue'));

  console.log('Consumer is subscribed to queues: addition_queue');
};