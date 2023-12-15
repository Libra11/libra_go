package jwts

import (
	"errors"
	"fmt"
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

func createToken(config TokenConfig, memId int64) (string, int64) {
	tokenExp := time.Now().Add(config.Exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": memId,
		"exp":   tokenExp,
	})
	tokenStr, err := accessToken.SignedString([]byte(config.Secret))
	if err != nil {
		panic(err)
	}
	return tokenStr, tokenExp
}

func CreateJwtToken(accessConfig TokenConfig, refreshConfig TokenConfig, memId int64) *JwtToken {
	AccessToken, AccessExp := createToken(accessConfig, memId)
	RefreshToken, RefreshExp := createToken(refreshConfig, memId)
	return &JwtToken{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
		AccessExp:    AccessExp,
		RefreshExp:   RefreshExp,
	}
}

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		t := fmt.Sprintf("%v", claims["token"].(float64))
		var exp = int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return "", errors.New("token已过期")
		}
		return t, nil

	} else {
		return "", err
	}
}
