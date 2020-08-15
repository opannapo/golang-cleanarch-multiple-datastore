package base

type ResDefault struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type ResError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
