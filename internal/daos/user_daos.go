package daos

import "time"

type User struct {
	ID           uint
	Nama         string `gorm:"type:varchar(255)"`
	KataSandi    string `gorm:"type:varchar(255)"`
	Notelp       string `gorm:"type:varchar(255);unique"`
	TanggalLahir time.Time
	JenisKelamin string `gorm:"type:varchar(255)"`
	Tentang      string `gorm:"type:text"`
	Pekerjaan    string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255);unique"`
	IDProvinsi   string `gorm:"type:varchar(255)"`
	IDKota       string `gorm:"type:varchar(255)"`
	IsAdmin      bool
	Toko         Toko
	UpdatedAt    time.Time
	CreatedAt    time.Time
	Alamat       []Alamat
	TRX          []TRX
}
