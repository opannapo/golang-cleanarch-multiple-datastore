package middleware

import (
	"app/app/base"
	"app/app/v1/apis/param"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCreateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p param.UserCreate
		err := c.ShouldBindJSON(&p)
		if err != nil {
			base.OutFailed(c, http.StatusBadRequest, "err1")
			c.AbortWithStatus(200)
			return
		}

		if p.Credential == nil {
			base.OutFailed(c, http.StatusBadRequest, "Param Credential Required")
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
