package middleware

import (
	"NeChat/config"
	"NeChat/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"code": "1", "msg": "账户未登录"})
			c.Abort()
			return
		}
		t, err := jwt.ParseWithClaims(token, &models.MyClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ServerConfig.Server.JWTSingedKey), nil
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": "1", "msg": "Login out"})
			c.Abort()
			return
		}
		c.Set("Username", t.Claims.(*models.MyClaim).Username)
	}
}
