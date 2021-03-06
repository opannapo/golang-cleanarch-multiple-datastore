package middleware

import (
	"app/app/v1/apis/endpoints/base"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

//ValidateHeaderToken function as middleware to handling user token validation
func ValidateHeaderToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		returnErr := func(msg string) {
			base.OutFailed(context, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)+" "+msg)
			context.Abort()
		}

		bearToken := context.Request.Header.Get("Authorization")

		//validate empty header
		if len(bearToken) == 0 {
			returnErr("")
			return
		}

		//validate length & header structure
		splitHeaderAuth := strings.Split(bearToken, " ")
		if !(len(splitHeaderAuth) == 2) || !(strings.ToLower(splitHeaderAuth[0]) == "bearer") {
			returnErr("")
			return
		}

		//validate jwt
		tokenString := splitHeaderAuth[1]
		key := viper.GetString("jwtKey")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			//fmt.Printf("Token Parse Token  %+v\n", token)
			if token.Method != jwt.GetSigningMethod("HS256") { //Validate Algorithm HS256
				return nil, fmt.Errorf("Invalid Algorithm for token %v ", tokenString)
			}
			return []byte(key), nil
		})

		if err != nil {
			returnErr(err.Error())
		}

		context.Next()
	}
}
