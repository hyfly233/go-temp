package initialize

import (
	"fmt"
	"go-temp/handler"
	"go-temp/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"go-temp/global"
)

func InitGrpc(_ *string) {
	grpcServer := grpc.NewServer()

	// ---- 注册服务 ----
	proto.RegisterExampleOrderServer(grpcServer, &handler.ExampleOrderServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.GrpcConfig.Ip, global.GrpcConfig.Port))

	if err != nil {
		panic("grpc failed to listen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	//启动服务
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
}
