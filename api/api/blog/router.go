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
	hw := NewHW()
	group := r.Group("/blog")
	group.Use(middleware.TokenVerify()) // 使用TokenVerify中间件
	group.POST("/deleteTag", h.deleteTag)
	group.POST("/deleteCategory", h.deleteCategory)
	group.POST("/addTag", h.addTag)
	group.POST("/addCategory", h.addCategory)
	group.POST("/addBlog", h.addBlog)
	group.GET("/getOssToken", h.getOssToken)
	group.POST("/deleteBlog", h.deleteBlog)
	group.POST("/addWord", hw.addWord)
	group.POST("/deleteWord", hw.deleteWord)
	group.GET("/getWordTranslate", hw.getWordTranslate)
	group2 := r.Group("/blog")
	group2.GET("/getAllTags", h.getAllTags)
	group2.GET("/getAllCategory", h.getAllCategory)
	group2.POST("/getBlogs", h.getBlogs)
	group2.GET("/getBlogById", h.getBlogById)
	group2.POST("/getWords", hw.getWords)
	group2.GET("/getWordById", hw.getWordById)
}
