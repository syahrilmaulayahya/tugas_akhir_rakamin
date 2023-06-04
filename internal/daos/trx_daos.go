package daos

import "time"

type TRX struct {
	ID          uint
	UserID      uint
	AlamatID    uint
	HargaTotal  uint
	KodeInvoice string `gorm:"type:varchar(255)"`
	MethodBayar string `gorm:"type:varchar(255)"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
	DetailTRX   []DetailTRX
}

type TRXResponse struct {
	ID          uint
	UserID      uint
	AlamatID    uint
	Alamat      Alamat
	HargaTotal  uint
	KodeInvoice string `gorm:"type:varchar(255)"`
	MethodBayar string `gorm:"type:varchar(255)"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
	DetailTRX   []DetailTRXResponse
}

type FilterTRX struct {
	Limit  int
	Offset int
}
