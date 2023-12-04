package router

import (
	"github.com/gin-gonic/gin"
	"libra.com/api/api/user"
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
	rg := New()
	rg.Route(&user.RouterUser{}, r)
}
