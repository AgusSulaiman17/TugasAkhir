package routes

import (
    "backend/controllers"
    "backend/middleware" // Pastikan middleware sudah diimpor
    "github.com/labstack/echo/v4"
)

func PenulisRoutes(e *echo.Echo) {
    // Hanya admin yang bisa melakukan operasi CRUD pada penulis
    adminGroup := e.Group("/penulis")
    adminGroup.Use(middleware.AdminMiddleware) // Menggunakan middleware admin

    adminGroup.POST("", controllers.CreatePenulis)         // Menambahkan penulis
    adminGroup.GET("/:id", controllers.GetPenulis)        // Mendapatkan penulis berdasarkan ID
    adminGroup.PUT("/:id", controllers.UpdatePenulis)     // Memperbarui penulis berdasarkan ID
    adminGroup.DELETE("/:id", controllers.DeletePenulis)  // Menghapus penulis berdasarkan ID
    e.GET("/penulis", middleware.JWTMiddleware(controllers.GetAllPenulis))
}
