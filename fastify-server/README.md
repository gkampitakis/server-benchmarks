# Fastify Server

## Description

Fastify server

## Useful Commands

- `kubectl port-forward svc/fastify-server 8080:80`
- `kubectl edit hpa fastify-hpa -n fastify-server` If you want to edit hpa.

## Run benchmarks

### Locally

Start the service by installing dependencies `npm i` and then run the server with `npm run start`.

### Docker

Build image and run the container, after you can start benchmarking or load testing.

- You can check the container stats with `docker stats <container-name or id>`

### Kubernetes

Spin up a `node-box` pod and connect inside by running `kubectl exec -it node-box -n fastify-server -- bash`.

Install autocannon inside the pod with `npm i autocannon -g` and then you can start using the autocannon by hitting fastify-server in `http://fastify-server.fastify-server/` and the route you want to reach.
