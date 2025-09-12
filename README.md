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