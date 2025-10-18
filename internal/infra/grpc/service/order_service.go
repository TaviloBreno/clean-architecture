package service

import (
	"context"

	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewOrderService(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrdersUseCase *usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.CreateOrderInputDTO{
		ID:    req.GetId(),
		Price: req.GetPrice(),
		Tax:   req.GetTax(),
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.OrderResponse
	for _, order := range orders {
		orderResponse := &pb.OrderResponse{
			Id:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.ListOrdersResponse{
		Orders: ordersResponse,
	}, nil
}
