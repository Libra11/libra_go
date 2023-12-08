package model

import (
	"libra.com/common/errs"
)

var (
	NoLegalEmail          = errs.NewError(4001, "邮箱不合法")
	NoLegalMobile         = errs.NewError(4002, "手机号不合法")
	PasswordNotConsistent = errs.NewError(4003, "密码不一致")
)
