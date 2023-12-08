package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"libra.com/api/config"
	"libra.com/common/discovery"
	"libra.com/common/logs"
	userService "libra.com/grpc/service/user"
	"log"
)

var ClientUser userService.UserServiceClient

func InitUserRpcClient() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientUser = userService.NewUserServiceClient(conn)
}
