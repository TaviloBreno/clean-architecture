# Clean Architecture Order System

![Go](https://img.shields.io/badge/Go-1.21-blue.svg)
![MySQL](https://img.shields.io/badge/MySQL-8.0-orange.svg)
![gRPC](https://img.shields.io/badge/gRPC-Latest-green.svg)
![GraphQL](https://img.shields.io/badge/GraphQL-Latest-pink.svg)
![Docker](https://img.shields.io/badge/Docker-Compose-blue.svg)

**A comprehensive order management system built with Clean Architecture principles, featuring multiple API interfaces: REST, gRPC, and GraphQL.**

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Technology Stack](#technology-stack)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Development](#development)
- [Testing](#testing)

---

## Overview

This project implements a robust order management system following Clean Architecture principles. It demonstrates how to build a single application that exposes the same business logic through multiple interfaces.

### Features

- **Order Creation**: Create new orders with price and tax calculations
- **Order Listing**: Retrieve all orders with pagination support
- **Multi-Interface API**: Access the same functionality via REST, gRPC, or GraphQL
- **Clean Architecture**: Separation of concerns with clear dependency boundaries
- **Database Integration**: MySQL with automated migrations
- **Containerized Deployment**: Docker Compose for easy setup

---

## Architecture

The project follows Robert C. Martin's Clean Architecture principles:

```
┌─────────────────────────────────────────────────────────┐
│                    Frameworks & Drivers                 │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   │
│  │   Web   │  │  gRPC   │  │GraphQL  │  │  MySQL  │   │
│  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │
└─────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────┐
│                Interface Adapters                       │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   │
│  │Controllers│ │Services │ │Resolvers│ │Repository│   │
│  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │
└─────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────┐
│                   Use Cases                             │
│  ┌─────────────────┐  ┌─────────────────┐             │
│  │  CreateOrder    │  │   ListOrders    │             │
│  └─────────────────┘  └─────────────────┘             │
└─────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────┐
│                    Entities                             │
│  ┌─────────────────┐  ┌─────────────────┐             │
│  │     Order       │  │   Repository    │             │
│  │   (Domain)      │  │  (Interface)    │             │
│  └─────────────────┘  └─────────────────┘             │
└─────────────────────────────────────────────────────────┘
```

### Directory Structure

```
├── cmd/server/               # Application entry point
├── internal/
│   ├── entity/              # Business entities and interfaces
│   ├── usecase/             # Application business rules
│   └── infra/               # External interfaces
│       ├── repository/      # Data access layer
│       ├── web/            # REST API handlers
│       ├── grpc/           # gRPC service implementation
│       └── graphql/        # GraphQL resolvers
├── proto/                   # Protocol Buffer definitions
├── migrations/              # Database migrations
└── docker-compose.yaml      # Container orchestration
```

---

## Technology Stack

| Component | Technology | Version |
|-----------|------------|---------|
| **Language** | Go | 1.21+ |
| **Database** | MySQL | 8.0 |
| **API Protocols** | REST, gRPC, GraphQL | Latest |
| **Serialization** | Protocol Buffers | v3 |
| **Containerization** | Docker Compose | Latest |
| **Dependency Management** | Go Modules | Latest |

---

## Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/) for cloning the repository

### Installation & Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/TaviloBreno/clean-architecture.git
   cd clean-architecture
   ```

2. **Start all services**
   ```bash
   docker compose up --build
   ```

3. **Verify services are running**
   ```bash
   # Check application health
   curl http://localhost:8080/order
   
   # View container logs
   docker compose logs -f
   ```

The application will automatically:
- Start MySQL database with proper schema
- Run database migrations
- Initialize all API services

### Service Endpoints

| Service | Port | Endpoint | Description |
|---------|------|----------|-------------|
| **REST API** | 8080 | `http://localhost:8080` | HTTP JSON API |
| **gRPC** | 50051 | `localhost:50051` | Binary protocol |
| **GraphQL** | 8081 | `http://localhost:8081` | Query language API |
| **MySQL** | 3306 | `localhost:3306` | Database |

## API Documentation

### REST API (Port 8080)

#### Create Order
```http
POST /order
Content-Type: application/json

{
  "id": "",
  "price": "100.50", 
  "tax": "10.50"
}
```

**Response:**
```json
{
  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "price": 100.50,
  "tax": 10.50,
  "final_price": 111.00
}
```

#### List Orders
```http
GET /order
```

**Response:**
```json
[
  {
    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "price": 100.50,
    "tax": 10.50,
    "final_price": 111.00
  }
]
```

### gRPC API (Port 50051)

#### Service Definition
```protobuf
service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
  rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);
}
```

#### Usage Examples
```bash
# Install grpcurl (if needed)
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List available services
grpcurl -plaintext localhost:50051 list

# Create order
grpcurl -plaintext -d '{"price": "100.50", "tax": "10.50"}' \
  localhost:50051 pb.OrderService/CreateOrder

# List orders
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

### GraphQL API (Port 8081)

#### Interactive Playground
Visit: `http://localhost:8081/`

#### Query Examples
**List Orders:**
```graphql
query {
  orders {
    id
    price
    tax
    final_price
  }
}
```

**HTTP Request:**
```bash
curl -X POST http://localhost:8081/query \
  -H "Content-Type: application/json" \
  -d '{"query": "{ orders { id price tax final_price } }"}'
```

---

## Development

### Local Development Setup

1. **Install Go 1.21+**
   ```bash
   # Check Go version
   go version
   ```

2. **Start MySQL locally**
   ```bash
   docker run --name mysql-orders -p 3306:3306 \
     -e MYSQL_ROOT_PASSWORD=root \
     -e MYSQL_DATABASE=orders \
     -e MYSQL_USER=user \
     -e MYSQL_PASSWORD=password \
     -d mysql:8.0
   ```

3. **Run migrations**
   ```sql
   USE orders;
   CREATE TABLE IF NOT EXISTS orders (
       id VARCHAR(36) PRIMARY KEY,
       price DECIMAL(10,2) NOT NULL,
       tax DECIMAL(10,2) NOT NULL,
       final_price DECIMAL(10,2) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

4. **Start the application**
   ```bash
   go run cmd/server/main.go
   ```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `3306` | Database port |
| `DB_USER` | `user` | Database username |
| `DB_PASSWORD` | `password` | Database password |
| `DB_NAME` | `orders` | Database name |

### Database Schema

```sql
CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    price DECIMAL(10,2) NOT NULL,
    tax DECIMAL(10,2) NOT NULL, 
    final_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Testing

### Using api.http File

The project includes an `api.http` file with ready-to-use examples for VS Code with the REST Client extension:

1. Open `api.http` in VS Code
2. Install the REST Client extension
3. Click "Send Request" above each example
4. Test order creation and listing

### Manual Testing

**Health Check:**
```bash
curl http://localhost:8080/order
```

**Create Order:**
```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"price": "100.50", "tax": "10.50"}'
```

### Service Monitoring

Monitor service logs:
```bash
# All services
docker compose logs -f

# Specific service
docker compose logs -f app
docker compose logs -f mysql
```

Expected startup logs:
```
app_1    | gRPC server running on port 50051
app_1    | GraphQL server running on port 8081 
app_1    | REST server running on port 8080
```

---

### Estrutura dos Use Cases

O projeto implementa os seguintes casos de uso:

1. **CreateOrderUseCase**: Responsável pela criação de novos pedidos
2. **ListOrdersUseCase**: Responsável pela listagem de pedidos

### Implementação da Clean Architecture

- **Entities**: Regras de negócio centrais (Order)
- **Use Cases**: Regras de negócio da aplicação
- **Interface Adapters**: Controllers REST, gRPC Services, GraphQL Resolvers
- **Frameworks & Drivers**: Database, Web Framework

## 📋 Troubleshooting

### Problemas Comuns

1. **Erro de conexão com MySQL**: Certifique-se de que o MySQL está rodando e as credenciais estão corretas
2. **Porta já em uso**: Verifique se as portas 3306, 8080, 50051, e 8081 estão livres
3. **Erro de build**: Execute `go mod tidy` para baixar as dependências

### Verificar se os serviços estão rodando

```bash
# Verificar se o MySQL está rodando
docker ps | grep mysql

# Verificar se a aplicação está respondendo
curl http://localhost:8080/order
```

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto é parte do curso Go Expert da Full Cycle e é destinado para fins educacionais.#   c l e a n - a r c h i t e c t u r e 
 
 