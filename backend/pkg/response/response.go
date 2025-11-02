package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(ctx echo.Context, httpstatus int, msg string, data interface{}, status bool) error {
	return ctx.JSON(httpstatus, Response{
		Success: status,
		Msg:     msg,
		Data:    data,
	})
}
