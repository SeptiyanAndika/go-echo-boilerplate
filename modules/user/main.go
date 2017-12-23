package user

import (
	"boilerplate/utils"
	"net/http"
	"sync"

	"github.com/labstack/echo"
)

var UserLogic LogicInterface
var once sync.Once

func init() {
	once.Do(func() {
		UserLogic = NewLogic()
	})
}

func Routes(e *echo.Echo) {
	e.POST("/register", Register)
	e.POST("/login", Login)
	e.GET("/restricted", restricted, utils.Authorizer())
	e.GET("/restricted-user", restricted, utils.Authorizer("user"))
	e.GET("/restricted-admin", restricted, utils.Authorizer("admin"))
	e.GET("/restricted-user-admin", restricted, utils.Authorizer("admin", "user"))
}

func restricted(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome ")
}

func Register(c echo.Context) error {
	// parsing
	request := new(RegisterRequest)
	errParsing := utils.ParsingParamter(c, request)
	if errParsing != nil {
		return utils.ErrorResponse(c, errParsing)
	}
	// validate
	errValidate := utils.ValidateParamter(c, request)
	if errValidate != nil {
		return utils.ErrorResponse(c, errValidate)
	}
	err, result := UserLogic.Register(request.FirstName, request.LastName, request.Email, request.Password)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, result)
}

func Login(c echo.Context) error {
	// parsing
	request := new(LoginRequest)
	errParsing := utils.ParsingParamter(c, request)
	if errParsing != nil {
		return utils.ErrorResponse(c, errParsing)
	}
	// validate
	errValidate := utils.ValidateParamter(c, request)
	if errValidate != nil {
		return utils.ErrorResponse(c, errValidate)
	}
	err, result := UserLogic.Login(request.Email, request.Password)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, result)
}

// func InitRoute(e *echo.Echo) {
// 	logic := NewLogic()

// 	e.GET("/", Register(logic))
// 	// e.GET("/", func(c echo.Context) error {
// 	// 	err, result := logic.Register()

// 	// 	if err != nil {
// 	// 		return utils.ErrorResponse(c, err)
// 	// 	}
// 	// 	return utils.SuccessResponse(c, result)
// 	// })

// 	// e.GET("/plus", func(c echo.Context) error {
// 	// 	// parsing
// 	// 	calc := new(CalculationStruct)
// 	// 	errParsing := utils.ParsingParamter(c, calc)
// 	// 	if errParsing != nil {
// 	// 		return utils.ErrorResponse(c, errParsing)
// 	// 	}
// 	// 	// validate
// 	// 	errValidate := utils.ValidateParamter(calc)
// 	// 	if errValidate != nil {
// 	// 		return utils.ErrorResponse(c, errValidate)
// 	// 	}

// 	// 	// logic
// 	// 	err, result := logic.Plus(calc.ValA, calc.ValB)
// 	// 	if err != nil {
// 	// 		return utils.ErrorResponse(c, err)
// 	// 	}
// 	// 	data := map[string]interface{}{
// 	// 		"result": result,
// 	// 	}
// 	// 	return utils.SuccessResponse(c, data)
// 	// })
// }

// func Register(c echo.Context) error {
// 	err, result := logic.Register("", "", "", "")

// 	if err != nil {
// 		return utils.ErrorResponse(c, err)
// 	}
// 	return utils.SuccessResponse(c, result)

// }
