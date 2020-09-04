package endpoints

import (
	"app/app/v1/apis/endpoints/base"
	"app/app/v1/apis/middleware"
	"app/app/v1/injection/services"
	"github.com/gin-gonic/gin"
)

//TopicTypeEndpoint struct for Topic Type endpoint
type TopicTypeEndpoint struct {
	services *serviceinjection.ServiceInjection
}

//NewTopicTypeEndpoint new instance of NewTopicTypeEndpoint
func NewTopicTypeEndpoint(g *gin.RouterGroup, injection *serviceinjection.ServiceInjection) {
	instance := &TopicTypeEndpoint{services: injection}
	gUserAuth := g.Group("", middleware.ValidateHeaderToken())
	{
		gUserAuth.GET("topic_types", instance.getTopicTypes)
	}
}

func (instance *TopicTypeEndpoint) getTopicTypes(c *gin.Context) {
	result, err := instance.services.TopicTypeService.GetAll()
	if err != nil {
		base.OutFailed(c, 400, err.Error())
	}
	base.OutOk(c, result)
}
