package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Request and Response types
type CreateOrderRequest struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax   string `protobuf:"bytes,3,opt,name=tax,proto3" json:"tax,omitempty"`
}

func (x *CreateOrderRequest) Reset()                             { *x = CreateOrderRequest{} }
func (x *CreateOrderRequest) String() string                     { return "CreateOrderRequest" }
func (*CreateOrderRequest) ProtoMessage()                        {}
func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message { return nil }

func (x *CreateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateOrderRequest) GetTax() string {
	if x != nil {
		return x.Tax
	}
	return ""
}

type CreateOrderResponse struct {
	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price      float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float64 `protobuf:"fixed64,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float64 `protobuf:"fixed64,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *CreateOrderResponse) Reset()                             { *x = CreateOrderResponse{} }
func (x *CreateOrderResponse) String() string                     { return "CreateOrderResponse" }
func (*CreateOrderResponse) ProtoMessage()                        {}
func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message { return nil }

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderResponse) GetTax() float64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *CreateOrderResponse) GetFinalPrice() float64 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type ListOrdersRequest struct{}

func (x *ListOrdersRequest) Reset()                             { *x = ListOrdersRequest{} }
func (x *ListOrdersRequest) String() string                     { return "ListOrdersRequest" }
func (*ListOrdersRequest) ProtoMessage()                        {}
func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message { return nil }

type OrderResponse struct {
	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price      float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float64 `protobuf:"fixed64,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float64 `protobuf:"fixed64,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *OrderResponse) Reset()                             { *x = OrderResponse{} }
func (x *OrderResponse) String() string                     { return "OrderResponse" }
func (*OrderResponse) ProtoMessage()                        {}
func (x *OrderResponse) ProtoReflect() protoreflect.Message { return nil }

func (x *OrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderResponse) GetTax() float64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *OrderResponse) GetFinalPrice() float64 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type ListOrdersResponse struct {
	Orders []*OrderResponse `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListOrdersResponse) Reset()                             { *x = ListOrdersResponse{} }
func (x *ListOrdersResponse) String() string                     { return "ListOrdersResponse" }
func (*ListOrdersResponse) ProtoMessage()                        {}
func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message { return nil }

func (x *ListOrdersResponse) GetOrders() []*OrderResponse {
	if x != nil {
		return x.Orders
	}
	return nil
}

// OrderServiceClient interface
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer interface
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations
type UnimplementedOrderServiceServer struct{}

func (*UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func (*UnimplementedOrderServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
