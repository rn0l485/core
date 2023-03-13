package web

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)



func TestJWT(t *testing.T) {
	secret := "test-secret"
	issuer := "test-issuer"
	userID := "test-user-id"
	loginFrom := "test-login-from"
	expireDuration := time.Hour * 24 // 1 天

	// 测试生成 JWT
	token, err := GenJWT(secret, issuer, userID, loginFrom, expireDuration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 测试解析 JWT
	claims, err := ParseJWT(secret, token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, issuer, claims.Issuer)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, loginFrom, claims.LoginFrom)
	assert.True(t, claims.ExpiresAt.Time().Sub(time.Now()) > 0)
}

func TestJWT_InvalidToken(t *testing.T) {
	secret := "test-secret"
	invalidToken := "invalid-token"

	// 测试解析无效的 JWT
	claims, err := ParseJWT(secret, invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestJWT_ExpiredToken(t *testing.T) {
	secret := "test-secret"
	issuer := "test-issuer"
	userID := "test-user-id"
	loginFrom := "test-login-from"
	expireDuration := time.Nanosecond // 1 纳秒

	// 测试生成过期的 JWT
	token, err := GenJWT(secret, issuer, userID, loginFrom, expireDuration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 测试解析过期的 JWT
	time.Sleep(time.Nanosecond) // 等待 JWT 过期
	claims, err := ParseJWT(secret, token)
	assert.Error(t, err)
	assert.Nil(t, claims)
}


func TestJWT_InvalidSecret(t *testing.T) {
	secret1 := "test-secret-1"
	secret2 := "test-secret-2"
	issuer := "test-issuer"
	userID := "test-user-id"
	loginFrom := "test-login-from"
	expireDuration := time.Hour * 24 // 1 天

	// 测试生成 JWT
	token, err := GenJWT(secret1, issuer, userID, loginFrom, expireDuration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 测试使用错误的密钥解析 JWT
	claims, err := ParseJWT(secret2, token)
	assert.Error(t, err)
	assert.Nil(t, claims)
}