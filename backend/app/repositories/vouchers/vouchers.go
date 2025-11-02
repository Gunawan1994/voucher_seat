package vouchers

import (
	"errors"
	"math/rand"
	"time"
	"voucher_seat/app/domain"
	"voucher_seat/app/model"
	"voucher_seat/app/repositories"
	"voucher_seat/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type event struct {
	Logger *log.Entry
	Db     *gorm.DB
}

func New(logger *log.Entry, db *gorm.DB) repositories.Vouchers {
	return &event{
		Logger: logger,
		Db:     db,
	}
}

func (v *event) CheckFlightNumber(c echo.Context, req model.CheckFlightNumberReq) (count int64, e error) {
	logger := c.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: CheckFlightNumber")

	v.Db.Model(&domain.Assignments{}).
		Where("flight_number = ? AND flight_date = ?", req.FlightNumber, req.Date).
		Count(&count)

	return
}

func (v *event) GenerateSeatNumber(c echo.Context, req model.GenerateSeatNumberReq) (data []string, e error) {
	logger := c.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: GenerateSeatNumber")

	tx := v.Db.Begin()
	defer tx.Rollback()

	var existing int64
	tx.Model(&domain.Assignments{}).
		Where("flight_number = ? AND flight_date = ?", req.FlightNumber, req.Date).
		Count(&existing)

	if existing > 0 {
		tx.Rollback()
		return nil, errors.New("vouchers already generated for this flight and date")
	}

	var listAllSeat []string
	listAllSeat, e = utils.SeatListForAircraft(req.Aircraft)
	if e != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	data, e = utils.PickRandomSeats(listAllSeat, 3)
	if e != nil {
		return
	}

	if e = tx.Create(&domain.Assignments{
		CrewName:     req.Name,
		CrewID:       req.ID,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.Aircraft,
		Seat1:        data[0],
		Seat2:        data[1],
		Seat3:        data[2],
	}).Error; e != nil {
		return
	}

	if e = tx.Commit().Error; e != nil {
	}

	return
}
