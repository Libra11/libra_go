package jwts

import (
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	// 创建一个 JwtToken
	accessConfig := TokenConfig{
		Val:    "access",
		Exp:    time.Hour,
		Secret: "access_secret",
	}
	refreshConfig := TokenConfig{
		Val:    "refresh",
		Exp:    24 * time.Hour,
		Secret: "refresh_secret",
	}
	jwtToken := CreateJwtToken(accessConfig, refreshConfig, 1)

	// 测试 AccessToken
	token, err := ParseToken(jwtToken.AccessToken, accessConfig.Secret)
	if err != nil {
		t.Errorf("ParseToken failed: %v", err)
	}
	if token != accessConfig.Val {
		t.Errorf("ParseToken returned wrong value: got %v, want %v", token, accessConfig.Val)
	}

	// 测试 RefreshToken
	token, err = ParseToken(jwtToken.RefreshToken, refreshConfig.Secret)
	if err != nil {
		t.Errorf("ParseToken failed: %v", err)
	}
	if token != refreshConfig.Val {
		t.Errorf("ParseToken returned wrong value: got %v, want %v", token, refreshConfig.Val)
	}
}
