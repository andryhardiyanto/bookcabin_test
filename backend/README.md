# BookCabin Backend

A Go-based REST API service for managing airline crew vouchers and seat assignments.

## Prerequisites

- Go 1.24.0 or higher
- SQLite3
- Docker (optional)

## Environment Setup

1. **Copy environment file**
   ```bash
   cp .env.example .env
   ```

2. **Configure environment variables in `.env`**
   ```env
   ENV=development
   APP_NAME=bookcabin
   APP_PORT=8080
   DB_PATH=./data/vouchers.db
   DB_MAX_OPEN_CONNS=10
   DB_MAX_IDLE_CONNS=5
   DB_CONN_MAX_IDLE_TIME=5m
   DB_CONN_MAX_LIFETIME=1h
   ```

## How to Run

### Option 1: Run Locally (without Docker)

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Create data directory**
   ```bash
   mkdir -p data
   ```
   
   **Note:** The `vouchers.db` file will be automatically created inside the `data` folder when you first run the application.

3. **Run the application**
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`

### Option 2: Run with Docker

1. **Build and run with Docker Compose**
   ```bash
   docker-compose up -d  # old version
   docker compose up -d  # new version
   ```

The server will start on `http://localhost:8080`

## API Endpoints

- `GET /health` - Health check
- `POST /api/check` - Check voucher availability
- `POST /api/generate` - Generate new voucher
