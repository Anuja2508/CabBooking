package server

import (
	"GoCab/controller"
	"GoCab/model"
	"GoCab/routing"
	"GoCab/token"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server struct
type Server struct {
	*echo.Echo
}

var (
	UserController controller.User
	Cabcontroller  controller.Cab
)

// New Server
func New(database *gorm.DB) (*Server, error) {
	server := Server{echo.New()}

	// Middleware
	server.HTTPErrorHandler = func(err error, context echo.Context) {
		message := err.Error()
		code := context.Response().Status
		context.JSON(code, map[string]map[string]interface{}{ // sub level mapping
			"error": {
				"message": message,
			},
		})
	}

	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	accessKey, err := token.CreateAccessKey()
	if err != nil {
		return nil, err
	}

	u := new(model.UserClaims)

	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/auth") {
				return true
			}
			return false
		},
		Claims:     u,
		SigningKey: []byte(accessKey),
	}))

	// Initializing Controller
	UserController = controller.NewUserController(database)
	Cabcontroller = controller.NewCabController(database)

	// initializing cab router
	signupRouter := routing.NewSignUpRouter(UserController)
	cabRouter := routing.NewCabRouter(Cabcontroller)

	signupRouter.Register(server.Group("/auth"))
	cabRouter.Register(server.Group("/cab", LoginUserMiddleware))

	return &server, nil
}
