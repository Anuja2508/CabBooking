package server

import (
	"GoCab/token"
	"errors"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// LoginUserMiddleware :-  This middleware is for getting login user information from jwt token
func LoginUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// extract current user's email from access token
		PhoneNumber, err := token.GetPhoneNumberFromJWT(c)
		if err != nil {
			log.Error(err)
			return errors.New("Error while getting Phone number from User Login Token")
		}

		// find user in database
		user, err := UserController.FindUser(PhoneNumber)
		if err != nil {

			log.Error(err)
			return errors.New("error while finding user")

		}
		c.Set("User", user)
		return next(c)
	}
}
