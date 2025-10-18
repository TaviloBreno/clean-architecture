package graph

import (
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/usecase"
)

type Resolver struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewResolver(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrdersUseCase *usecase.ListOrdersUseCase,
) *Resolver {
	return &Resolver{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}
