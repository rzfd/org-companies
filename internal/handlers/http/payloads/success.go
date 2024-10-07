package payloads

import "github.com/labstack/echo/v4"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"Data"`
}

func HandleSuccess(e echo.Context, data interface{}, message string, status int) {
	res := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	e.JSON(status, res)
}
