package initialize

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpc() {
	conn, err := grpc.Dial("localhost:38080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
}
