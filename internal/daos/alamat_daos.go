package daos

import "time"

type Alamat struct {
	ID           uint
	UserID       uint   `gorm:"not null"`
	JudulAlamat  string `gorm:"type:varchar(255)"`
	NamaPenerima string `gorm:"type:varchar(255)"`
	NoTelp       string `gorm:"type:varchar(255)"`
	DetailAlamat string `gorm:"type:varchar(255)"`
	UpdatedAt    time.Time
	CreatedAt    time.Time
	TRX          []TRX
}

type FilterAlamat struct {
	JudulAlamat string
}
