import express from 'express';
import { latestRoute } from './routes/latest';

export const createServer = (): express.Application => {
  const app = express();

  // Register routes
  app.use(latestRoute);

  return app;
};