# 📚 Maktaba - Book Management API

## 🌟 Overview

Maktaba is a robust, production-ready Book Management API built with Go, offering comprehensive features for book catalog management, user authentication, and system integration.

## 🚀 Features

- **Book Management**

  - Create, read, update, and delete book entries
  - Flexible book information storage
  - Advanced filtering and search capabilities

- **User Authentication**

  - Secure user registration
  - Token-based authentication
  - Role-based access control

- **Performance & Security**
  - Rate limiting
  - CORS support
  - Connection pooling for database
  - Environment-specific configurations

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Authentication**: Statfull token based auth

## 📡 API Endpoints

### Books

| Method | Endpoint        | Description            | Permission    |
| ------ | --------------- | ---------------------- | ------------- |
| GET    | `/v1/books`     | List all books         | `books:read`  |
| POST   | `/v1/books`     | Create a new book      | `books:write` |
| GET    | `/v1/books/:id` | Retrieve specific book | `books:read`  |
| PATCH  | `/v1/books/:id` | Update a book          | `books:write` |
| DELETE | `/v1/books/:id` | Delete a book          | `books:write` |

### Authentication

| Method | Endpoint                    | Description           |
| ------ | --------------------------- | --------------------- |
| POST   | `/v1/users`                 | Register new user     |
| PUT    | `/v1/users/activated`       | Activate user account |
| POST   | `/v1/tokens/authentication` | Generate auth token   |

## 🔧 Configuration

### Environment Variables

| Variable   | Description                  | Default       |
| ---------- | ---------------------------- | ------------- |
| `DB_DSN`   | PostgreSQL connection string | Required      |
| `PORT`     | API server port              | `4000`        |
| `ENV_MODE` | Environment mode             | `development` |

### Command Line Flags

```bash
go run ./cmd/api -db-dsn="postgres://maktaba:maktaba@localhost:5432/maktaba?sslmode=disable" -port=4000

```

## 📦 Installation

### Prerequisites

- Go 1.23+
- PostgreSQL 12+
- Git

### Setup Steps

1. Clone the repository

```bash
git clone https://github.com/xarafeddine/maktaba.git
cd maktaba
```

2. Install dependencies

```bash
go mod tidy
# or
make audit
```

## Local Setup

### Database Setup

Set up PostgreSQL database

1. Connect to PostgreSQL:

```bash
psql -h localhost -p 5432 -U postgres -W
```

2. Create and configure database:

```sql
CREATE DATABASE maktaba;
\c maktaba
CREATE ROLE maktaba WITH LOGIN PASSWORD 'maktaba';
CREATE EXTENSION IF NOT EXISTS citext;
ALTER DATABASE maktaba OWNER TO maktaba;
```

### Migrations

1. Install migration tool:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Run migrations:

```bash
migrate -path=./migrations -database="postgres://maktaba:maktaba@localhost:5432/maktaba?sslmode=disable" up
```

4. Run the application

```bash
go run ./cmd/api -db-dsn=${DB_DSN}
# or
make run/api
```

## 🐳 Docker Setup

### Prerequisites

- Docker 20.10+

### Running with Docker

1. Create a Docker network:

```bash
docker network create maktaba-net
```

2. Start the PostgreSQL container:

```bash
docker run --rm --name maktaba-db \
  --network maktaba-net \
  -e POSTGRES_DB=maktaba \
  -e POSTGRES_USER=maktaba \
  -e POSTGRES_PASSWORD=maktaba \
  postgres:15-alpine
```

3. Build the image

```bash
# Build the image
docker build -t maktaba-api .
```

4. run the API container:

```bash
docker run -d \
  --name maktaba-api \
  --network maktaba-net \
  -e DB_DSN=postgres://maktaba:secret@maktaba-db:5432/maktaba?sslmode=disable \
  -p 4000:4000 \
  maktaba-api
```

or

```bash
docker run -it --rm --name maktaba-api \
  --network maktaba-net \
  -p 4000:4000 \
  --env-file .env \
  maktaba-api
```

The migrations will run automatically when the container starts. You can check the migration status with:

```bash
docker logs maktaba-api
```

### Useful Docker Commands

```bash
# View container logs
docker logs maktaba-api
docker logs maktaba-db

# Check container status
docker ps

# Enter container shell
docker exec -it maktaba-api sh
docker exec -it maktaba-db psql -U maktaba

# Stop containers
docker stop maktaba-api maktaba-db

# Remove containers
docker rm maktaba-api maktaba-db
```

The API will be accessible at `http://localhost:4000`

## 🎡 Kubernetes Deployment

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

## 🧪 Testing

Run unit and integration tests:

```bash
go test ./...
```

## 📊 Monitoring

- Prometheus metrics endpoint
- Expvar debugging endpoint at `/debug/vars`

## 🔒 Security Features

- Token based authentication
- Role-based access control
- Rate limiting
- CORS protection
- Secure password hashing

## 📝 Logging

- Structured logging
- Log levels: INFO, WARN, ERROR
- Configurable log output

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
