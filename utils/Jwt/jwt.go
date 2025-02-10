package Jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yurongjie2003/ginblog/config"
	"time"
)

// go get github.com/golang-jwt/jwt/v5

// CustomClaims 包含自定义的业务字段
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID uint) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JwtEffectiveTime) * time.Minute)),
			Issuer:    config.JwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}

// ParseToken 解析 Jwt Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	// 处理解析错误
	if err != nil {
		return nil, err // 直接将错误抛给调用方处理
	}

	// 验证Claims有效性
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
