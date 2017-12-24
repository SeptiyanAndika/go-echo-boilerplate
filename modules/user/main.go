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
		UserLogic = NewLogic(NewUserRepository())
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

	err, result := utils.RequestHandler(c, new(LoginRequest), UserLogic.Login)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, result)

}
