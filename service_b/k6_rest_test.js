import http from 'k6/http';
import { check } from 'k6';
import { sleep } from 'k6';
// Options for the test scenario
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
        http_req_duration: ['p(95)<500'], // 95% of requests should be below 500ms
        http_req_failed: ['rate<0.01'],   // Less than 1% of requests should fail
    },
};
export default () => {
  const url = 'http://localhost:3000/latest';
  const response = http.get(url);

  check(response, {
    'status is 200': (r) => r.status === 200,
    'response has latestValue': (r) => JSON.parse(r.body).hasOwnProperty('latestValue'),
  });
};