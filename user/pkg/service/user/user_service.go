package user

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"libra.com/common"
	"libra.com/common/encrypts"
	"libra.com/common/errs"
	"libra.com/common/jwts"
	"libra.com/grpc/service/user"
	"libra.com/user/config"
	"libra.com/user/internal/dao"
	"libra.com/user/internal/data/member"
	"libra.com/user/internal/repo"
	"libra.com/user/pkg/model"
	"time"
)

type ServiceUser struct {
	user.UnimplementedUserServiceServer
	cache  repo.Cache
	member repo.MemberRepo
}

func New() *ServiceUser {
	return &ServiceUser{
		cache:  dao.Rc,
		member: dao.NewMemberDao(),
	}
}

func (l *ServiceUser) GetCaptcha(ctx context.Context, message *user.CaptchaMessage) (*user.CaptchaResponse, error) {
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
	return &user.CaptchaResponse{Code: code}, nil
}

func (l *ServiceUser) Register(ctx context.Context, message *user.RegisterMessage) (*user.RegisterResponse, error) {
	redisValue, err := l.cache.Get(ctx, "REGISTER_"+message.Mobile)
	if err == redis.Nil {
		return &user.RegisterResponse{}, errs.GrpcErr(model.CaptchaExpire)
	}
	if err != nil {
		zap.L().Error("Register redis search fail", zap.Error(err))
		return &user.RegisterResponse{}, errs.GrpcErr(model.RedisGetError)
	}
	if redisValue != message.Captcha {
		return &user.RegisterResponse{}, errs.GrpcErr(model.CodeError)
	}
	exist, err := l.member.GetMemberByAccount(ctx, message.Name)
	if err != nil {
		zap.L().Error("Register GetMemberByAccount db fail", zap.Error(err))
		return &user.RegisterResponse{}, errs.GrpcErr(model.GormGetError)
	}
	if exist {
		return &user.RegisterResponse{}, errs.GrpcErr(model.AccountExist)
	}
	exist, err = l.member.GetMemberByEmail(ctx, message.Email)
	if err != nil {
		zap.L().Error("Register GetMemberByEmail db fail", zap.Error(err))
		return &user.RegisterResponse{}, errs.GrpcErr(model.GormGetError)
	}
	if exist {
		return &user.RegisterResponse{}, errs.GrpcErr(model.EmailExist)
	}
	exist, err = l.member.GetMemberByMobile(ctx, message.Mobile)
	if err != nil {
		zap.L().Error("Register GetMemberByMobile db fail", zap.Error(err))
		return &user.RegisterResponse{}, errs.GrpcErr(model.GormGetError)
	}
	if exist {
		return &user.RegisterResponse{}, errs.GrpcErr(model.MobileExist)
	}

	pwd := encrypts.Md5(message.Password)
	mem := &member.Member{
		Account:       message.Name,
		Password:      pwd,
		Name:          message.Name,
		Mobile:        message.Mobile,
		Email:         message.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.Normal,
	}
	err = l.member.SaveMember(ctx, mem)
	if err != nil {
		zap.L().Error("register save member db err", zap.Error(err))
		return &user.RegisterResponse{}, errs.GrpcErr(model.GormSetError)
	}
	return &user.RegisterResponse{}, nil
}

func (l *ServiceUser) Login(ctx context.Context, message *user.LoginMessage) (*user.LoginResponse, error) {
	pwd := encrypts.Md5(message.Password)
	mem, err := l.member.FindMember(ctx, message.Name, pwd)
	if mem == nil {
		return &user.LoginResponse{}, model.AccountAndPwdError
	}
	if err != nil {
		zap.L().Error("Login FindMember db fail", zap.Error(err))
		return &user.LoginResponse{}, model.GormGetError
	}

	accessConf := jwts.TokenConfig{
		Val:    message.Name,
		Exp:    time.Duration(config.C.JC.AccessExp * 3600 * 24 * 1000),
		Secret: config.C.JC.AccessSecret,
	}
	refreshConf := jwts.TokenConfig{
		Val:    message.Name,
		Exp:    time.Duration(config.C.JC.RefreshExp * 3600 * 24 * 1000),
		Secret: config.C.JC.RefreshSecret,
	}
	jwtToken := jwts.CreateJwtToken(accessConf, refreshConf)
	return &user.LoginResponse{
		AccessToken:    jwtToken.AccessToken,
		RefreshToken:   jwtToken.RefreshToken,
		AccessTokenExp: jwtToken.AccessExp,
		TokenType:      "bearer",
	}, nil
}
