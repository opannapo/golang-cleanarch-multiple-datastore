package endpoints

import (
	"app/app/v1/apis/endpoints/base"
	"app/app/v1/apis/param"
	"app/app/v1/injection/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthEndpoint struct
type AuthEndpoint struct {
	services *services.ServiceInjection
}

//NewAuthEndpoint create new instance for AuthEndpoint
func NewAuthEndpoint(g *gin.RouterGroup, services *services.ServiceInjection) {
	instance := &AuthEndpoint{services: services}
	g.POST("/auth", instance.doAuth)
}

func (instance *AuthEndpoint) doAuth(c *gin.Context) {
	var p param.AuthParam
	err := c.ShouldBindJSON(&p)
	if err != nil {
		base.OutFailed(c, http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	credential, token, err := instance.services.AuthService.ValidateCredential(&p)
	if err != nil {
		base.OutFailed(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	credential.User.FollowingTopic = nil //omitempty
	credential.User.Phone = ""           //omitempty
	credential.User.Email = ""           //omitempty
	data := map[string]interface{}{
		"user":  credential.User,
		"token": token,
	}
	base.OutOk(c, data)
}
