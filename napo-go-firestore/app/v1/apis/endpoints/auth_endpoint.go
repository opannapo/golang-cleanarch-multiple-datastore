package endpoints

import (
	super "app/app/base"
	"app/app/v1/apis/param"
	"app/app/v1/injection"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthEndpoint struct {
	services *injection.ServiceInjection
}

func NewAuthEndpoint(g *gin.RouterGroup, services *injection.ServiceInjection) {
	instance := &AuthEndpoint{services: services}
	g.POST("/auth", instance.doAuth)
}

func (instance *AuthEndpoint) doAuth(c *gin.Context) {
	var p param.AuthParam
	err := c.ShouldBindJSON(&p)
	if err != nil {
		super.OutFailed(c, http.StatusUnauthorized, err.Error())
		return
	}

	instance.services.MysqlUserService.GetUser()
	super.OutOk(c, p)
}
