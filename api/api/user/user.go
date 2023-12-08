package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"libra.com/api/pkg/model/user"
	"libra.com/common"
	"libra.com/common/errs"
	userService "libra.com/grpc/service/user"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(r *gin.Context) {
	result := common.Result{}
	// receive json data
	body := make(map[string]string)
	err := r.ShouldBind(&body)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp, err := ClientUser.GetCaptcha(ctx, &userService.CaptchaMessage{Mobile: body["mobile"]})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(http.StatusOK, result.Success(rsp.Code))
}

func (*HandlerUser) register(r *gin.Context) {
	result := &common.Result{}
	var req user.RegisterReq
	err := r.ShouldBind(&req)
	if err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	//参数校验
	if err := req.Verify(); err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//处理业务
	msg := &userService.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数有误"))
		return
	}
	_, err = ClientUser.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(200, result.Success(nil))
}

func (*HandlerUser) login(r *gin.Context) {
	result := &common.Result{}
	var req user.LoginReq
	err := r.ShouldBind(&req)
	if err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//处理业务
	msg := &userService.LoginMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数有误"))
		return
	}
	var rsp user.LoginResp
	rspMsg := &userService.LoginResponse{}
	rspMsg, err = ClientUser.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	err = copier.Copy(&rsp, rspMsg)
	if err != nil {
		r.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数有误"))
		return
	}
	r.JSON(200, result.Success(rsp))
}
