FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM alpine:3.19
RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /app
COPY --from=builder /app/api .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations ./migrations



EXPOSE 4000
CMD echo "Running migrations..." && \
migrate -path=/app/migrations -database="${DB_DSN}" up && \
echo "Starting application..." && \
./api -db-dsn="${DB_DSN}" -port="${PORT}"
