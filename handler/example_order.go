package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go-temp/proto"
)

type ExampleOrderServer struct {
	proto.UnimplementedExampleOrderServer
}

func (e *ExampleOrderServer) CreateOrder(ctx context.Context, request *proto.ExampleOrderRequest) (*proto.ExampleOrderInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExampleOrderServer) OrderList(ctx context.Context, request *proto.ExampleOrderFilterRequest) (*proto.ExampleOrderListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExampleOrderServer) OrderDetail(ctx context.Context, request *proto.ExampleOrderRequest) (*proto.ExampleOrderInfoDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExampleOrderServer) UpdateOrderStatus(ctx context.Context, status *proto.ExampleOrderStatus) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
