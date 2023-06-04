package daos

import "time"

type DetailTRX struct {
	ID          uint
	TRXID       uint
	LogProdukID uint
	TokoID      uint
	Kuantitas   uint
	HargaTotal  uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type DetailTRXResponse struct {
	ID          uint
	TRXID       uint
	LogProdukID uint
	TokoID      uint
	Toko        Toko
	Kuantitas   uint
	HargaTotal  uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
	LogProduk   LogProduk
}
