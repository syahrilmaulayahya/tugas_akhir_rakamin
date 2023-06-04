package daos

import "time"

type Toko struct {
	ID        uint
	UserID    uint   `gorm:"not null;unique"`
	NamaToko  string `gorm:"type:varchar(255);not null"`
	UrlFoto   string `gorm:"type:varchar(255)"`
	Produk    []Produk
	UpdatedAt time.Time
	CreatedAt time.Time
	LogProduk []LogProduk
	DetailTRX []DetailTRX
}

type FilterToko struct {
	Limit  int
	Offset int
	Nama   string
}
