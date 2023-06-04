package daos

import "time"

type LogProduk struct {
	ID            uint
	ProdukID      uint   `gorm:"not null"`
	NamaProduk    string `gorm:"type:varchar(255)"`
	Slug          string `gorm:"type:varchar(255)"`
	HargaReseller uint
	HargaKonsumen uint
	Deskripsi     string `gorm:"type:text"`
	TokoID        uint   `gorm:"not null"`
	Toko          Toko
	CategoryID    uint `gorm:"not null"`
	Category      Category
	LogFotoProduk []LogFotoProduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DetailTRX     DetailTRX
}

type LogFotoProduk struct {
	ID          uint
	LogProdukID uint
	URL         string `gorm:"type:varchar(255)"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
