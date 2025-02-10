package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/utils/Jwt"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, results.Error(codes.ErrorTokenNeed))
			c.Abort()
			return
		}
		claims, err := Jwt.ParseToken(token)
		if err != nil {
			var code codes.Code
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed):
				code = codes.ErrorTokenFormatWrong
			case errors.Is(err, jwt.ErrTokenExpired):
				code = codes.ErrorTokenExpired
			case errors.Is(err, jwt.ErrTokenUnverifiable):
				code = codes.ErrorTokenWrong
			default:
				code = codes.Error
			}
			c.JSON(http.StatusOK, results.Error(code))
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
