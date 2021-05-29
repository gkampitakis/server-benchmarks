# Go Server

## Description

Fastify server

## Useful Commands

## Run benchmarks

### Locally

Start the service by running `make all`.

### Docker

Build image and run the container, after you can start benchmarking or load testing.

-   You can check the container stats with `docker stats <container-name or id>`

### Kubernetes

Spin up a `node-box` pod and connect inside by running `kubectl exec -it node-box -n fastify-server -- bash`.

Install autocannon inside the pod with `npm i autocannon -g` and then you can start using the autocannon by hitting fastify-server in `http://fastify-server.fastify-server/` and the route you want to reach.

### Results locally inside docker

`autocannon -c 1000 -p 10 -d 120 -m 'POST' localhost:5000/postdata`

```bash
┌─────────┬────────┬────────┬─────────┬─────────┬───────────┬───────────┬─────────┐
│ Stat    │ 2.5%   │ 50%    │ 97.5%   │ 99%     │ Avg       │ Stdev     │ Max     │
├─────────┼────────┼────────┼─────────┼─────────┼───────────┼───────────┼─────────┤
│ Latency │ 127 ms │ 479 ms │ 2051 ms │ 2550 ms │ 579.05 ms │ 449.07 ms │ 5783 ms │
└─────────┴────────┴────────┴─────────┴─────────┴───────────┴───────────┴─────────┘
┌───────────┬────────┬────────┬─────────┬───────┬──────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%     │ 97.5% │ Avg      │ Stdev   │ Min    │
├───────────┼────────┼────────┼─────────┼───────┼──────────┼─────────┼────────┤
│ Req/Sec   │ 6251   │ 6943   │ 17743   │ 21871 │ 17347.05 │ 2930.42 │ 6090   │
├───────────┼────────┼────────┼─────────┼───────┼──────────┼─────────┼────────┤
│ Bytes/Sec │ 857 kB │ 951 kB │ 2.43 MB │ 3 MB  │ 2.38 MB  │ 401 kB  │ 834 kB │
└───────────┴────────┴────────┴─────────┴───────┴──────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

2082k requests in 120.85s, 285 MB read
2 errors (0 timeouts)
```
