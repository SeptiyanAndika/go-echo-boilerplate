package utils

import (
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Authorizer(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			tokenString := ""
			if len(req.Header["Authorization"]) != 0 {
				tokenString = req.Header["Authorization"][0]
			} else {
				return ErrorResponse(c, errors.New("Authorization not found"))
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte("secret"), nil
			})

			if err != nil {
				return ErrorResponse(c, err)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				rolesUserString := claims["roles"].(string)
				rolesUser := strings.Split(rolesUserString, ",")
				if len(roles) == 0 {
					c.Set("user", claims)
				} else if len(rolesUser) > 0 && intersectRoles(rolesUser, roles) {
					c.Set("user", claims)
				} else {
					return UnauthorizedResponse(c)
				}
			} else {
				return UnauthorizedResponse(c)
			}

			return next(c)
		}
	}
}

func intersectRoles(roles1 []string, roles2 []string) bool {
	for _, role1 := range roles1 {
		for _, role2 := range roles2 {
			if role1 == role2 {
				return true
			}
		}
	}
	return false
}
