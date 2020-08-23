package base

import (
	"github.com/gin-gonic/gin"
)

//ResDefault type struct Default Response Format
type ResDefault struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

//ResError type struct Default Response Error Format
type ResError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//OutOk is default method to return JsonResponse as status ok
func OutOk(c *gin.Context, data interface{}) {
	c.JSON(200, ResDefault{Data: data, Success: true})
}

//OutFailed is default method to return JsonResponse as failed
func OutFailed(c *gin.Context, code int, err string) {
	c.JSON(200, &ResDefault{Error: ResError{Message: err, Code: code}, Success: false})
}
