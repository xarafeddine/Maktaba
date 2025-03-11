#!/bin/sh
set -e

# Install migrate tool
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Wait for postgres to be ready
echo "Waiting for postgres..."
while ! nc -z db 5432; do
  sleep 1
done

# Run migrations
echo "Running migrations..."
migrate -path=/app/migrations -database="${DB_DSN}" up

# Start the application
echo "Starting application..."
exec "$@"
