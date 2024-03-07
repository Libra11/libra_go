package blog

import (
	"context"
	"github.com/gin-gonic/gin"
	"libra.com/api/rpc"
	"libra.com/common"
	"libra.com/common/errs"
	wordService "libra.com/grpc/service/word"
	"net/http"
	"strconv"
	"time"
)

type HandlerWord struct{}

func (b HandlerWord) addWord(c *gin.Context) {
	result := common.Result{}
	var req wordService.AddWordRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	id, err := rpc.ClientWord.AddWord(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(id))
}

func (b HandlerWord) deleteWord(c *gin.Context) {
	result := common.Result{}
	var req wordService.DeleteWordRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ClientWord.DeleteWord(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (b HandlerWord) getWords(c *gin.Context) {
	result := common.Result{}
	var req wordService.GetWordsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientWord.GetWords(ctx, &req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res))
}

func (b HandlerWord) getWordById(c *gin.Context) {
	result := common.Result{}
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	var req = &wordService.GetWordByIdRequest{
		Id: id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientWord.GetWordById(ctx, req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res))
}

func (b HandlerWord) getWordTranslate(c *gin.Context) {
	result := common.Result{}
	word := c.Query("word")
	if word == "" {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	var req = &wordService.GetWordTranslateRequest{
		Word: word,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.ClientWord.GetWordTranslate(ctx, req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(res))
}

func NewHW() *HandlerWord {
	return &HandlerWord{}
}
