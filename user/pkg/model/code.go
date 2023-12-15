package model

import (
	"libra.com/common/errs"
)

var (
	RedisGetError = errs.NewError(1000, "redis错误")
	GormGetError  = errs.NewError(1001, "数据库读取失败")
	GormSetError  = errs.NewError(1002, "数据库存储失败")

	NoLegalMobile      = errs.NewError(4002, "手机号不合法")
	CodeError          = errs.NewError(4003, "验证码错误")
	AccountExist       = errs.NewError(4004, "账号已存在")
	EmailExist         = errs.NewError(4005, "邮箱已存在")
	MobileExist        = errs.NewError(4006, "手机号已存在")
	CaptchaExpire      = errs.NewError(4007, "验证码不存在或已过期")
	AccountAndPwdError = errs.NewError(4008, "账号或密码错误")
	NoLogin            = errs.NewError(4009, "用户未登录")
	NoMember           = errs.NewError(4010, "用户不存在")
)
