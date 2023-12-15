package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"libra.com/api/rpc"
	"libra.com/common"
	"libra.com/common/errs"
	userService "libra.com/grpc/service/user"
	"net/http"
	"time"
)

func TokenVerify() func(c *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		// remove Bearer
		tokenStr := c.GetHeader("authorization")
		if tokenStr == "" {
			c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "请先登录"))
			c.Abort()
			return
		}
		token := tokenStr[7:]
		//验证用户是否已经登录
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		member, err := rpc.ClientUser.VerifyToken(ctx, &userService.TokenVerifyMessage{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", member.Id)
		c.Next()
	}
}
