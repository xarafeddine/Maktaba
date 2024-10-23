# MyMovieApi

MyMovieApi is a JSON API written in Go for retrieving and managing information about movies. The core functionality is inspired by the Open Movie Database API, providing endpoints for CRUD operations on movie data, user management, and authentication. It is designed to serve as the backend for applications that need movie data and user authentication.

## Features

- Retrieve movie information
- Create, update, and delete movie entries
- User registration and management
- User authentication with token-based login
- Health check and monitoring endpoints

## Endpoints

The following endpoints and actions are supported:

### Health Check

| Method | URL Pattern       | Action                                   |
| ------ | ----------------- | ---------------------------------------- |
| GET    | `/v1/healthcheck` | Show application health and version info |

### Movie Endpoints

| Method | URL Pattern      | Action                                 |
| ------ | ---------------- | -------------------------------------- |
| GET    | `/v1/movies`     | Show details of all movies             |
| POST   | `/v1/movies`     | Create a new movie                     |
| GET    | `/v1/movies/:id` | Show details of a specific movie       |
| PATCH  | `/v1/movies/:id` | Update the details of a specific movie |
| DELETE | `/v1/movies/:id` | Delete a specific movie                |

### User Endpoints

| Method | URL Pattern           | Action                   |
| ------ | --------------------- | ------------------------ |
| POST   | `/v1/users`           | Register a new user      |
| PUT    | `/v1/users/activated` | Activate a specific user |
| PUT    | `/v1/users/password`  | Update a user's password |

### Authentication Endpoints

| Method | URL Pattern                 | Action                              |
| ------ | --------------------------- | ----------------------------------- |
| POST   | `/v1/tokens/authentication` | Generate a new authentication token |
| POST   | `/v1/tokens/password-reset` | Generate a new password-reset token |

### Debugging & Metrics

| Method | URL Pattern   | Action                      |
| ------ | ------------- | --------------------------- |
| GET    | `/debug/vars` | Display application metrics |

## Setup and Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/mymovieapi.git
   ```

2. Navigate into the project directory:

   ```bash
   cd mymovieapi
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## Configuration

The application reads configuration values from command-line flags. Here are the available flags:

| Flag                 | Description                                    |
| -------------------- | ---------------------------------------------- |
| `-db-dsn`            | PostgreSQL DSN                                 |
| `-db-max-idle-conns` | PostgreSQL max idle connections                |
| `-db-max-idle-time`  | PostgreSQL max connection idle time            |
| `-db-max-open-conns` | PostgreSQL max open connections                |
| `-env`               | Environment (development, staging, production) |
| `-port`              | API server port                                |

### Example

To run the application with custom configuration, use the following command:

```bash
go run main.go -db-dsn="your-dsn" -port=8080 -env=production
```
