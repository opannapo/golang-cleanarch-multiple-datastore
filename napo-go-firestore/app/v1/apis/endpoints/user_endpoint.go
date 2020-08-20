package endpoints

import (
	super "app/app/base"
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
		super.OutFailed(c, 0, err.Error())
	} else {
		super.OutOk(c, result)
	}
}

func (instance *UserEndpoint) getUsers(c *gin.Context) {
	mysqlUser := instance.services.MysqlUserService
	data, err := mysqlUser.GetUsers()
	//data, err := instance.FirestoreUserService.GetUsers()
	if err != nil {
		super.OutFailed(c, 0, err.Error())
	} else {
		super.OutOk(c, data)
	}
}

func (instance *UserEndpoint) addUser(c *gin.Context) {
	var p param.UserCreate
	err := c.ShouldBindJSON(&p)
	if err != nil {
		super.OutFailed(c, http.StatusBadRequest, err.Error())
		return
	}

	err = instance.services.MysqlUserService.AddUser(&p)
	if err != nil {
		super.OutFailed(c, 500, err.Error())
		return
	}
	super.OutOk(c, p)
}

func (instance *UserEndpoint) updateUser(c *gin.Context) {

}

func (instance *UserEndpoint) deleteUser(c *gin.Context) {

}

func NewUserEndpoint(g *gin.RouterGroup, services *services.ServiceInjection) {
	instance := &UserEndpoint{
		services: services,
	}
	g.GET("user", instance.getUser)
	g.GET("users", instance.getUsers)
	g.POST("user/add", instance.addUser)
	g.POST("user/update", instance.updateUser)
	g.POST("user/delete", instance.deleteUser)
}
