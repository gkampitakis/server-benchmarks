# Server Benchmarks

## Description

A repository testing various concepts on how to set up kubernetes, and run load testing or benchmarks.

Containing Concepts: 
- Benchmarking locally, on docker or on minikube cluster.
- Load testing locally, on docker or on minikube cluster.
- Node affinity. 
- Node draining and Pod Disruption Budget.
- Horizontal Pod Autoscaler.

## Useful Commands

- `docker login --username <username>`
- `docker build . -t gkabitakis/<serserver-typever>-server:<version>`
- `docker push gkabitakis/<server-type>-server:<version>`
- `docker run -p 5000:5000 -- rm --name <server-type>-server <server-type>-server`
- `docker stop <server-type>-server`
- `docker stats <server-type>-server`
- `minikube start --nodes 3 -p multinod --cpus 4`
- `kubectl drain node <node-name> --grace-period=30 --ignore-daemonsets=true` Drain a specific node or you can 
specify which pods from a node to drain by running `--pod-selector=''` and a label.
- `kubectl uncordeon <node-name>` After draining you must set node again to accept pods to be scheduled


Create docker-hub secret by running (No longer needed)

- `kubectl create secret docker-registry docker-hub --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>`

## Containing

- Fastify server Dockerized and manifest files for kubernetes.
- Go server Dockerized and manifest files for kubernetes.

## Benchmarking tools

- [autocannon](https://www.npmjs.com/package/autocannon)
```bash
# Example
autocannon -c 1000 -p 10 -d 120 {URL}
```

## How to run

You need to create a namespace for each server, its easier to maintain it like this
