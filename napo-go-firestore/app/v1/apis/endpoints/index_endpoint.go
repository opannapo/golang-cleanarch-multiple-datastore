package endpoints

import (
	super "app/app/base"
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
	super.OutOk(c, explorer)
}

func directory(c *gin.Context) {
	explorer := []string{
		"/",
		"/siswa",
		"/siswa/{ID}",
		"/employee",
		"/employee/{ID}",
	}
	super.OutOk(c, explorer)
}
