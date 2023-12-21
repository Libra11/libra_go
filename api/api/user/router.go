package user

import (
	"github.com/gin-gonic/gin"
	"libra.com/api/middleware"
	"libra.com/api/rpc"
)

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	rpc.InitUserRpcClient()
	h := New()
	// 路由组1: 不需要token验证
	group1 := r.Group("/user")
	group1.POST("/getCaptcha", h.getCaptcha)
	group1.POST("/register", h.register)
	group1.POST("/login", h.login)

	// 路由组2: 需要token验证
	group2 := r.Group("/user")
	group2.Use(middleware.TokenVerify()) // 使用TokenVerify中间件
	group2.GET("/getUserInfo", h.getUserInfo)
}
