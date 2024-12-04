package routes

import (
    "github.com/labstack/echo/v4"
    "backend/controllers"
)

func UserRoutes(e *echo.Echo) {
    e.GET("/user/:id", controllers.GetUser)
    e.GET("/users", controllers.GetUsers)
    e.PUT("/user/:id", controllers.UpdateUser)
    e.DELETE("/user/:id", controllers.DeleteUser)
    e.POST("/user",controllers.Register)
}
