package blog

import (
	"context"
	"github.com/gin-gonic/gin"
	"libra.com/api/rpc"
	"libra.com/common"
	"libra.com/common/errs"
	blogService "libra.com/grpc/service/blog"
	"net/http"
	"time"
)

type HandlerBlog struct {
}

func (b HandlerBlog) getAllTags(r *gin.Context) {
	result := common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tags, err := rpc.ClientBlog.GetAllTags(ctx, &blogService.GetAllTagsRequest{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(http.StatusOK, result.Success(tags))
}

func (b HandlerBlog) getAllCategory(r *gin.Context) {
	result := common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	category, err := rpc.ClientBlog.GetAllCategory(ctx, &blogService.GetAllCategoryRequest{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(http.StatusOK, result.Success(category))
}

func (b HandlerBlog) addBlog(r *gin.Context) {
	result := common.Result{}
	var req blogService.AddBlogRequest
	if err := r.ShouldBind(&req); err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ClientBlog.AddBlog(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(http.StatusOK, result.Success(nil))
}

func (b HandlerBlog) deleteCategory(c *gin.Context) {
	result := common.Result{}
	var req blogService.DeleteCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ClientBlog.DeleteCategory(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (b HandlerBlog) deleteTag(c *gin.Context) {
	result := common.Result{}
	var req blogService.DeleteTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ClientBlog.DeleteTag(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (b HandlerBlog) addTag(c *gin.Context) {
	result := common.Result{}
	var req blogService.AddTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientBlog.AddTag(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res.Id))
}

func (b HandlerBlog) addCategory(c *gin.Context) {
	result := common.Result{}
	var req blogService.AddCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientBlog.AddCategory(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res.Id))
}

func (b HandlerBlog) getOssToken(c *gin.Context) {
	result := common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientBlog.GetOssToken(ctx, &blogService.GetOssTokenRequest{})
	println(res)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res))
}

func New() *HandlerBlog {
	return &HandlerBlog{}
}
