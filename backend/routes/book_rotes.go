package routes

import (
    "github.com/labstack/echo/v4"
    "backend/controllers"
    "backend/middleware"
)

func BookRoutes(e *echo.Echo) {
    g := e.Group("/buku")
    g.Use(middleware.AdminMiddleware) // Hanya admin yang dapat mengakses route ini
    g.POST("", controllers.CreateBuku)
    g.PUT("/:id", controllers.UpdateBuku)
    g.DELETE("/:id", controllers.DeleteBuku)

    e.GET("/buku/:id", middleware.JWTMiddleware(controllers.GetBukuByID))
    e.GET("/buku", middleware.JWTMiddleware(controllers.GetBuku))
}
