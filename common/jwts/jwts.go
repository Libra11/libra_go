package jwts

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenConfig struct {
	Val    string
	Exp    time.Duration
	Secret string
}

type JwtToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExp    int64  `json:"accessExp"`
	RefreshExp   int64  `json:"refreshExp"`
}

func createToken(config TokenConfig) (string, int64) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": config.Val,
		"exp":   config.Exp,
	})
	tokenExp := time.Now().Add(config.Exp).Unix()
	tokenStr, _ := accessToken.SignedString([]byte(config.Secret))
	return tokenStr, tokenExp
}

func CreateJwtToken(accessConfig TokenConfig, refreshConfig TokenConfig) *JwtToken {
	AccessToken, AccessExp := createToken(accessConfig)
	RefreshToken, RefreshExp := createToken(refreshConfig)
	return &JwtToken{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
		AccessExp:    AccessExp,
		RefreshExp:   RefreshExp,
	}
}
