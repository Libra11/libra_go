package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"libra.com/blog/config"
	blogService "libra.com/blog/pkg/service/blog"
	wordService "libra.com/blog/pkg/service/word"
	"libra.com/common/discovery"
	"libra.com/common/logs"
	"libra.com/grpc/service/blog"
	"libra.com/grpc/service/word"
	"log"
	"net"
)

type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Route(r)
}

func InitRouter(r *gin.Engine) {

}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func InitGRPCServer() *grpc.Server {
	c := &gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(s *grpc.Server) {
			blog.RegisterBlogServiceServer(s, blogService.New())
			word.RegisterWordServiceServer(s, wordService.New())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", config.C.GC.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	log.Println("grpc server start at", config.C.GC.Addr)
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight,
	}
	r := discovery.NewRegister(config.C.EC.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
