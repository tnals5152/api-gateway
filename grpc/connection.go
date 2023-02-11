package grpc

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/utils"
)

func ConnectGrpcClient(host, port string) (conn *grpc.ClientConn, cancel context.CancelFunc, err error) {
	// port := viper.GetString(constant.GatewayGrpcPort)

	target := host + ":" + port

	if host == "localhost" || host == "127.0.0.1" {
		target = ":" + port
	}

	ctx, cancel := utils.GetContext(viper.GetString(constant.GrpcTimeout)) // 외부에서 닫을 수 있게 calcel return

	conn, err = grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials())) //grpc.WithBlock())

	return
}
