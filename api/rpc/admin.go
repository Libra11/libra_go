package admin

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"libra.com/api/config"
	"libra.com/common/discovery"
	"libra.com/common/logs"
	adminService "libra.com/grpc/service/admin"
	"log"
)

var ClientAdmin adminService.AdminServiceClient

func InitAdminRpcClient() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///admin", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientAdmin = adminService.NewAdminServiceClient(conn)
}
