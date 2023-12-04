package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"libra.com/common"
	"libra.com/common/errs"
	loginServiceV1 "libra.com/user/pkg/service/login.service.v1"
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
	mobile := r.PostForm("mobile")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp, err := ClientUser.GetCaptcha(ctx, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		r.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	r.JSON(http.StatusOK, result.Success(rsp.Code))
}

func (*HandlerUser) register(r *gin.Context) {

}
