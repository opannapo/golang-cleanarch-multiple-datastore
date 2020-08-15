package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func OutOk(c *gin.Context, data interface{}) {
	c.JSON(200, ResDefault{Data: data, Success: true})
}

func OutFailed(c *gin.Context, code int, err string) {
	if strings.ContainsAny(err, "record found not") {
		code = http.StatusNoContent
	}

	c.JSON(200, &ResDefault{Error: ResError{Message: err, Code: code}, Success: false})
}
