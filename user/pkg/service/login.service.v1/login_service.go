package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"libra.com/common"
	"libra.com/common/errs"
	"libra.com/user/internal/dao"
	"libra.com/user/internal/repo"
	"libra.com/user/pkg/model"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (l *LoginService) GetCaptcha(ctx context.Context, message *CaptchaMessage) (*CaptchaResponse, error) {
	//1. 获取参数
	mobile := message.Mobile
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcErr(model.NoLegalMobile)
	}
	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("调用短信平台发送短信")
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := l.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error("存入redis失败" + err.Error())
		}
	}()
	return &CaptchaResponse{Code: code}, nil
}
