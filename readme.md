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

| Variable      | Description                  | Default       |
| ------------- | ---------------------------- | ------------- |
| `DB_DSN`      | PostgreSQL connection string | Required      |
| `SERVER_PORT` | API server port              | `4000`        |
| `ENV_MODE`    | Environment mode             | `development` |

### Command Line Flags

```bash
go run main.go \
  -db-dsn="postgres://username:password@localhost/dbname" \
  -port=4000 \
  -env=production
```

## 📦 Installation

### Prerequisites

- Go 1.22+
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

3. Set up PostgreSQL database

```bash
# Create database
CREATE DATABASE maktaba;

# Create necessary tables using migrations
make db/migrate/up
```

4. Run the application

```bash
go run ./cmd/api -db-dsn=${MAKTABA_DB_DSN}
# or
make run/api
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
