import { runConsumer } from './amqp/consumer';
import { createServer } from './server';

const startApp = async (): Promise<void> => {
  try {
    // Start RabbitMQ consumer
    await runConsumer();
    console.log('Consumer is running...');

    // Start Express server
    const app = createServer();
    app.listen(3000, () => {
      console.log('Server is running on http://localhost:3000/latest');
    });
  } catch (error) {
    console.error('Failed to start application:', error);
  }
};

startApp(); 