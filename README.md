# Clean Architecture Order System# Clean Architecture Order System



![Go](https://img.shields.io/badge/Go-1.21-blue.svg)![Go](https://img.shields.io/badge/Go-1.21-blue.svg)

![MySQL](https://img.shields.io/badge/MySQL-8.0-orange.svg)![MySQL](https://img.shields.io/badge/MySQL-8.0-orange.svg)

![gRPC](https://img.shields.io/badge/gRPC-Latest-green.svg)![gRPC](https://img.shields.io/badge/gRPC-Latest-green.svg)

![GraphQL](https://img.shields.io/badge/GraphQL-Latest-pink.svg)![GraphQL](https://img.shields.io/badge/GraphQL-Latest-pink.svg)

![Docker](https://img.shields.io/badge/Docker-Compose-blue.svg)![Docker](https://img.shields.io/badge/Docker-Compose-blue.svg)



**A comprehensive order management system built with Clean Architecture principles, featuring multiple API interfaces: REST, gRPC, and GraphQL.****A comprehensive order management system built with Clean Architecture principles, featuring multiple API interfaces: REST, gRPC, and GraphQL.**



------



## Table of Contents## Table of Contents



- [Overview](#overview)- [Overview](#overview)

- [Architecture](#architecture)- [Architecture](#architecture)

- [Technology Stack](#technology-stack)- [Technology Stack](#technology-stack)

- [Quick Start](#quick-start)- [Quick Start](#quick-start)

- [API Documentation](#api-documentation)- [API Documentation](#api-documentation)

- [Development](#development)- [Development](#development)

- [Testing](#testing)- [Testing](#testing)

- [Troubleshooting](#troubleshooting)

- [Contributing](#contributing)---



---## Overview



## OverviewThis project implements a robust order management system following Clean Architecture principles. It demonstrates how to build a single application that exposes the same business logic through multiple interfaces.



This project implements a robust order management system following Clean Architecture principles. It demonstrates how to build a single application that exposes the same business logic through multiple interfaces.### Features



### Features- **Order Creation**: Create new orders with price and tax calculations

- **Order Listing**: Retrieve all orders with pagination support

- **Order Creation**: Create new orders with price and tax calculations- **Multi-Interface API**: Access the same functionality via REST, gRPC, or GraphQL

- **Order Listing**: Retrieve all orders with pagination support- **Clean Architecture**: Separation of concerns with clear dependency boundaries

- **Multi-Interface API**: Access the same functionality via REST, gRPC, or GraphQL- **Database Integration**: MySQL with automated migrations

- **Clean Architecture**: Separation of concerns with clear dependency boundaries- **Containerized Deployment**: Docker Compose for easy setup

- **Database Integration**: MySQL with automated migrations

- **Containerized Deployment**: Docker Compose for easy setup---



---## Architecture



## ArchitectureThe project follows Robert C. Martin's Clean Architecture principles:



The project follows Robert C. Martin's Clean Architecture principles:```

┌─────────────────────────────────────────────────────────┐

```│                    Frameworks & Drivers                 │

┌─────────────────────────────────────────────────────────┐│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   │

│                    Frameworks & Drivers                 ││  │   Web   │  │  gRPC   │  │GraphQL  │  │  MySQL  │   │

│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   ││  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │

│  │   Web   │  │  gRPC   │  │GraphQL  │  │  MySQL  │   │└─────────────────────────────────────────────────────────┘

│  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │                              │

└─────────────────────────────────────────────────────────┘┌─────────────────────────────────────────────────────────┐

                              ││                Interface Adapters                       │

┌─────────────────────────────────────────────────────────┐│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   │

│                Interface Adapters                       ││  │Controllers│ │Services │ │Resolvers│ │Repository│   │

│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐   ││  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │

│  │Controllers│ │Services │ │Resolvers│ │Repository│   │└─────────────────────────────────────────────────────────┘

│  └─────────┘  └─────────┘  └─────────┘  └─────────┘   │                              │

└─────────────────────────────────────────────────────────┘┌─────────────────────────────────────────────────────────┐

                              ││                   Use Cases                             │

┌─────────────────────────────────────────────────────────┐│  ┌─────────────────┐  ┌─────────────────┐             │

│                   Use Cases                             ││  │  CreateOrder    │  │   ListOrders    │             │

│  ┌─────────────────┐  ┌─────────────────┐             ││  └─────────────────┘  └─────────────────┘             │

│  │  CreateOrder    │  │   ListOrders    │             │└─────────────────────────────────────────────────────────┘

│  └─────────────────┘  └─────────────────┘             │                              │

└─────────────────────────────────────────────────────────┘┌─────────────────────────────────────────────────────────┐

                              ││                    Entities                             │

┌─────────────────────────────────────────────────────────┐│  ┌─────────────────┐  ┌─────────────────┐             │

│                    Entities                             ││  │     Order       │  │   Repository    │             │

│  ┌─────────────────┐  ┌─────────────────┐             ││  │   (Domain)      │  │  (Interface)    │             │

│  │     Order       │  │   Repository    │             ││  └─────────────────┘  └─────────────────┘             │

│  │   (Domain)      │  │  (Interface)    │             │└─────────────────────────────────────────────────────────┘

│  └─────────────────┘  └─────────────────┘             │```

└─────────────────────────────────────────────────────────┘

```### Directory Structure



### Directory Structure```

├── cmd/server/               # Application entry point

```├── internal/

├── cmd/server/               # Application entry point│   ├── entity/              # Business entities and interfaces

├── internal/│   ├── usecase/             # Application business rules

│   ├── entity/              # Business entities and interfaces│   └── infra/               # External interfaces

│   ├── usecase/             # Application business rules│       ├── repository/      # Data access layer

│   └── infra/               # External interfaces│       ├── web/            # REST API handlers

│       ├── repository/      # Data access layer│       ├── grpc/           # gRPC service implementation

│       ├── web/            # REST API handlers│       └── graphql/        # GraphQL resolvers

│       ├── grpc/           # gRPC service implementation├── proto/                   # Protocol Buffer definitions

│       └── graphql/        # GraphQL resolvers├── migrations/              # Database migrations

├── proto/                   # Protocol Buffer definitions└── docker-compose.yaml      # Container orchestration

├── migrations/              # Database migrations```

└── docker-compose.yaml      # Container orchestration

```---



---## Technology Stack



## Technology Stack| Component | Technology | Version |

|-----------|------------|---------|

| Component | Technology | Version || **Language** | Go | 1.21+ |

|-----------|------------|---------|| **Database** | MySQL | 8.0 |

| **Language** | Go | 1.21+ || **API Protocols** | REST, gRPC, GraphQL | Latest |

| **Database** | MySQL | 8.0 || **Serialization** | Protocol Buffers | v3 |

| **API Protocols** | REST, gRPC, GraphQL | Latest || **Containerization** | Docker Compose | Latest |

| **Serialization** | Protocol Buffers | v3 || **Dependency Management** | Go Modules | Latest |

| **Containerization** | Docker Compose | Latest |

| **Dependency Management** | Go Modules | Latest |---



---## Quick Start



## Quick Start### Prerequisites



### Prerequisites- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

- [Git](https://git-scm.com/) for cloning the repository

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

- [Git](https://git-scm.com/) for cloning the repository### Installation & Setup



### Installation & Setup1. **Clone the repository**

   ```bash

1. **Clone the repository**   git clone https://github.com/TaviloBreno/clean-architecture.git

   ```bash   cd clean-architecture

   git clone https://github.com/TaviloBreno/clean-architecture.git   ```

   cd clean-architecture

   ```2. **Start all services**

   ```bash

2. **Start all services**   docker compose up --build

   ```bash   ```

   docker compose up --build

   ```3. **Verify services are running**

   ```bash

3. **Verify services are running**   # Check application health

   ```bash   curl http://localhost:8080/order

   # Check application health   

   curl http://localhost:8080/order   # View container logs

      docker compose logs -f

   # View container logs   ```

   docker compose logs -f

   ```The application will automatically:

- Start MySQL database with proper schema

The application will automatically:- Run database migrations

- Start MySQL database with proper schema- Initialize all API services

- Run database migrations

- Initialize all API services### Service Endpoints



### Service Endpoints| Service | Port | Endpoint | Description |

|---------|------|----------|-------------|

| Service | Port | Endpoint | Description || **REST API** | 8080 | `http://localhost:8080` | HTTP JSON API |

|---------|------|----------|-------------|| **gRPC** | 50051 | `localhost:50051` | Binary protocol |

| **REST API** | 8080 | `http://localhost:8080` | HTTP JSON API || **GraphQL** | 8081 | `http://localhost:8081` | Query language API |

| **gRPC** | 50051 | `localhost:50051` | Binary protocol || **MySQL** | 3306 | `localhost:3306` | Database |

| **GraphQL** | 8081 | `http://localhost:8081` | Query language API |

| **MySQL** | 3306 | `localhost:3306` | Database |## API Documentation



---### REST API (Port 8080)



## API Documentation#### Create Order

```http

### REST API (Port 8080)POST /order

Content-Type: application/json

#### Create Order

```http{

POST /order  "id": "",

Content-Type: application/json  "price": "100.50", 

  "tax": "10.50"

{}

  "id": "",```

  "price": "100.50", 

  "tax": "10.50"**Response:**

}```json

```{

  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",

**Response:**  "price": 100.50,

```json  "tax": 10.50,

{  "final_price": 111.00

  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",}

  "price": 100.50,```

  "tax": 10.50,

  "final_price": 111.00#### List Orders

}```http

```GET /order

```

#### List Orders

```http**Response:**

GET /order```json

```[

  {

**Response:**    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",

```json    "price": 100.50,

[    "tax": 10.50,

  {    "final_price": 111.00

    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",  }

    "price": 100.50,]

    "tax": 10.50,```

    "final_price": 111.00

  }### gRPC API (Port 50051)

]

```#### Service Definition

```protobuf

### gRPC API (Port 50051)service OrderService {

  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);

#### Service Definition  rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);

```protobuf}

service OrderService {```

  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);

  rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);#### Usage Examples

}```bash

```# Install grpcurl (if needed)

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

#### Usage Examples

```bash# List available services

# Install grpcurl (if needed)grpcurl -plaintext localhost:50051 list

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Create order

# List available servicesgrpcurl -plaintext -d '{"price": "100.50", "tax": "10.50"}' \

grpcurl -plaintext localhost:50051 list  localhost:50051 pb.OrderService/CreateOrder



# Create order# List orders

grpcurl -plaintext -d '{"price": "100.50", "tax": "10.50"}' \grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

  localhost:50051 pb.OrderService/CreateOrder```



# List orders### GraphQL API (Port 8081)

grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

```#### Interactive Playground

Visit: `http://localhost:8081/`

### GraphQL API (Port 8081)

#### Query Examples

#### Interactive Playground**List Orders:**

Visit: `http://localhost:8081/````graphql

query {

#### Query Examples  orders {

**List Orders:**    id

```graphql    price

query {    tax

  orders {    final_price

    id  }

    price}

    tax```

    final_price

  }**HTTP Request:**

}```bash

```curl -X POST http://localhost:8081/query \

  -H "Content-Type: application/json" \

**HTTP Request:**  -d '{"query": "{ orders { id price tax final_price } }"}'

```bash```

curl -X POST http://localhost:8081/query \

  -H "Content-Type: application/json" \---

  -d '{"query": "{ orders { id price tax final_price } }"}'

```## Development



---### Local Development Setup



## Development1. **Install Go 1.21+**

   ```bash

### Local Development Setup   # Check Go version

   go version

1. **Install Go 1.21+**   ```

   ```bash

   # Check Go version2. **Start MySQL locally**

   go version   ```bash

   ```   docker run --name mysql-orders -p 3306:3306 \

     -e MYSQL_ROOT_PASSWORD=root \

2. **Start MySQL locally**     -e MYSQL_DATABASE=orders \

   ```bash     -e MYSQL_USER=user \

   docker run --name mysql-orders -p 3306:3306 \     -e MYSQL_PASSWORD=password \

     -e MYSQL_ROOT_PASSWORD=root \     -d mysql:8.0

     -e MYSQL_DATABASE=orders \   ```

     -e MYSQL_USER=user \

     -e MYSQL_PASSWORD=password \3. **Run migrations**

     -d mysql:8.0   ```sql

   ```   USE orders;

   CREATE TABLE IF NOT EXISTS orders (

3. **Run migrations**       id VARCHAR(36) PRIMARY KEY,

   ```sql       price DECIMAL(10,2) NOT NULL,

   USE orders;       tax DECIMAL(10,2) NOT NULL,

   CREATE TABLE IF NOT EXISTS orders (       final_price DECIMAL(10,2) NOT NULL,

       id VARCHAR(36) PRIMARY KEY,       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

       price DECIMAL(10,2) NOT NULL,   );

       tax DECIMAL(10,2) NOT NULL,   ```

       final_price DECIMAL(10,2) NOT NULL,

       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP4. **Start the application**

   );   ```bash

   ```   go run cmd/server/main.go

   ```

4. **Start the application**

   ```bash### Environment Variables

   go run cmd/server/main.go

   ```| Variable | Default | Description |

|----------|---------|-------------|

### Environment Variables| `DB_HOST` | `localhost` | Database host |

| `DB_PORT` | `3306` | Database port |

| Variable | Default | Description || `DB_USER` | `user` | Database username |

|----------|---------|-------------|| `DB_PASSWORD` | `password` | Database password |

| `DB_HOST` | `localhost` | Database host || `DB_NAME` | `orders` | Database name |

| `DB_PORT` | `3306` | Database port |

| `DB_USER` | `user` | Database username |### Database Schema

| `DB_PASSWORD` | `password` | Database password |

| `DB_NAME` | `orders` | Database name |```sql

CREATE TABLE orders (

### Database Schema    id VARCHAR(36) PRIMARY KEY,

    price DECIMAL(10,2) NOT NULL,

```sql    tax DECIMAL(10,2) NOT NULL, 

CREATE TABLE orders (    final_price DECIMAL(10,2) NOT NULL,

    id VARCHAR(36) PRIMARY KEY,    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

    price DECIMAL(10,2) NOT NULL,);

    tax DECIMAL(10,2) NOT NULL, ```

    final_price DECIMAL(10,2) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP---

);

```## Testing



---### Using api.http File



## TestingThe project includes an `api.http` file with ready-to-use examples for VS Code with the REST Client extension:



### Using api.http File1. Open `api.http` in VS Code

2. Install the REST Client extension

The project includes an `api.http` file with ready-to-use examples for VS Code with the REST Client extension:3. Click "Send Request" above each example

4. Test order creation and listing

1. Open `api.http` in VS Code

2. Install the REST Client extension### Manual Testing

3. Click "Send Request" above each example

4. Test order creation and listing**Health Check:**

```bash

### Manual Testingcurl http://localhost:8080/order

```

**Health Check:**

```bash**Create Order:**

curl http://localhost:8080/order```bash

```curl -X POST http://localhost:8080/order \

  -H "Content-Type: application/json" \

**Create Order:**  -d '{"price": "100.50", "tax": "10.50"}'

```bash```

curl -X POST http://localhost:8080/order \

  -H "Content-Type: application/json" \### Service Monitoring

  -d '{"price": "100.50", "tax": "10.50"}'

```Monitor service logs:

```bash

### Service Monitoring# All services

docker compose logs -f

Monitor service logs:

```bash# Specific service

# All servicesdocker compose logs -f app

docker compose logs -fdocker compose logs -f mysql

```

# Specific service

docker compose logs -f appExpected startup logs:

docker compose logs -f mysql```

```app_1    | gRPC server running on port 50051

app_1    | GraphQL server running on port 8081 

Expected startup logs:app_1    | REST server running on port 8080

``````

app_1    | gRPC server running on port 50051

app_1    | GraphQL server running on port 8081 ---

app_1    | REST server running on port 8080

```### Estrutura dos Use Cases



---O projeto implementa os seguintes casos de uso:



## Architecture Details1. **CreateOrderUseCase**: Responsável pela criação de novos pedidos

2. **ListOrdersUseCase**: Responsável pela listagem de pedidos

### Use Case Implementation

### Implementação da Clean Architecture

The application implements the following use cases following Clean Architecture principles:

- **Entities**: Regras de negócio centrais (Order)

- **CreateOrderUseCase**: Handles new order creation with business validation- **Use Cases**: Regras de negócio da aplicação

- **ListOrdersUseCase**: Retrieves orders with proper data transformation- **Interface Adapters**: Controllers REST, gRPC Services, GraphQL Resolvers

- **Frameworks & Drivers**: Database, Web Framework

### Layer Responsibilities

## 📋 Troubleshooting

| Layer | Components | Responsibilities |

|-------|------------|------------------|### Problemas Comuns

| **Entities** | `Order`, `OrderRepository` | Core business logic and rules |

| **Use Cases** | `CreateOrderUseCase`, `ListOrdersUseCase` | Application-specific business rules |1. **Erro de conexão com MySQL**: Certifique-se de que o MySQL está rodando e as credenciais estão corretas

| **Interface Adapters** | REST handlers, gRPC services, GraphQL resolvers | Convert data between use cases and external world |2. **Porta já em uso**: Verifique se as portas 3306, 8080, 50051, e 8081 estão livres

| **Frameworks & Drivers** | MySQL, HTTP server, gRPC server | External tools and frameworks |3. **Erro de build**: Execute `go mod tidy` para baixar as dependências



---### Verificar se os serviços estão rodando



## Troubleshooting```bash

# Verificar se o MySQL está rodando

### Common Issuesdocker ps | grep mysql



| Issue | Solution |# Verificar se a aplicação está respondendo

|-------|----------|curl http://localhost:8080/order

| **Port already in use** | Check if ports 3306, 8080, 50051, 8081 are available |```

| **MySQL connection failed** | Ensure MySQL container is running and credentials are correct |

| **Build errors** | Run `go mod tidy` to download dependencies |## 🤝 Contribuição

| **Container startup issues** | Check Docker logs with `docker compose logs` |

1. Faça um fork do projeto

### Health Checks2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)

3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)

```bash4. Push para a branch (`git push origin feature/nova-feature`)

# Check if all services are running5. Abra um Pull Request

docker compose ps

## 📄 Licença

# Test database connection

docker compose exec mysql mysql -u user -ppassword orders -e "SHOW TABLES;"Este projeto é parte do curso Go Expert da Full Cycle e é destinado para fins educacionais.#   c l e a n - a r c h i t e c t u r e 

 

# Test API endpoints 
curl -f http://localhost:8080/order || echo "REST API not responding"
grpcurl -plaintext localhost:50051 list || echo "gRPC not responding"  
curl -f http://localhost:8081/ || echo "GraphQL not responding"
```

### Performance Monitoring

```bash
# Monitor container resources
docker stats

# Check application logs
docker compose logs -f app

# Monitor database queries (if needed)
docker compose exec mysql mysql -u root -proot -e "SHOW PROCESSLIST;"
```

---

## Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/amazing-feature`
3. **Commit** your changes: `git commit -m 'Add amazing feature'`
4. **Push** to the branch: `git push origin feature/amazing-feature`
5. **Open** a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Maintain Clean Architecture principles
- Add tests for new functionality
- Update documentation as needed
- Use conventional commits

---

## License

This project is part of the **Go Expert Course** by [Full Cycle](https://fullcycle.com.br/) and is intended for educational purposes.

---

## Support

If you encounter any issues or have questions:

- Check the [Troubleshooting](#troubleshooting) section
- Review the [API Documentation](#api-documentation)
- Open an issue in the repository

---

**Built with ❤️ for learning Clean Architecture principles in Go**