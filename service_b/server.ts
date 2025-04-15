import express from 'express';
import { latestRoute } from './routes/latest';
import { prometheusRoute } from './routes/prometheus';
import { middleWareMetrics } from './prometheus/client';

export const createServer = (): express.Application => {
  const app = express();

  // Middleware Metrics
  app.use(middleWareMetrics);
  // Register routes
  app.use(latestRoute);
  app.use(prometheusRoute);

  return app;
};