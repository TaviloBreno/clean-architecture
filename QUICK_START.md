# 🚀 Quick Start Guide

## Iniciar o Projeto

```bash
docker compose up --build
```

## Testar os Endpoints

### 1. REST API (Porta 8080)

**Criar Pedido:**
```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"price": "100.50", "tax": "10.50"}'
```

**Listar Pedidos:**
```bash
curl http://localhost:8080/order
```

### 2. GraphQL API (Porta 8081)

**Acessar Playground:**
- Abra: http://localhost:8081/

**Listar Pedidos:**
```bash
curl -X POST http://localhost:8081/query \
  -H "Content-Type: application/json" \
  -d '{"query": "{ orders { id price tax final_price } }"}'
```

### 3. gRPC (Porta 50051)

**Com grpcurl:**
```bash
# Instalar grpcurl (se necessário)
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Listar serviços
grpcurl -plaintext localhost:50051 list

# Criar pedido
grpcurl -plaintext -d '{"price": "100.50", "tax": "10.50"}' localhost:50051 pb.OrderService/CreateOrder

# Listar pedidos  
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

## Verificar Status

```bash
# Ver logs dos containers
docker compose logs -f

# Verificar se os serviços estão rodando
curl http://localhost:8080/order
```

## Parar o Projeto

```bash
docker compose down
```