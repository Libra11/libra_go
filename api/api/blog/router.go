package blog

import (
	"github.com/gin-gonic/gin"
	"libra.com/api/middleware"
	"libra.com/api/rpc"
)

type RouterBlog struct {
}

func (*RouterBlog) Route(r *gin.Engine) {
	rpc.InitBlogRpcClient()
	h := New()
	group := r.Group("/blog")
	group.Use(middleware.TokenVerify()) // 使用TokenVerify中间件
	group.GET("/getAllTags", h.getAllTags)
	group.GET("/getAllCategory", h.getAllCategory)
	group.POST("/deleteTag", h.deleteTag)
	group.POST("/deleteCategory", h.deleteCategory)
	group.POST("/addTag", h.addTag)
	group.POST("/addCategory", h.addCategory)
	group.POST("/addBlog", h.addBlog)
	group.GET("/getOssToken", h.getOssToken)
	group.POST("/getBlogs", h.getBlogs)
}
