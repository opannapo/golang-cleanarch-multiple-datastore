package middleware

import (
	"app/app/base"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateHeaderToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearToken := context.Request.Header.Get("Authorization")
		fmt.Println("auth bearToken : " + bearToken)

		returnErr := func() {
			base.OutFailed(context, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			context.Abort()
		}

		if len(bearToken) == 0 {
			returnErr()
			return
		}

		splitHeaderAuth := strings.Split(bearToken, " ")
		if !(len(splitHeaderAuth) == 2) || !(strings.ToLower(splitHeaderAuth[0]) == "bearer") {
			returnErr()
			return
		}

		context.Next()
	}
}
