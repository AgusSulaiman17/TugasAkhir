package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"strconv"
	"backend/utils"
	"backend/config"
	"backend/models"
)


// Create Peminjaman
func CreatePeminjaman(c echo.Context) error {
    var peminjaman models.Peminjaman

    // Bind data dari request
    if err := c.Bind(&peminjaman); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format data tidak valid"})
    }

    // Validasi data tidak boleh kosong
    if peminjaman.IDBuku == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID buku tidak boleh kosong"})
    }
    if peminjaman.DurasiHari < 1 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Durasi hari harus minimal 1"})
    }

    // Ambil ID user dari token JWT
    userId := c.Get("userId").(int)
    peminjaman.IDUser = userId

    // Set tanggal pinjam menjadi sekarang
    peminjaman.TanggalPinjam = time.Now()

    // Set status pengembalian menjadi "Belum melakukan pengembalian"
    peminjaman.StatusPengembalian = "Belum melakukan pengembalian"

    // Ambil data buku
    var buku models.Buku
    if err := config.DB.First(&buku, peminjaman.IDBuku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Buku tidak ditemukan"})
    }

    // Pastikan jumlah buku yang tersedia cukup
    if buku.Jumlah <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Buku tidak tersedia"})
    }

    // Kurangi jumlah buku yang tersedia
    buku.Jumlah--

    // Simpan peminjaman ke database
    if err := config.DB.Create(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan data peminjaman"})
    }

    // Perbarui jumlah buku
    if err := config.DB.Save(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui jumlah buku"})
    }

    // Ambil data pengguna untuk email
    var user models.User
    if err := config.DB.First(&user, peminjaman.IDUser).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data user"})
    }

    // Kirim email konfirmasi
    userEmail := user.Email
    subject := "Konfirmasi Peminjaman Buku"
    body := "Terima kasih telah meminjam buku. Harap kembalikan sesuai tanggal yang ditentukan."
    utils.SendEmail(userEmail, subject, body)

    return c.JSON(http.StatusCreated, peminjaman)
}



func GetPeminjamanbyUser(c echo.Context) error {
    userId := c.Get("userId").(int)

    var peminjaman []models.Peminjaman
    if err := config.DB.Preload("Buku").Preload("Buku.Penulis").Preload("Buku.Genre").
        Where("id_user = ?", userId).Find(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }


    return c.JSON(http.StatusOK, peminjaman)
}

func GetAllPeminjaman(c echo.Context) error {
    var peminjaman []models.Peminjaman

    if err := config.DB.Preload("Buku").Preload("Buku.Penulis").Preload("Buku.Genre").Find(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data peminjaman: " + err.Error()})
    }



    return c.JSON(http.StatusOK, peminjaman)
}



// Update Peminjaman
func UpdatePeminjaman(c echo.Context) error {
	id := c.Param("id")
	var peminjaman models.Peminjaman

	if err := config.DB.First(&peminjaman, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Peminjaman not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if err := c.Bind(&peminjaman); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := config.DB.Save(&peminjaman).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, peminjaman)
}

// Delete Peminjaman
func DeletePeminjaman(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Peminjaman{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Peminjaman deleted"})
}

// ReturnBook mengelola pengembalian buku oleh user
func ReturnBook(c echo.Context) error {
    idPeminjaman, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID peminjaman tidak valid"})
    }

    // Ambil ID user dari JWT
    userId := c.Get("userId").(int)

    var peminjaman models.Peminjaman
    if err := config.DB.First(&peminjaman, idPeminjaman).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Data peminjaman tidak ditemukan"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari data peminjaman"})
    }

    // Validasi bahwa hanya user yang melakukan peminjaman dapat mengembalikan
    if peminjaman.IDUser != userId {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda tidak berhak mengembalikan peminjaman ini"})
    }

    // Cek apakah buku sudah dalam proses pengembalian atau dikembalikan
    if peminjaman.StatusPengembalian != "Belum melakukan pengembalian" && peminjaman.StatusKembali {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Buku sudah dikembalikan atau sedang menunggu persetujuan"})
    }

    // Set status pengembalian menjadi Pending
    peminjaman.StatusPengembalian = "Pending"
    peminjaman.StatusKembali = true  // Set StatusKembali menjadi true saat pengembalian dimulai
    peminjaman.TanggalKembali = time.Now()

    if err := config.DB.Save(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate data peminjaman"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Pengembalian dalam proses persetujuan admin"})
}


// VerifyReturn mengelola verifikasi pengembalian oleh admin
func VerifyReturn(c echo.Context) error {
    idPeminjaman, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID peminjaman tidak valid"})
    }

    var peminjaman models.Peminjaman
    if err := config.DB.First(&peminjaman, idPeminjaman).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Data peminjaman tidak ditemukan"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari data peminjaman"})
    }

    // Validasi status pengembalian
    if peminjaman.StatusPengembalian != "Pending" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Pengembalian sudah diverifikasi"})
    }

    // Perbarui status pengembalian dan jumlah buku
    peminjaman.StatusPengembalian = "Approved"
    peminjaman.StatusKembali = true

    var buku models.Buku
    if err := config.DB.First(&buku, peminjaman.IDBuku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data buku"})
    }

    // Tambah jumlah buku yang tersedia
    buku.Jumlah++

    // Simpan perubahan
    if err := config.DB.Save(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate data peminjaman"})
    }
    if err := config.DB.Save(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui jumlah buku"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Pengembalian berhasil diverifikasi"})
}

// RejectReturn mengelola penolakan pengembalian oleh admin
func RejectReturn(c echo.Context) error {
    idPeminjaman, _ := strconv.Atoi(c.Param("id"))
    var requestData struct {
        Alasan string `json:"alasan"`
    }
    if err := c.Bind(&requestData); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Alasan penolakan harus diisi"})
    }

    // Ambil data peminjaman
    var peminjaman models.Peminjaman
    if err := config.DB.First(&peminjaman, idPeminjaman).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Data peminjaman tidak ditemukan"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari data peminjaman"})
    }

    // Perbarui status pengembalian menjadi ditolak
    peminjaman.StatusPengembalian = "Rejected"
    peminjaman.StatusKembali = false // status kembali tetap dipinjam

    // Simpan perubahan status peminjaman
    if err := config.DB.Save(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui status peminjaman"})
    }

    // Kirim email alasan penolakan ke pengguna
    var user models.User
    if err := config.DB.First(&user, peminjaman.IDUser).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data pengguna"})
    }
    subject := "Alasan Penolakan Pengembalian Buku"
    body := "Pengembalian buku Anda ditolak. Alasan penolakan: " + requestData.Alasan
    if err := utils.SendEmail(user.Email, subject, body); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengirim email penolakan"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Pengembalian berhasil ditolak"})
}

// GetPengembalianPending mengambil daftar pengembalian dengan status Pending
func GetPengembalianPending(c echo.Context) error {
    var peminjaman []models.Peminjaman
    if err := config.DB.Preload("Buku").Where("status_pengembalian = ?", "Pending").Find(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data pengembalian"})
    }
    return c.JSON(http.StatusOK, peminjaman)
}
