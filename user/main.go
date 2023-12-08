package main

import (
	"github.com/gin-gonic/gin"
	srv "libra.com/common"
	"libra.com/user/config"
	"libra.com/user/internal/database/gorms"
	"libra.com/user/router"
)

func main() {
	r := gin.Default()
	println("init main")
	config.C.InitZapLog()
	router.InitRouter(r)
	gorms.InitDB()
	gc := router.InitGRPCServer()
	router.RegisterEtcdServer()
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, gc.Stop)
}
