package controllers

import (
	"net/http"
	"time"

	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"backend/config"
	"backend/models"
)

// GetUser retrieves a user by ID
func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id_user = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User tidak ditemukan"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Kesalahan saat mengambil user"})
	}

	return c.JSON(http.StatusOK, user)
}

// GetUsers retrieves all users
func GetUsers(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Kesalahan saat mengambil data pengguna"})
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser updates user information
func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	// Bind data dari request
	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Input tidak valid"})
	}

	// Cari user yang akan diperbarui
	var user models.User
	if err := config.DB.First(&user, "id_user = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Pengguna tidak ditemukan"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Kesalahan saat mengambil data pengguna"})
	}

	// Perbarui data pengguna
	user.Nama = updatedUser.Nama
	user.Email = updatedUser.Email
	user.Role = updatedUser.Role
	user.DiperbaruiPada = time.Now()

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Kesalahan saat memperbarui pengguna"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Pengguna berhasil diperbarui"})
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	// Hapus data pengguna
	if err := config.DB.Where("id_user = ?", id).Delete(&models.User{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Kesalahan saat menghapus pengguna"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Pengguna berhasil dihapus"})
}
