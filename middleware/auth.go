package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if !session.Get("loggedin").(bool) {
			c.JSON(401, gin.H{"message": "Unauthorized, please input password"})
			c.Abort()
			return
		}
		c.Next()
	}
}
