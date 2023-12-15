package admin

import "github.com/gin-gonic/gin"

type RouterAdmin struct {
}

func (*RouterAdmin) Route(r *gin.Engine) {
	InitAdminRpcClient()
	//h := New()
}
