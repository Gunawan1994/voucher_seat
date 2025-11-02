package repositories

import (
	"voucher_seat/app/model"

	"github.com/labstack/echo/v4"
)

type Vouchers interface {
	CheckFlightNumber(c echo.Context, req model.CheckFlightNumberReq) (count int64, e error)
	GenerateSeatNumber(c echo.Context, req model.GenerateSeatNumberReq) (data []string, e error)
}
