# fizzbuzz-api

A minimal Go service showcasing modular design, observability, and robust request handling.  
The system is composed of **two APIs** and a **gateway**, designed to demonstrate clean architecture, inter-service gRPC, and basic analytics.

---

## 📦 Overview

This project implements the classic **FizzBuzz** as a service and adds request statistics.

- **api-fizzbuzz**
    - REST endpoint (`/fizzbuzz`) that computes FizzBuzz.
    - Validates input; returns JSON.
    - For each HTTP request, sends a **gRPC** message to `api-statistics` to record the event.

- **api-statistics**
    - Stores historical data for analytics.
    - REST endpoints under `/statistics` to query most requested result, top numbers, and time-series.

- **gateway**
  - **Reverse proxy** that accepts client HTTP calls and routes them to `api-fizzbuzz` or `api-statistics`.

- **docs** → Node.js + Swagger UI service for API documentation.

## 🏗️ Architecture
```
                                     ┌───────────┐
                                     │   Client  │
                                     └─────┬─────┘
                                           │ HTTP
                                           ▼
                                    ┌───────────────┐
                                    │   Gateway     │
                                    │  (Reverse     │
                                    │   Proxy)      │
                                    └─────┬─────────┘
                                          │
                               ┌──────────┴────────────┐
                               │                       │
                               ▼                       ▼
                        ┌─────────────┐       ┌────────────────┐
                        │ api-fizzbuzz│  gRPC │ api-statistics │
                        │   (REST)    │──────▶│  (gRPC + REST) │
                        └─────────────┘       └────────────────┘
```

---

## 🛠 Tech Stack

- **Go 1.25.0**
- **gRPC** (api-fizzbuzz → api-statistics)
- **net/http** (REST)
- **Postgres** (configurable) for statistics storage
- **Docker Compose** for local setup

---

## 🚀 Quickstart

### Run everything with Docker Compose

```bash 
docker-compose up --build
```

## 🌐 Services & Routing

All services run inside Docker and expose their own ports, **but you don’t need to call them directly**.  
The **Gateway** at [http://localhost:8080](http://localhost:8080) acts as the single entrypoint, forwarding requests to the right service.

### Direct service endpoints
- **Gateway** → [http://localhost:8080](http://localhost:8080)
- **FizzBuzz API** → [http://localhost:8081/fizzbuzz](http://localhost:8081/fizzbuzz)
- **Statistics API** → [http://localhost:8082/stats](http://localhost:8082/stats)
- **Swagger Docs** → [http://localhost:8085/docs](http://localhost:8085/docs)  
  _(also proxied by the gateway at [http://localhost:8080/docs/#](http://localhost:8080/docs/#))_

### Access through the Gateway (recommended)

- **FizzBuzz API**
    - `POST /fizzbuzz/get` → compute a FizzBuzz sequence  
      **Example:**
      ```bash
      curl -X POST http://localhost:8080/fizzbuzz/get \
        -H "Content-Type: application/json" \
        -d '{"limit":15,"multiple1":3,"multiple2":5,"replacement_string1":"Fizz","replacement_string2":"Buzz"}'
      ```

- **Statistics API**
    - `GET /statistics/get` → most frequent request
      ```bash
      curl http://localhost:8080/statistics/get
      ```

- **Swagger Docs**
    - `GET /docs` → Swagger UI served by the docs service  
      Accessible at:
        - [http://localhost:8085/docs](http://localhost:8085/docs) (direct)
        - [http://localhost:8080/docs](http://localhost:8080/docs) (via gateway)

## 🔮 Future Considerations
For production, the architecture could evolve to support:
- CQRS with separate read/write models
- Event-driven stats collection using a message queue (instead of direct gRPC)
- Enhanced observability and monitoring

## 📜 License
Licensed under the [MIT License](LICENSE).