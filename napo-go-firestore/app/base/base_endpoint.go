package base

import (
	"github.com/gin-gonic/gin"
)

type ResDefault struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type ResError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func OutOk(c *gin.Context, data interface{}) {
	c.JSON(200, ResDefault{Data: data, Success: true})
}

func OutFailed(c *gin.Context, code int, err string) {
	c.JSON(200, &ResDefault{Error: ResError{Message: err, Code: code}, Success: false})
}
