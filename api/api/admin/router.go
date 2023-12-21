package admin

import (
	"github.com/gin-gonic/gin"
	"libra.com/api/rpc"
)

type RouterAdmin struct {
}

func (*RouterAdmin) Route(r *gin.Engine) {
	rpc.InitAdminRpcClient()
	//h := New()
}
