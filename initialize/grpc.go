package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"net"

	"go-temp/global"
	"go-temp/handler"
	"go-temp/proto"
	"go-temp/utils/register/consul"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGrpc(_ *string) {
	zap.S().Info("初始化 grpc ------------------")
	grpcServer := grpc.NewServer()

	// ---- 注册服务 ----
	serverConfig := global.ServerConfig
	proto.RegisterExampleOrderServer(grpcServer, &handler.ExampleOrderServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.RpcPort))

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

	zap.S().Info("grpc 初始化成功 ------------------")

	// 服务注册到 consul
	zap.S().Info("服务注册到 consul ------------------")
	consulInfo := serverConfig.ConsulInfo

	registerClient := consul.NewRegistryClient(consulInfo.Host, consulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registerClient.Register(serverConfig.Host, serverConfig.RpcPort, serverConfig.Name, serverConfig.Tags, serviceId)

	if err != nil {
		zap.S().Panic("rpc 服务注册到 consul 失败:", err.Error())
	}

	zap.S().Info("服务注册到 consul 成功 ------------------")
}
