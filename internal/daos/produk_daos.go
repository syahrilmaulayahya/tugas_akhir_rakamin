package daos

import "time"

type Produk struct {
	ID            uint
	NamaProduk    string `gorm:"type:varchar(255)"`
	Slug          string `gorm:"type:varchar(255)"`
	HargaReseller uint
	HargaKonsumen uint
	Stok          uint
	Deskripsi     string `gorm:"type:text"`
	TokoID        uint   `gorm:"not null"`
	Toko          Toko
	CategoryID    uint `gorm:"not null"`
	Category      Category
	FotoProduk    []FotoProduk
	UpdatedAt     time.Time
	CreatedAt     time.Time
	LogProduk     []LogProduk
}

type FilterProduk struct {
	NamaProduk string
	Limit      int
	Offset     int
	CategoryID uint
	TokoID     uint
	MaxHarga   uint
	MinHarga   uint
}
