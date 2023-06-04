package daos

import "time"

type FotoProduk struct {
	ID        uint
	ProdukID  uint
	URL       string `gorm:"type:varchar(255)"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
