package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	InitUserRpcClient()
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/getCaptcha", h.register)
}
