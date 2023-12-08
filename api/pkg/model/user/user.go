package user

import (
	"libra.com/api/pkg/model"
	"libra.com/common"
)

type RegisterReq struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

func (r RegisterReq) Verify() error {
	if !common.VerifyEmailFormat(r.Email) {
		return model.NoLegalEmail
	}
	if !common.VerifyMobile(r.Mobile) {
		return model.NoLegalMobile
	}
	if !r.VerifyPassword() {
		return model.PasswordNotConsistent
	}
	return nil
}
