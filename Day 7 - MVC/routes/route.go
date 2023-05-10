package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {

	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByID)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	return e

}
