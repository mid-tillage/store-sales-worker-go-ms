# store-sales-worker-go-ms

## Table of Contents
- [Description](#description)
- [Installation](#installation)
- [Running the App](#running-the-app)
- [Test](#test)
- [Docker](#docker)
  - [Image Resource Usage Metrics](#image-resource-usage-metrics)
- [Kubernetes](#kubernetes)
  - [Pod Resource Usage Metrics](#pod-resource-usage-metrics)

## Description

Store's Admin Web Service example using [Nest](https://github.com/nestjs/nest) framework.

## Installation

```bash
$ go mod download
```

## Running the app
The following commands allow you to run the application

```bash
# development
go run .
```

### Swagger API documentation
You can access the Swagger documentation at: `http://localhost:8080/swagger/index.html`

## Docker

```bash
# Build Docker image
docker build -t store-sales-worker-go-ms:latest -f Dockerfile .

# Run Docker container (with example port mappings and environment variables)
docker run -e QUEUE_REDIS_IP="host.docker.internal" -e QUEUE_REDIS_PORT="6379" -e QUEUE_REDIS_PASSWORD="mypassword"  store-sales-worker-go-ms
```

### Image resource usage metrics

The table below shows resource usage metrics for the `store-sales-worker-go-ms` Docker container.

| REPOSITORY                  | TAG    | IMAGE ID      | CREATED    | SIZE    |
|-----------------------------|--------|---------------|------------|---------|
| store-sales-worker-go-ms    | latest | ea98a671f394  | 6 minutes  | 16.6MB  |


## Kubernetes

```bash
# Start Minikube to create a local Kubernetes cluster
minikube start

# Configure the shell to use Minikube's Docker daemon
& minikube -p minikube docker-env --shell powershell | Invoke-Expression

# Build Docker image with a specific tag and Dockerfile
docker build -t store-sales-worker-go-ms:latest -f Dockerfile .

# Apply Kubernetes configuration to create a pod
kubectl apply -f kubernetes/pod.yaml
```

### Pod resource usage metrics

The table below shows resource usage metrics for the `store-sales-worker-go-ms-pod` pod.

```bash
minikube addons enable metrics-server
kubectl top pods
```

**Note:** If you just enabled the metrics-server addon, remember to wait a couple of seconds before running the `kubectl top pods` command.


| NAME                          | CPU(cores)  | MEMORY(bytes) |
|-------------------------------|-------------|---------------|
| store-sales-worker-go-ms-pod  | 11m         | 6Mi           |
