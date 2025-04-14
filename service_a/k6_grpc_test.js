import grpc from 'k6/net/grpc';
import { check } from 'k6';

const client = new grpc.Client();
client.load(['./internal/proto'], 'addition.proto');
export const options = {
    scenarios: {
        basic: {
            executor: 'constant-vus',
            vus: 10,
            duration: '30s',
        },
        load: {
            executor: 'ramping-vus',
            startVUs: 0,
            stages: [
                { duration: '20s', target: 25 },
                { duration: '30s', target: 25 },
                { duration: '10s', target: 0 },
            ],
            gracefulRampDown: '5s',
        },
    },
    thresholds: {
        'checks': ['rate>0.95'],
        'http_req_duration': ['p(95)<500'],
    },
};
export default () => {
  client.connect('localhost:50051', { plaintext: true });

  const data = { a: 5, b: 3 };
  const response = client.invoke('addition.AdditionService/Add', data);

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
    'result is correct': (r) => r && r.message.result === 8,
  });

  client.close();
};