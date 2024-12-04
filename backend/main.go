package main

import (
	"backend/config"
	"backend/middleware"
	"backend/models"
	"backend/routes"
	"backend/utils"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()
	if config.DB == nil {
		panic("Database connection failed. Please check your configuration.")
	}

	// Inisialisasi Echo
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/uploads", "uploads")

	// Daftarkan routes
	routes.AuthRoutes(e)
	routes.UserRoutes(e)
	routes.GenreRoutes(e)
	routes.PenulisRoutes(e)
	routes.BookRoutes(e)
	routes.PeminjamanRoutes(e)

	// Jalankan scheduler untuk notifikasi email
	go scheduleEmailNotifications()

	// Jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}

// Fungsi untuk mengatur scheduler
func scheduleEmailNotifications() {
	s := gocron.NewScheduler(time.UTC)

	// Jadwalkan fungsi pengingat pengembalian buku
	s.Every(24).Hours().Do(SendReminderEmail)

	// Menjalankan scheduler secara asinkron
	s.StartAsync()
}


// Fungsi untuk pengingat pengembalian buku (berdasarkan logika tambahan)
func SendReminderEmail() {
	var peminjaman []models.Peminjaman
	now := time.Now()

	// Ambil data peminjaman yang sesuai dengan kriteria
	if err := config.DB.Preload("Buku").Where("status_pengembalian = ? AND tanggal_pinjam + interval '1 day' * durasi_hari = ?", "Belum melakukan pengembalian", now.Format("2006-01-02")).Find(&peminjaman).Error; err != nil {
		fmt.Println("Error fetching peminjaman:", err)
		return
	}

	// Kirim email untuk setiap peminjaman yang ditemukan
	for _, p := range peminjaman {
		var user models.User
		if err := config.DB.First(&user, p.IDUser).Error; err == nil {
			subject := "Pengingat Pengembalian Buku"
			body := "Harap kembalikan buku: " + p.Buku.Judul + " sebelum tanggal " + p.TanggalPinjam.AddDate(0, 0, p.DurasiHari).Format("02-01-2006") + "."
			if err := utils.SendEmail(user.Email, subject, body); err != nil {
				fmt.Println("Gagal mengirim email ke:", user.Email, "Error:", err)
			}
		}
	}
}
