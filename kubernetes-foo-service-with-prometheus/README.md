# Kubernetes Application Deployment Guide

This repository contains Kubernetes configurations for deploying a Go application (foo-server) with Prometheus monitoring.

## Prerequisites

- Docker
- kubectl
- Minikube

## 1. Install Minikube

```bash
brew install minikube
```

## 2. Start Minikube

```bash
minikube start
```

## 3. Create Persistent Volumes

Create the required directories inside the Minikube VM for persistent storage:

```bash
minikube ssh "sudo mkdir -p /tmp/prometheus && sudo chmod 777 /tmp/prometheus"
```

## 4. Build and Load Docker Image

```bash
docker build -t foo-server:1.0.1 .
```

Load the image into Minikube
```
minikube image load foo-server:1.0.1
```
## 5. Deploy All Resources

```bash
kubectl apply -f kubernetes/ --recursive
```
## 8. Access Applications
```bash
kubectl port-forward -n monitoring svc/prometheus 9090:9090
```
