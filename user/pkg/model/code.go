package model

import (
	"libra.com/common/errs"
)

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
