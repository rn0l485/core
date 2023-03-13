package web

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rn0l485/core/utility"
)



type PixuClaims struct {
	UserID string `json:"UserID"`
	LoginFrom string `json:"LoginFrom"`
	jwt.RegisteredClaims
}

func GenJWT(secret, Issuer, UserID, LoginFrom string, ExpireDuration time.Duration) (string, error) {
	claims := PixuClaims{
		UserID: UserID,
		LoginFrom: LoginFrom,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: Issuer,
			ID: utility.NewID("jwt"),
		},
	}
	
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func ParseJWT(secret, RawToken string) (*PixuClaims, error) {
	token, err := jwt.ParseWithClaims(RawToken, &PixuClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*PixuClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}