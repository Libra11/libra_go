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
	r.POST("/user/getCaptcha", h.getCaptcha)
	r.POST("/user/register", h.register)
	r.POST("/user/login", h.login)
	group := r.Group("/user/getUserInfo")
	group.Use(middleware.TokenVerify())
	group.GET("/", h.getUserInfo)
}
