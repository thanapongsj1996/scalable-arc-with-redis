import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
    vus: 1, // Number of virtual users (simulated users)
    duration: '10s', // Duration of the test
};

export default function () {
    const randomNumber = Math.floor(Math.random() * 100000) + 1
    let payload = {
        id: `${randomNumber}`,
    };

    let headers = {
        'Content-Type': 'application/json',
    };

    let res = http.post('http://localhost:8085/register-db', JSON.stringify(payload), { headers: headers });
    // let res = http.post('http://localhost:8085/register-redis', JSON.stringify(payload), { headers: headers });
    // let res = http.post('http://localhost:8085/register-buffer', JSON.stringify(payload), { headers: headers });

    check(res, {
        'Status is 200': (r) => r.status === 200,
    });
}
