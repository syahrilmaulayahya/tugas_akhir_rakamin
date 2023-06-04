package daos

import "time"

type Category struct {
	ID           uint
	NamaCategory string `gorm:"type varchar(255);not null;unique"`
	Produk       []Produk
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LogProduk    []LogProduk
}
