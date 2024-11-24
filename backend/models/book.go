package models

import (
    "time"
)

type Buku struct {
    ID             int       `gorm:"column:id_buku;primaryKey" json:"id_buku"`
    Judul          string    `gorm:"type:varchar(100);not null" json:"judul"`
    IdPenulis      int       `gorm:"column:id_penulis" json:"id_penulis"`
    IdGenre        int       `gorm:"column:id_genre" json:"id_genre"`
    Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
    Jumlah         int       `gorm:"not null" json:"jumlah"`
    Gambar         string    `gorm:"type:varchar(255)" json:"gambar"`
    Status         bool      `gorm:"default:true" json:"status"`
    DibuatPada     time.Time `gorm:"column:dibuat_pada;default:CURRENT_TIMESTAMP" json:"dibuat_pada"`
    DiperbaruiPada time.Time `gorm:"column:diperbarui_pada;default:CURRENT_TIMESTAMP" json:"diperbarui_pada"`

    Penulis        Penulis  `gorm:"foreignKey:IdPenulis;references:IDPenulis" json:"penulis"`
    Genre          Genre    `gorm:"foreignKey:IdGenre;references:IDGenre" json:"genre"`
}

// TableName memberi tahu GORM untuk menggunakan nama tabel 'books'
func (Buku) TableName() string {
    return "books"
}
