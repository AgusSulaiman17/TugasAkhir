package controllers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strconv"

    "backend/config"
    "backend/models"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

// CreateBuku untuk menambahkan buku baru
// CreateBuku untuk menambahkan buku baru
func CreateBuku(c echo.Context) error {
    var buku models.Buku

    // Cek apakah judul buku sudah ada
    judul := c.FormValue("judul")
    var existingBuku models.Buku
    if err := config.DB.Where("judul = ?", judul).First(&existingBuku).Error; err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Book with this title already exists"})
    }

    file, err := c.FormFile("gambar")
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "No image provided"})
    }

    // Simpan gambar
    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to open image"})
    }
    defer src.Close()

    // Pastikan folder "uploads" ada
    if _, err := os.Stat("uploads"); os.IsNotExist(err) {
        os.Mkdir("uploads", os.ModePerm)
    }

    dst := filepath.Join("uploads", file.Filename)
    out, err := os.Create(dst)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save image"})
    }
    defer out.Close()

    if _, err = io.Copy(out, src); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to copy image"})
    }

    buku.Judul = judul
    buku.IdPenulis, _ = strconv.Atoi(c.FormValue("id_penulis"))
    buku.IdGenre, _ = strconv.Atoi(c.FormValue("id_genre"))
    buku.Deskripsi = c.FormValue("deskripsi")
    buku.Jumlah, _ = strconv.Atoi(c.FormValue("jumlah"))
    buku.Gambar = "/uploads/" + file.Filename // Path gambar lokal
    buku.Status = true

    if err := config.DB.Create(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create book"})
    }

    // Menambahkan preload untuk relasi Penulis dan Genre
    if err := config.DB.Preload("Penulis").Preload("Genre").First(&buku, buku.ID).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to load related Penulis or Genre"})
    }

    return c.JSON(http.StatusCreated, buku)
}

// GetBuku untuk mendapatkan daftar semua buku
func GetBuku(c echo.Context) error {
    var bukus []models.Buku
    if err := config.DB.Preload("Penulis").Preload("Genre").Find(&bukus).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve books", "error": err.Error()})
    }
    return c.JSON(http.StatusOK, bukus)
}

// GetBukuByID untuk mendapatkan buku berdasarkan ID
func GetBukuByID(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID format"})
    }

    var buku models.Buku
    if err := config.DB.Preload("Penulis").Preload("Genre").First(&buku, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve book", "error": err.Error()})
    }

    return c.JSON(http.StatusOK, buku)
}

// UpdateBuku untuk memperbarui informasi buku
func UpdateBuku(c echo.Context) error {
    id := c.Param("id")
    var buku models.Buku

    // Mencari buku berdasarkan ID
    if err := config.DB.First(&buku, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve book", "error": err.Error()})
    }

    // Memeriksa apakah ada buku lain dengan judul yang sama (kecuali yang sedang diupdate)
    var existingBuku models.Buku
    if err := config.DB.Where("judul = ? AND id_buku != ?", buku.Judul, id).First(&existingBuku).Error; err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Judul buku sudah ada"})
    }

    // Memperbarui informasi buku jika ada perubahan
    if c.FormValue("judul") != "" {
        buku.Judul = c.FormValue("judul")
    }
    if c.FormValue("id_penulis") != "" {
        buku.IdPenulis, _ = strconv.Atoi(c.FormValue("id_penulis"))
    }
    if c.FormValue("id_genre") != "" {
        buku.IdGenre, _ = strconv.Atoi(c.FormValue("id_genre"))
    }
    if c.FormValue("deskripsi") != "" {
        buku.Deskripsi = c.FormValue("deskripsi")
    }
    if c.FormValue("jumlah") != "" {
        buku.Jumlah, _ = strconv.Atoi(c.FormValue("jumlah"))
    }

    // Memeriksa dan mengupdate gambar jika ada
    file, err := c.FormFile("gambar")
    if err == nil {
        src, err := file.Open()
        if err == nil {
            defer src.Close()
            dst := filepath.Join("uploads", file.Filename)
            out, err := os.Create(dst)
            if err == nil {
                defer out.Close()
                io.Copy(out, src)
                buku.Gambar = "/uploads/" + file.Filename
            } else {
                return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save image"})
            }
        } else {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to open image"})
        }
    }

    // Menyimpan perubahan buku
    if err := config.DB.Save(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update book"})
    }

    // Menambahkan preload untuk relasi Penulis dan Genre setelah update
    if err := config.DB.Preload("Penulis").Preload("Genre").First(&buku, buku.ID).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to load related Penulis or Genre"})
    }

    return c.JSON(http.StatusOK, buku)
}

// DeleteBuku untuk menghapus buku berdasarkan ID
func DeleteBuku(c echo.Context) error {
    id := c.Param("id")
    var buku models.Buku

    if err := config.DB.First(&buku, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve book", "error": err.Error()})
    }

    if err := config.DB.Delete(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete book", "error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Book deleted"})
}
