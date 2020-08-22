package endpoints

import (
	"app/app/v1/apis/endpoints/base"
	"app/app/v1/apis/middleware"
	"app/app/v1/apis/param"
	"app/app/v1/injection/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserEndpoint struct {
	services *services.ServiceInjection
}

func (instance *UserEndpoint) getUser(c *gin.Context) {
	mysqlUser := instance.services.MysqlUserService
	req := c.Request.URL.Query()
	id, _ := strconv.Atoi(req.Get("id"))
	result, err := mysqlUser.GetUser(id)
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, result)
	}
}

func (instance *UserEndpoint) getUsers(c *gin.Context) {
	mysqlUser := instance.services.MysqlUserService
	data, err := mysqlUser.GetUsers()
	//data, err := instance.FirestoreUserService.GetUsers()
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}

func (instance *UserEndpoint) addUser(c *gin.Context) {
	var p param.UserCreate
	err := c.ShouldBindJSON(&p)
	if err != nil {
		base.OutFailed(c, http.StatusBadRequest, err.Error())
		return
	}

	if p.Credential == nil {
		base.OutFailed(c, http.StatusBadRequest, "Param Credential Required")
		c.AbortWithStatus(200)
		return
	}

	err = instance.services.MysqlUserService.AddUser(&p)
	if err != nil {
		base.OutFailed(c, 500, err.Error())
		return
	}
	base.OutOk(c, p)
}

func (instance *UserEndpoint) updateUser(c *gin.Context) {

}

func (instance *UserEndpoint) deleteUser(c *gin.Context) {

}

func NewUserEndpoint(g *gin.RouterGroup, services *services.ServiceInjection) {
	instance := &UserEndpoint{
		services: services,
	}

	//No Middleware
	g.POST("user/add", instance.addUser)

	//Auth Middleware only for this request
	gUserAuth := g.Group("", middleware.ValidateHeaderToken())
	{
		gUserAuth.GET("user", instance.getUser)
		gUserAuth.GET("users", instance.getUsers)
		gUserAuth.POST("user/update", instance.updateUser)
		gUserAuth.POST("user/delete", instance.deleteUser)
	}
}
