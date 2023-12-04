package main

import (
	"github.com/gin-gonic/gin"
	"libra.com/api/config"
	"libra.com/api/router"
	srv "libra.com/common"
)

func main() {
	r := gin.Default()
	println("init main")
	config.C = config.InitConfig()
	config.C.InitZapLog()
	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
