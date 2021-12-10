package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/r1soX/kissme/errors"
	"github.com/r1soX/kissme/session"
	"github.com/r1soX/kissme/until"
)

// AuthMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := session.GetUser(c)

		if user == nil || user.ID == 0 {
			until.WriteResponse(c, 401, gin.H{
				"result": "fail",
			}, errors.ErrNotAuthenticated)
			c.Abort()
			return
		}

		c.Next()
	}
}
