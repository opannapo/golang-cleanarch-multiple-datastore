package app

import (
	"app/app/v1/apis/endpoints"
	"app/app/v1/apis/endpoints/base"
	"app/app/v1/injection/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SetupRoute Router Setup
func SetupRoute(services *services.ServiceInjection) {
	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	router.NoRoute(func(context *gin.Context) {
		err := base.ResError{
			Message: "Service Not Found",
			Code:    404,
		}
		res := base.ResDefault{
			Error:   err,
			Data:    nil,
			Success: false,
		}
		context.JSON(200, res)
	})

	v1 := router.Group("api/v1")
	{
		//initial controllers
		endpoints.NewIndexEndpoint(v1)
		endpoints.NewUserEndpoint(v1, services)
		endpoints.NewAuthEndpoint(v1, services)
	}

	_ = router.Run(viper.GetString("server.address"))
}
