package user

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

type (
	LogicInterface interface {
		Register(firstName, lastName, email, password string) (error, interface{})
		Login(ctx echo.Context, params interface{}) (error, interface{})
		ForgotPassword(email string) (error, map[string]interface{})
		Activated(token string) (error, map[string]interface{})
	}
	Logic struct {
		userRepo UserRepositoryInterface
	}
)

func NewLogic(userRepo UserRepositoryInterface) LogicInterface {
	return &Logic{userRepo: userRepo}
}

func (l *Logic) Register(firstName, lastName, email, password string) (error, interface{}) {
	hash, _ := l.hashPassword(password)
	err, user := l.userRepo.Create(firstName, lastName, email, hash)
	return err, user
}

func (l *Logic) Login(ctx echo.Context, params interface{}) (error, interface{}) {
	paramater := params.(*LoginRequest)
	user := UserSchema{}
	var err error
	err, user = l.userRepo.FindByEmail(paramater.Email)
	if err != nil {
		return err, nil
	}
	if !l.checkPasswordHash(paramater.Password, user.Password) {
		return errors.New("please check again username or password"), nil
	}
	token, err := l.createToken(user)
	if err != nil {
		return err, nil
	}

	data := map[string]interface{}{
		"token": token,
	}
	return nil, data
}

func (l *Logic) ForgotPassword(email string) (error, map[string]interface{}) {
	data := map[string]interface{}{
		"message": "hello",
	}
	return nil, data
}

func (l *Logic) Activated(token string) (error, map[string]interface{}) {
	data := map[string]interface{}{
		"message": "hello",
	}
	return nil, data
}

func (l *Logic) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (l *Logic) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (l *Logic) createToken(user UserSchema) (t string, err error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["firstname"] = user.FirstName
	claims["lastname"] = user.LastName
	claims["email"] = user.Email
	claims["roles"] = user.Roles
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte("secret"))
	return
}
