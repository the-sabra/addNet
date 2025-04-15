import promClient from 'prom-client';

export type Payload = {
    value: number;
    sent_at: string;
}

const queueLatency = new promClient.Histogram({
    name: 'queue_delivery_latency_ms',
    help: 'Time between message sent and consumer received',
    buckets: [100, 200, 300, 400, 500, 600, 700, 800, 900, 1000],
  });

const requestCount = new promClient.Counter({
    name: 'example_requests_count',
    help: 'Request counter per method, url and status code',
    labelNames: ['method', 'url', 'status'],
  });
  
const latencyHistogram = new promClient.Histogram({
    name: 'example_request_latency_ms',
    help: 'Histogram for request latency in millisecond',
    labelNames: ['method', 'url'],
    buckets: [100, 200, 300, 400, 500, 600, 700, 800, 900, 1000],
  });
  
promClient.collectDefaultMetrics();
promClient.register.registerMetric(queueLatency);
promClient.register.registerMetric(requestCount);
promClient.register.registerMetric(latencyHistogram);

export const middleWareMetrics = (req: any, res: any, next: any) => {
    const start = Date.now();

    res.on('finish', () => {
      const duration = Date.now() - start;
  
      requestCount.inc({
        method: req.method,
        url: req.route ? req.route.path : req.path,
        status: res.statusCode.toString(),
      });
  
      latencyHistogram.observe({
        method: req.method,
        url: req.route ? req.route.path : req.path,
      }, duration);
    });
  
    next();
}

export function observeLatencyMetrics(payload: Payload) {
    const sentAt = new Date(payload.sent_at);
    const latency = Date.now() - sentAt.getTime();
    queueLatency.observe(latency);
  
    console.log(`[Consumer] result=${payload.value}, latency=${latency}ms`);
} 

export default promClient;