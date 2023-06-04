package dto

type TRX struct {
	UserID      uint        `json:"-" validate:"required"`
	KodeInvoice string      `json:"-"`
	MethodBayar string      `json:"method_bayar" validate:"required"`
	AlamatID    uint        `json:"alamat_kirim" validate:"required"`
	DetailTRX   []DetailTRX `json:"detail_trx" validate:"required"`
}

type DetailTRX struct {
	ProductID uint `json:"product_id" validate:"required"`
	Kuantitas uint `json:"kuantitas" validate:"required"`
}

type TRXGetResponse struct {
	ID          uint                   `json:"id"`
	HargaTotal  uint                   `json:"harga_total"`
	KodeInvoice string                 `json:"kode_invoice"`
	MethodBayar string                 `json:"method_bayar"`
	Alamat      AlamatTRX              `json:"alamat_kirim"`
	DetailTRX   []DetailTRXGetResponse `json:"detail_trx"`
}

type DetailTRXGetResponse struct {
	Product    LogProdukGetResponse `json:"product"`
	Toko       GetTokoByIDResponse  `json:"toko"`
	Kuantitas  uint                 `json:"kuantitas"`
	HargaTotal uint                 `json:"harga_total"`
}

type LogProdukGetResponse struct {
	ID            uint                    `json:"id"`
	ProdukID      uint                    `json:"produk_id"`
	NamaProduk    string                  `json:"nama_produk"`
	Slug          string                  `json:"slug"`
	HargaReseller uint                    `json:"harga_reseller"`
	HargaKonsumen uint                    `json:"harga_konsumen"`
	Deskripsi     string                  `json:"deskripsi"`
	Toko          TokoTRX                 `json:"toko"`
	Category      CategoryWithID          `json:"category"`
	Photos        []LogFotoProdukResponse `json:"photos"`
}

type LogFotoProdukResponse struct {
	ID          uint   `json:"id"`
	LogProdukID uint   `json:"log_produk_id"`
	URL         string `json:"url"`
}

type FilterTRX struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}
