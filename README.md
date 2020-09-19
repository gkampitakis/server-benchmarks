# Server Benchmarks

## Description

## Useful Commands

- `docker login --username <username>`
- `docker build . -t gkabitakis/<server>-server:<version>`
- `docker push gkabitakis/<server>-server:<version>`
- `minikube start --nodes 2 -p multinode-demo`

Create docker-hub secret by running 

- `kubectl create secret docker-registry docker-hub --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>`

## Benchmarking tools

- [autocannon](https://www.npmjs.com/package/autocannon)
```bash
# Example
autocannon -c 1000 -p 10 -d 120 http://localhost:8080/listen\?test\=1\&test2\=2
```