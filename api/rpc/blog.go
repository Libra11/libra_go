package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"libra.com/api/config"
	"libra.com/common/discovery"
	"libra.com/common/logs"
	blogService "libra.com/grpc/service/blog"
	"log"
)

var ClientBlog blogService.BlogServiceClient

func InitBlogRpcClient() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///blog", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientBlog = blogService.NewBlogServiceClient(conn)
}
