package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

type TokenMiddleware struct {
	Key       string
	Name      string
}

type TokenClaims struct {
	UserInfo      string  `json:"userInfo"`
	jwt.StandardClaims
}

// 验证 token 中间件函数
func (tm TokenMiddleware) Handle(c *gin.Context) {
	tokenString := c.GetHeader("token")

	if tokenString == "" {
		c.AbortWithStatusJSON(200, gin.H {"status": 401, "msg": "当前没有权限查看"})
		return
	}

	userInfoString, isPassValidated := tm.ValidateToken(tokenString)

	if !isPassValidated {
		c.AbortWithStatusJSON(200, gin.H {"status": 401, "msg": "token 解析错误"})
		return
	}

	c.Set("userInfo", userInfoString)

	c.Next()
}

func (tm TokenMiddleware) SignToken(userInfoString string) (string, error) {
	claims := TokenClaims{
		UserInfo: userInfoString,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: int64(time.Now().Add(time.Hour * 24).Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, error := token.SignedString([]byte(tm.Key))

	if error != nil {
		return "", error
	}

	return tokenString, nil
}

func (tm TokenMiddleware) ValidateToken(tokenString string) (string, bool) {
	token, error := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.Key), nil
	})

	if error != nil {
		return "", false
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.UserInfo, true
	}

	return "", false
}

func NewTokenMiddleware() TokenMiddleware {
	return TokenMiddleware { Key: "token-middleware", Name: "token" }
}
