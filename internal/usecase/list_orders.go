package usecase

import "github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/entity"

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []ListOrdersOutputDTO
	for _, order := range orders {
		dto := ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersOutput = append(ordersOutput, dto)
	}

	return ordersOutput, nil
}
