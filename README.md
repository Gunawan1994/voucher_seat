# Voucher seat

---

## Project Structure

```
.
├── krakend/
│   └── krakend.json          # API Gateway configuration
├── backend/                  
│   ├── Dockerfile
│   └── main.go
├── frontend/                     
│   ├── Dockerfile
│   └── src
└── docker-compose.yml
```

---

## Prerequisites

* Docker & Docker Compose
* Go 1.25+ (for building services)
* React (npm 10.8.2)
* GORM

---

## Environment Variables

### Backend Service (PostgreSQL)

* `LISTEN_PORT` : port the service listens on (e.g., `:8888`)
* `POSTGRES_ADDR` : PostgreSQL host (service name in Docker, e.g., `postgres`)
* `POSTGRES_PORT` : PostgreSQL port (e.g., `5432`)
* `POSTGRES_DATABASE` : Database name (e.g., `vouchers`)
* `POSTGRES_PASSWORD` : PostgreSQL password

---

## Docker Compose

Start all services including Krakend:

```bash
docker compose up -d
```

* **Forontend Service:** [http://localhost:3000](http://localhost:3000)
* **Bakend Service:** [http://localhost:8888](http://localhost:8888)
* **Krakend API Gateway:** [http://localhost:8080](http://localhost:8080)

---

## API Endpoints

### Backend Service

* `GET /api/check` → check flight number
* `GET /api/generate` → generate seats number

---

## Commands

* Build all services:

```bash
docker compose up --build
```

* Stop and remove containers:

```bash
docker compose down -v
```

* View logs:

```bash
docker compose logs -f
```
