package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/swlee/Ecommerce-sw/tokens"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Header에 Token 가져오기
		ClientToken := c.Request.Header.Get("token")
		// Token 유무 확인
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "No authrization Header provied"})
			c.Abort()
			return
		}
		// Header Token 검증
		claims, err := tokens.ValidateToken(ClientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
		0
	}
}
