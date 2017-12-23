package main

import (
	userModules "boilerplate/modules/user"
	"net/http"
	"time"

	"boilerplate/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func LogsEndpoint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		start := time.Now()
		req := c.Request()
		res := c.Response()
		c.Set("RequestURI", req.RequestURI)

		log.WithFields(log.Fields{
			"requestId":   res.Header().Get(echo.HeaderXRequestID),
			"method":      req.Method,
			"timeRequest": start.Format("Mon Jan _2 15:04:05 2006"),
		}).Info(req.RequestURI)
		c.Set("requestURI", req.RequestURI)

		if err = next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()
		status := res.Status

		log.WithFields(log.Fields{
			"requestId": res.Header().Get(echo.HeaderXRequestID),
			"latency":   stop.Sub(start).String(),
			"status":    status,
		}).Info(req.RequestURI)

		return
	}
}

func main() {
	autoMigrate()
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(LogsEndpoint)
	//e.Use(middleware.Logger())
	//	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, utils.Config.App.Name)
	})
	userModules.Routes(e)
	e.Logger.Fatal(e.Start(utils.Config.App.Port))

	defer utils.GetInstanceDB().Db.Close()
}

func autoMigrate() {
	db := utils.GetInstanceDB().Db
	db.AutoMigrate(&userModules.UserSchema{})
}
