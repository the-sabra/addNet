import amqp from 'amqplib';
import { updateNumber } from '../save-file';
import { observeLatencyMetrics, type Payload } from '../prometheus/client';

export const runConsumer = async (): Promise<void> => {
  const connection = await amqp.connect(process.env.RABBITMQ_URL || 'amqp://localhost:5672');
  const channel = await connection.createChannel();

  const handleMessage = (queue: string) => async (message: amqp.ConsumeMessage | null): Promise<void> => {
    if (message) {
      const content = JSON.parse(message.content.toString()) as Payload;
      if (queue === 'addition_queue') {
        observeLatencyMetrics(content)
        await updateNumber(content);
      }

      channel.ack(message);
    }
  };

  await channel.assertQueue('addition_queue', { durable: false });
  await channel.consume('addition_queue', handleMessage('addition_queue'));

  console.log('Consumer is subscribed to queues: addition_queue');
};