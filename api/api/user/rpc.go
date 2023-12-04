package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"libra.com/api/config"
	"libra.com/common/discovery"
	"libra.com/common/logs"
	loginServiceV1 "libra.com/user/pkg/service/login.service.v1"
	"log"
)

var ClientUser loginServiceV1.LoginServiceClient

func InitUserRpcClient() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientUser = loginServiceV1.NewLoginServiceClient(conn)
}
