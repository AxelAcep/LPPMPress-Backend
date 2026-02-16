package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Jenis     string `gorm:"type:varchar(20);not null"` // 'admin' atau 'pelanggan'
	CreatedAt time.Time
}

type Admin struct {
	UserID uint   `gorm:"primaryKey"`
	Nama   string `gorm:"size:50;not null"`
	User   User   `gorm:"foreignKey:UserID"`
}

type Pelanggan struct {
	UserID      uint   `gorm:"primaryKey"`
	NamaLengkap string `gorm:"size:50;not null"`
	NoTelp      string `gorm:"size:15"`
	Alamat      string `gorm:"type:text"`
	User        User   `gorm:"foreignKey:UserID"`
}

type Kategori struct {
	ID           uint   `gorm:"primaryKey"`
	NamaKategori string `gorm:"size:50;not null"`
	Buku         []Buku `gorm:"foreignKey:KategoriID"`
}

type Buku struct {
	ID            uint `gorm:"primaryKey"`
	KategoriID    uint
	NamaBuku      string  `gorm:"size:255;not null"`
	Penulis       string  `gorm:"size:100"`
	Harga         float64 `gorm:"type:decimal(12,2)"`
	Stok          int
	StatusFisik   bool   `gorm:"default:false"`
	StatusDigital bool   `gorm:"default:true"`
	Deskripsi     string `gorm:"type:text"`
}

type Pesanan struct {
	ID              uint `gorm:"primaryKey"`
	PelangganID     uint
	TanggalPesan    time.Time  `gorm:"autoCreateTime"`
	TanggalSampai   *time.Time // Pointer agar bisa nil/null
	TotalHarga      float64    `gorm:"type:decimal(12,2)"`
	Status          string     `gorm:"default:'pembayaran'"`
	JenisPengiriman string
	Items           []PesananItem `gorm:"foreignKey:PesananID"`
}

type PesananItem struct {
	ID        uint `gorm:"primaryKey"`
	PesananID uint
	BukuID    uint
	Jumlah    int
	Subtotal  float64 `gorm:"type:decimal(12,2)"`
}
