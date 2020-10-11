# Server Benchmarks

## Description

A repository testing various concepts on how to set up kubernetes, and run load testing or benchmarks.

Containing Concepts: 
- Benchmarking locally, on docker or on minikube cluster.
- Load testing locally, on docker or on minikube cluster.
- Node affinity. 
- Node draining and Pod Disruption Budget.
- Horizontal Pod Autoscaler.
- Node taints

## Useful Commands

- `docker login --username <username>`
- `docker build . -t gkabitakis/<serserver-typever>-server:<version>`
- `docker push gkabitakis/<server-type>-server:<version>`
- `docker run -p 5000:5000 -- rm --name <server-type>-server <server-type>-server`
- `docker stop <server-type>-server`
- `docker stats <server-type>-server`
- `minikube start --nodes 3 -p multinod --cpus 4`
- `kubectl drain <node-name> --grace-period=30 --ignore-daemonsets=true` Drain a specific node or you can 
specify which pods from a node to drain by running `--pod-selector=''` and a label.
- `kubectl uncordon <node-name>` After draining you must set node again to accept pods to be scheduled
- `kubectl taint nodes node1 key=value:NoSchedule`


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
- [vegeta](https://github.com/tsenart/vegeta)
```bash
# Example
  echo "GET http://localhost:5000/listen?test1=1&test=2" | vegeta attack -duration=100s -rate=1000 | tee results.bin | vegeta report
```

## How to run

You need to create a namespace for each server, its easier to maintain it like this

You also need to label nodes if you want to test pod affinity,node selectors and drain node. Here I use `node.type` label and assign master and slave values.
You can achieve that by running `kubectl label nodes <node-name> <label-key>=<label-value>`.

You need to also have metrics server enabled. When running in `minikube` you can just enable it by running `minikube addons enable metrics-server` or you can [install](https://github.com/kubernetes-sigs/metrics-server) it in your cluster.

## Kubernetes apply taint to node

You can apply a taint to a node. Taints are the opposite of Node affinities, they allow to repel a set of pods.

If you want to apply `NoExecute` taint be careful that when you apply this taint it evicts pod that don't respect it. So if you want to apply this type of taint make sure first draining your node.

- `kubectl drain <node-name> --grace-period=40 --ignore-daemonsets=true --pod-selector='<key-label>=<key-value>'`
- Add taint
```bash
kubectl taint nodes <node-name>  key=value:TaintEffect
# Taint effect can be `NoSchedule`, `NoExecute` or `PreferNoSchedule`
```
- Remove taint
```bash
kubectl taint nodes <node-name>  key=value:TaintEffect-
```