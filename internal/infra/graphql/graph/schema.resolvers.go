package graph

import (
	"context"

	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/graphql/graph/model"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/usecase"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error) {
	dto := usecase.CreateOrderInputDTO{
		Price: input.Price,
		Tax:   input.Tax,
	}

	if input.ID != nil {
		dto.ID = *input.ID
	}

	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:         output.ID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
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

	return result, nil
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type MutationResolver interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error)
}
type QueryResolver interface {
	Orders(ctx context.Context) ([]*model.Order, error)
}
