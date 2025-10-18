package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/grpc/service"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/graphql/graph"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/graphql/graph/model"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/repository"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/web"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/usecase"
)

func main() {
	// Database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnv("DB_USER", "user"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "orders"),
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repository
	orderRepository := repository.NewOrderRepository(db)

	// Use cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)

	// Start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		orderService := service.NewOrderService(createOrderUseCase, listOrdersUseCase)
		pb.RegisterOrderServiceServer(grpcServer, orderService)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			panic(err)
		}
		log.Println("gRPC server running on port 50051")
		grpcServer.Serve(lis)
	}()

	// Start GraphQL server
	go func() {
		resolver := graph.NewResolver(createOrderUseCase, listOrdersUseCase)
		
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`
				<!DOCTYPE html>
				<html>
				<head>
					<title>GraphQL Playground</title>
				</head>
				<body>
					<h1>GraphQL API</h1>
					<p>Use POST /query to execute GraphQL queries</p>
					<h2>Example Query:</h2>
					<pre>
					{
						"query": "{ orders { id price tax final_price } }"
					}
					</pre>
					<h2>Example Mutation:</h2>
					<pre>
					{
						"query": "mutation { createOrder(input: {price: \"100.0\", tax: \"10.0\"}) { id price tax final_price } }"
					}
					</pre>
				</body>
				</html>
			`))
		})

		http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			
			var req struct {
				Query string `json:"query"`
			}
			
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			// Simple GraphQL query handling
			if req.Query == "{ orders { id price tax final_price } }" {
				orders, err := resolver.ListOrdersUseCase.Execute()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				
				var result []*model.Order
				for _, order := range orders {
					result = append(result, &model.Order{
						ID:         order.ID,
						Price:      order.Price,
						Tax:        order.Tax,
						FinalPrice: order.FinalPrice,
					})
				}
				
				response := map[string]interface{}{
					"data": map[string]interface{}{
						"orders": result,
					},
				}
				
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
				return
			}
			
			// Handle other queries...
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"data": nil,
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Query not supported yet",
					},
				},
			})
		})

		log.Println("GraphQL server running on port 8081")
		log.Println("connect to http://localhost:8081/ for GraphQL playground")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	// REST API server
	webOrderHandler := web.NewWebOrderHandler(createOrderUseCase, listOrdersUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			webOrderHandler.Create(w, r)
		case "GET":
			webOrderHandler.List(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("REST server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}