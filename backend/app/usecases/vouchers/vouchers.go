package vouchers

import (
	"net/http"
	"voucher_seat/app/model"
	"voucher_seat/app/repositories"
	"voucher_seat/app/repositories/vouchers"
	"voucher_seat/app/usecases"
	"voucher_seat/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type event struct {
	RepoVouchers repositories.Vouchers
	Db           *gorm.DB
}

func New(logger *log.Entry, db *gorm.DB) usecases.Vouchers {
	return &event{
		RepoVouchers: vouchers.New(logger, db),
	}
}

func (v *event) CheckFlightNumber(ctx echo.Context, req model.CheckFlightNumberReq) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: CheckFlightNumber")

	var count int64
	count, e = v.RepoVouchers.CheckFlightNumber(ctx, req)
	if e != nil {
		e = response.SetResponse(ctx, http.StatusBadRequest,
			e.Error(), nil, false)
		return
	}
	if count == 0 {
		e = response.SetResponse(ctx, http.StatusOK, "Success", model.CheckFlightNumberRes{
			Exists: count < 0,
		}, true)
		return
	}

	e = response.SetResponse(ctx, http.StatusOK, "Success", model.CheckFlightNumberRes{
		Exists: count > 0,
	}, true)

	return
}

func (v *event) GenerateSeatNumber(ctx echo.Context, req model.GenerateSeatNumberReq) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: GenerateSeatNumber")

	var data []string
	data, e = v.RepoVouchers.GenerateSeatNumber(ctx, req)
	if e != nil {
		e = response.SetResponse(ctx, http.StatusBadRequest,
			e.Error(), nil, false)
		return
	}

	if len(data) == 0 {
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"vouchers already generated for this flight and date", nil, false)
		return
	}

	e = response.SetResponse(ctx, http.StatusOK, "Success", model.GenerateSeatNumberRes{
		Seats: data,
	}, true)

	return
}
