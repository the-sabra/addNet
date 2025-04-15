import { runConsumer } from './amqp/consumer';
import { createServer } from './server';

const PORT =  process.env.PORT || 2113; // Default to 2113 if PORT is not set

const startApp = async (): Promise<void> => {
  try {
    // Start RabbitMQ consumer
    await runConsumer();
    console.log('Consumer is running...');

    // Start Express server
    const app = createServer();
    app.listen(PORT, () => {
      console.log(`Server is running on http://localhost:${PORT}/latest`);
    });
  } catch (error) {
    console.error('Failed to start application:', error);
  }
};

startApp(); 