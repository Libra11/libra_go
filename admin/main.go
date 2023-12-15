package main

import (
	"github.com/gin-gonic/gin"
	"libra.com/admin/config"
	"libra.com/admin/internal/database/gorms"
	"libra.com/admin/router"
	srv "libra.com/common"
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
