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


func CreatePeminjaman(c echo.Context) error {
    var peminjaman models.Peminjaman

    // Bind input ke model Peminjaman
    if err := c.Bind(&peminjaman); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format data tidak valid"})
    }

    // Validasi input yang wajib
    if peminjaman.IDBuku == 0 || peminjaman.JumlahPinjam <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID buku dan jumlah pinjam tidak boleh kosong atau nol"})
    }

    // Ambil ID user dari JWT
    userId, ok := c.Get("userId").(int)
    if !ok || userId == 0 {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Pengguna tidak valid"})
    }

    // Jika IDUser tidak diberikan dalam request, gunakan userId dari context
    if peminjaman.IDUser == 0 {
        peminjaman.IDUser = userId
    }

    peminjaman.TanggalPinjam = time.Now()
    peminjaman.StatusPengembalian = "Belum melakukan pengembalian"

    // Validasi ketersediaan buku
    var buku models.Buku
    if err := config.DB.First(&buku, peminjaman.IDBuku).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Buku tidak ditemukan"})
    }

    if buku.Jumlah < peminjaman.JumlahPinjam {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Jumlah buku yang tersedia tidak mencukupi"})
    }

    // Kurangi jumlah buku yang tersedia
    buku.Jumlah -= peminjaman.JumlahPinjam
    if err := config.DB.Save(&buku).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui jumlah buku"})
    }

    // Simpan data peminjaman
    if err := config.DB.Create(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan data peminjaman"})
    }

    // Kirim email pemberitahuan
    var user models.User
    if err := config.DB.First(&user, peminjaman.IDUser).Error; err == nil {
        subject := "Konfirmasi Peminjaman Buku"
        body := "Anda telah meminjam buku: " + buku.Judul + ". Harap dikembalikan sebelum tanggal " + peminjaman.TanggalPinjam.AddDate(0, 0, peminjaman.DurasiHari).Format("02-01-2006") + "."
        utils.SendEmail(user.Email, subject, body)
    }

    return c.JSON(http.StatusCreated, peminjaman)
}

func GetPeminjamanbyUser(c echo.Context) error {
    userId := c.Get("userId").(int)

    var peminjaman []models.Peminjaman
    if err := config.DB.Preload("Buku").Preload("Buku.Penulis").Preload("Buku.Genre").
        Where("id_user = ?", userId).Find(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }
// Menampilkan data pengguna dalam respon
for i := range peminjaman {
	var user models.User
	if err := config.DB.First(&user, peminjaman[i].IDUser).Error; err == nil {
		peminjaman[i].User = user
	}
}

    return c.JSON(http.StatusOK, peminjaman)
}

func GetAllPeminjaman(c echo.Context) error {
    var peminjaman []models.Peminjaman

    if err := config.DB.Preload("Buku").Preload("Buku.Penulis").Preload("Buku.Genre").Find(&peminjaman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data peminjaman: " + err.Error()})
    }

// Menampilkan data pengguna dalam respon
for i := range peminjaman {
	var user models.User
	if err := config.DB.First(&user, peminjaman[i].IDUser).Error; err == nil {
		peminjaman[i].User = user
	}
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

	userId := c.Get("userId").(int)

	var peminjaman models.Peminjaman
	if err := config.DB.First(&peminjaman, idPeminjaman).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Data peminjaman tidak ditemukan"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari data peminjaman"})
	}

	if peminjaman.IDUser != userId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda tidak berhak mengembalikan peminjaman ini"})
	}

	// Jika status pengembalian ditolak, kita masih memperbolehkan pengembalian
	if peminjaman.StatusPengembalian == "Rejected" && peminjaman.StatusKembali {
		peminjaman.StatusPengembalian = "Pending" // Mengubah status pengembalian ke Pending lagi
		peminjaman.TanggalKembali = time.Now()

		if err := config.DB.Save(&peminjaman).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate data peminjaman"})
		}

		// Kirim email pemberitahuan pengembalian
		var user models.User
		if err := config.DB.First(&user, userId).Error; err == nil {
			subject := "Pengembalian Buku"
			body := "Pengembalian buku Anda sedang menunggu persetujuan admin."
			utils.SendEmail(user.Email, subject, body)
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Pengembalian dalam proses persetujuan admin"})
	}

	// Jika pengembalian sudah diproses atau sudah dikembalikan
	if peminjaman.StatusPengembalian != "Belum melakukan pengembalian" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Buku sudah dikembalikan atau sedang menunggu persetujuan"})
	}

	// Jika status pengembalian belum dilakukan
	peminjaman.StatusPengembalian = "Pending"
	peminjaman.StatusKembali = true
	peminjaman.TanggalKembali = time.Now()

	if err := config.DB.Save(&peminjaman).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate data peminjaman"})
	}

	// Kirim email pemberitahuan pengembalian
	var user models.User
	if err := config.DB.First(&user, userId).Error; err == nil {
		subject := "Pengembalian Buku"
		body := "Pengembalian buku Anda sedang menunggu persetujuan admin."
		utils.SendEmail(user.Email, subject, body)
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

	if peminjaman.StatusPengembalian != "Pending" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Pengembalian sudah diverifikasi"})
	}

	peminjaman.StatusPengembalian = "Approved"
	peminjaman.StatusKembali = true

	var buku models.Buku
	if err := config.DB.First(&buku, peminjaman.IDBuku).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data buku"})
	}

	buku.Jumlah += peminjaman.JumlahPinjam

	if err := config.DB.Save(&peminjaman).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate data peminjaman"})
	}

	if err := config.DB.Save(&buku).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui jumlah buku"})
	}

	// Kirim email pemberitahuan verifikasi pengembalian
	var user models.User
	if err := config.DB.First(&user, peminjaman.IDUser).Error; err == nil {
		subject := "Pengembalian Buku Diverifikasi"
		body := "Pengembalian buku Anda telah diverifikasi. Terima kasih."
		utils.SendEmail(user.Email, subject, body)
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

    // Perbarui status pengembalian menjadi ditolak, namun status peminjaman tetap dipinjam
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

// SendReminderEmail mengirimkan pengingat pengembalian buku
func SendReminderEmail() {
	var peminjaman []models.Peminjaman
	now := time.Now()

	if err := config.DB.Preload("Buku").Where("status_pengembalian = ? AND tanggal_pinjam + interval '1 day' * durasi_hari = ?", "Belum melakukan pengembalian", now.Format("2006-01-02")).Find(&peminjaman).Error; err != nil {
		return
	}

	for _, p := range peminjaman {
		var user models.User
		if err := config.DB.First(&user, p.IDUser).Error; err == nil {
			subject := "Pengingat Pengembalian Buku"
			body := "Harap kembalikan buku: " + p.Buku.Judul + " sebelum tanggal " + p.TanggalPinjam.AddDate(0, 0, p.DurasiHari).Format("02-01-2006") + "."
			utils.SendEmail(user.Email, subject, body)
		}
	}
}