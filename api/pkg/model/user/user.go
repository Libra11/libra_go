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

type LoginReq struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

//type Member struct {
//	Id            int64  `json:"id"`
//	Name          string `json:"name"`
//	Mobile        string `json:"mobile"`
//	RealName      string `json:"realname"`
//	Account       string `json:"account"`
//	Status        int32  `json:"status"`
//	LastLoginTime int64  `json:"lastLoginTime"`
//	Address       string `json:"address"`
//	Province      int32  `json:"province"`
//	City          int32  `json:"city"`
//	Area          int32  `json:"area"`
//	Email         string `json:"email"`
//}

type LoginResp struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenType      string `json:"tokenType"`
	AccessTokenExp int64  `json:"accessTokenExp"`
}
