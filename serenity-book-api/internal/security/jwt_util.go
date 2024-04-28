package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"serenity-book-api/pkg/global"
	"time"
)

var jwtKey = []byte(global.GConfig.Security.Secret)

type Claims struct {
	SessionId interface{} `json:"sessionId"` // 用户主建id
	jwt.StandardClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(SessionId interface{}) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		SessionId: SessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken 解析 JWT 令牌
func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
