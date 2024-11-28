package routes

import (
    "github.com/labstack/echo/v4"
    "backend/controllers"
    "backend/middleware"
)

func UserRoutes(e *echo.Echo) {
    e.GET("/user/:id", middleware.JWTMiddleware(controllers.GetUser))
    e.GET("/users", middleware.JWTMiddleware(controllers.GetUsers))
    e.PUT("/user/:id", middleware.JWTMiddleware(controllers.UpdateUser))
    e.DELETE("/user/:id", middleware.JWTMiddleware(controllers.DeleteUser))
}
