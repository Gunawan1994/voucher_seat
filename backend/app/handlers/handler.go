package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"voucher_seat/app/model"
	"voucher_seat/app/usecases"
	vouchers "voucher_seat/app/usecases/vouchers"
	"voucher_seat/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	ilog "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HTTP struct {
	usecaseVouchers usecases.Vouchers
}

func Handler(logger *ilog.Entry, db *gorm.DB) *HTTP {
	return &HTTP{
		usecaseVouchers: vouchers.New(logger, db),
	}
}

func (c *HTTP) CheckFlightNumberHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: CheckFlightNumberHandler")

	req := model.CheckFlightNumberReq{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, false)
		return
	}

	validate := validator.New()
	if e = validate.Struct(&req); e != nil {
		errs := e.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			logger.WithField("error", e.Error()).Error(fmt.Printf("field %s: %s\n", fieldErr.Field(), fieldErr.Tag()))
			e = response.SetResponse(ctx, http.StatusBadRequest,
				fmt.Sprintf("Missing mandatory parameter %s", fieldErr.Field()), nil, false)
			return
		}
		return
	}

	if _, err := time.Parse("2006-01-02", req.Date); err != nil {
		return errors.New("invalid date format; expected YYYY-MM-DD")
	}
	e = c.usecaseVouchers.CheckFlightNumber(ctx, req)

	return
}

func (c *HTTP) GenerateSeatNumberHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: GenerateSeatNumberHandler")

	req := model.GenerateSeatNumberReq{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, false)
		return
	}

	validate := validator.New()
	if e = validate.Struct(&req); e != nil {
		errs := e.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			logger.WithField("error", e.Error()).Error(fmt.Printf("field %s: %s\n", fieldErr.Field(), fieldErr.Tag()))
			e = response.SetResponse(ctx, http.StatusBadRequest,
				fmt.Sprintf("Missing mandatory parameter %s", fieldErr.Field()), nil, false)
			return
		}
		return
	}

	if _, err := time.Parse("2006-01-02", req.Date); err != nil {
		return errors.New("invalid date format; expected YYYY-MM-DD")
	}

	e = c.usecaseVouchers.GenerateSeatNumber(ctx, req)

	return
}
