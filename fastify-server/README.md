# Fastify Server

## Description

Fastify server

## Useful Commands

-   `kubectl port-forward svc/fastify-server 8080:80`
-   `kubectl edit hpa fastify-hpa -n fastify-server` If you want to edit hpa.

## Run benchmarks

### Locally

Start the service by installing dependencies `npm i` and then run the server with `npm run start`.

### Docker

Build image and run the container, after you can start benchmarking or load testing.

-   You can check the container stats with `docker stats <container-name or id>`

### Kubernetes

Spin up a `node-box` pod and connect inside by running `kubectl exec -it node-box -n fastify-server -- bash`.

Install autocannon inside the pod with `npm i autocannon -g` and then you can start using the autocannon by hitting fastify-server in `http://fastify-server.fastify-server/` and the route you want to reach.

### Results locally inside docker

`autocannon -c 1000 -p 10 -d 120 -m 'POST' localhost:5000/postdata`

```bash
┌─────────┬────────┬─────────┬─────────┬─────────┬────────────┬───────────┬─────────┐
│ Stat    │ 2.5%   │ 50%     │ 97.5%   │ 99%     │ Avg        │ Stdev     │ Max     │
├─────────┼────────┼─────────┼─────────┼─────────┼────────────┼───────────┼─────────┤
│ Latency │ 980 ms │ 1208 ms │ 2210 ms │ 2662 ms │ 1313.02 ms │ 350.84 ms │ 7995 ms │
└─────────┴────────┴─────────┴─────────┴─────────┴────────────┴───────────┴─────────┘
┌───────────┬────────┬────────┬─────────┬─────────┬─────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%     │ 97.5%   │ Avg     │ Stdev   │ Min    │
├───────────┼────────┼────────┼─────────┼─────────┼─────────┼─────────┼────────┤
│ Req/Sec   │ 1482   │ 1693   │ 7907    │ 9799    │ 7577.55 │ 1857.49 │ 1378   │
├───────────┼────────┼────────┼─────────┼─────────┼─────────┼─────────┼────────┤
│ Bytes/Sec │ 295 kB │ 337 kB │ 1.57 MB │ 1.95 MB │ 1.51 MB │ 370 kB  │ 274 kB │
└───────────┴────────┴────────┴─────────┴─────────┴─────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

909k requests in 120.28s, 181 MB read
```
