package endpoints

import (
	"app/app/v1/apis/endpoints/base"
	"github.com/gin-gonic/gin"
)

func NewIndexEndpoint(g *gin.RouterGroup) {
	g.GET("/", home)
	g.GET("/directory", directory)
}

func home(c *gin.Context) {
	explorer := []string{
		"/me",
		"/directory",
		"/help",
	}
	base.OutOk(c, explorer)
}

func directory(c *gin.Context) {
	explorer := []string{
		"/",
		"/siswa",
		"/siswa/{ID}",
		"/employee",
		"/employee/{ID}",
	}
	base.OutOk(c, explorer)
}
