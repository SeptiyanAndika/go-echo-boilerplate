package utils

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func SuccessResponseMap(ctx echo.Context, data map[string]interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "Success Response",
		"response":  responseData,
	}).Info(ctx.Request().RequestURI)

	return ctx.JSON(http.StatusOK, responseData)
}

func SuccessResponse(ctx echo.Context, data interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "Success Response",
		"response":  responseData,
	}).Info(ctx.Request().RequestURI)

	return ctx.JSON(http.StatusOK, responseData)
}

func ErrorResponse(ctx echo.Context, err error) error {

	responseData := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "Error Response",
		"response":  responseData,
	}).Info(ctx.Request().RequestURI)

	return ctx.JSON(http.StatusBadRequest, responseData)
}

func UnauthorizedResponse(ctx echo.Context) error {

	responseData := map[string]interface{}{
		"success": false,
		"message": "Unauthorized",
	}

	log.WithFields(log.Fields{
		"requestId": ctx.Response().Header().Get(echo.HeaderXRequestID),
		"action":    "Unauthorized Response",
		"response":  responseData,
	}).Info(ctx.Request().RequestURI)

	return ctx.JSON(http.StatusUnauthorized, responseData)
}
