# Maktaba API

A Go-based API service with PostgreSQL database running on Kubernetes.

### Prerequisites

- Minikube
- kubectl
- Docker

### Deployment Steps

1. Start Minikube:

```bash
minikube start
```

2. Enable Minikube addons:

```bash
minikube addons enable ingress
```

3. Point shell to minikube's docker-daemon:

```bash
eval $(minikube docker-env)
```

4. Build the Docker image:

```bash
docker build -t maktaba-api:latest .
```

5. Create Kubernetes resources:

```bash
# Create ConfigMap from migrations folder
kubectl create configmap migrations-config --from-file=../migrations/

# Create ConfigMap and Secret
kubectl apply -f k8s/postgres-config.yaml
kubectl apply -f k8s/postgres-secret.yaml

# Create PostgreSQL deployment and service
kubectl apply -f k8s/postgres-deployment.yaml
kubectl apply -f k8s/postgres-service.yaml

# Create API deployment and service
kubectl apply -f k8s/maktaba-api-deployment.yaml
kubectl apply -f k8s/maktaba-api-service.yaml
```

6. Verify deployment:

```bash
kubectl get pods
kubectl get services
kubectl get deployments
```

7. Access the application:

```bash
minikube service maktaba-api-service --url
```

### Useful Kubernetes Commands

```bash
# View logs
kubectl logs deployment/maktaba-api-deployment

# Scale deployment
kubectl scale deployment maktaba-api-deployment --replicas=3

# Describe resources
kubectl describe pod <pod-name>
kubectl describe service maktaba-api-service

# Delete resources
kubectl delete -f k8s/
```
