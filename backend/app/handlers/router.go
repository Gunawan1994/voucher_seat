package handlers

import (
	"net/http"

	"voucher_seat/app/middlewares"

	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type Routes struct {
	Db *gorm.DB
}

func NewRoutes(db *gorm.DB) *Routes {
	return &Routes{
		Db: db,
	}
}

func (route *Routes) RegisterServices(c *echo.Echo) {
	logger := log.WithFields(log.Fields{
		"job":    "RegisterServices",
		"msg_id": xid.New().String(),
	})
	logger.Debug("Running")

	handler := Handler(logger, route.Db)
	router := c.Group("/api")

	route.setMiddleware(router)
	router.POST("/check", handler.CheckFlightNumberHandler)
	router.POST("/generate", handler.GenerateSeatNumberHandler)
}

func (route *Routes) setMiddleware(rGroup *echo.Group) {
	rGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRealIP},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))

	m := middlewares.New("")
	rGroup.Use(m.AddLoggerToContext, m.DumpRequest)

}
