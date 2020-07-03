package routing

import (
	"GoCab/controller"
	"GoCab/model"
	"GoCab/token"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
)

// SignUp router
type SignUp struct {
	userController controller.User
}

// NewSignUpRouter will create new signup router
func NewSignUpRouter(

	userController controller.User,

) SignUp {
	return SignUp{

		userController: userController,
	}

}

// Register will create new routes
func (router SignUp) Register(group *echo.Group) {
	group.POST("/signup", router.signup)
	group.POST("/signin", router.signIn)
	group.POST("/register-cab", router.registerCab)

}

func (router SignUp) signup(context echo.Context) error {

	// create request
	req := new(model.User)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}
	// validate requesr
	if err := validator.New().Struct(req); err != nil {
		log.Error(err)
		return errors.New("Validation Error")
	}

	// create user in database
	if err := router.userController.Create(req); err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
	})

}

func (router SignUp) signIn(context echo.Context) error {
	type Request struct {
		PhoneNumber string `json:"phone_number"`
	}
	req := new(Request)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}

	// check in database if email or phone already exist
	user, err := router.userController.FindUser(req.PhoneNumber)
	if err != nil {
		return errors.New("Unable to Find User")
	}
	token, err := token.CreateMobileToken(req.PhoneNumber)
	if err != nil {
		return errors.New("Unable to create token")
	}

	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "success",
		"AccessToken": token,
		"user":        user,
	})

}

func (router SignUp) registerCab(context echo.Context) error {

	// create request
	req := new(model.Cab)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}
	// validate requesr
	if err := validator.New().Struct(req); err != nil {
		log.Error(err)
		return errors.New("Validation Error")
	}

	// create user in database
	if err := router.userController.RegisterCab(req); err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
	})

}
