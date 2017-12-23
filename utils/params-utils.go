package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"

	log "github.com/sirupsen/logrus"
)

func ParsingParamter(ctx echo.Context, i interface{}, logsInfo ...string) error {
	err := ctx.Bind(i)

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "parsing",
		"paramater": i,
		"err":       err,
	}).Info(ctx.Request().RequestURI)

	return err
}

func ValidateParamter(ctx echo.Context, i interface{}) error {
	_, err := govalidator.ValidateStruct(i)

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "Validate",
		"input":     i,
		"err":       err,
	}).Info(ctx.Request().RequestURI)

	return err
}
