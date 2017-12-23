package utils

import (
	"github.com/labstack/echo"
)

type Logic func(i interface{}) (error, interface{})

func RequestHandler(ctx echo.Context, i interface{}, logic Logic) (err error, result interface{}) {

	// execute parsing
	err = ParsingParamter(ctx, i)

	// execute validate
	err = ValidateParamter(ctx, i)

	// execute logic
	err, result = logic(i)
	return
}
