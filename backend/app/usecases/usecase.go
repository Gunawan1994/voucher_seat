package usecases

import (
	"voucher_seat/app/model"

	"github.com/labstack/echo/v4"
)

type Vouchers interface {
	CheckFlightNumber(ctx echo.Context, req model.CheckFlightNumberReq) (e error)
	GenerateSeatNumber(ctx echo.Context, req model.GenerateSeatNumberReq) (e error)
}
