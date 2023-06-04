package dto

type UploadProdukRequest struct {
	NamaProduk    string `validate:"reuqired"`
	CategoryID    uint   `validate:"reuqired"`
	TokoID        uint   `validate:"reuqired"`
	HargaReseller uint   `validate:"reuqired"`
	HargaKonsumen uint   `validate:"reuqired"`
	Stok          uint   `validate:"reuqired"`
	Deskripsi     string `validate:"reuqired"`
	Photos        []Photos
}

type Photos struct {
	URL string
}
type FotoProdukGetProduk struct {
	ID       uint   `json:"id"`
	ProdukID uint   `json:"produk_id"`
	URL      string `json:"url"`
}

type GetProduk struct {
	ID            uint                  `json:"id"`
	NamaProduk    string                `json:"nama_produk"`
	Slug          string                `json:"slug"`
	HargaReseller uint                  `json:"harga_reseller"`
	HargaKonsumen uint                  `json:"harga_konsumen"`
	Stok          uint                  `json:"stok"`
	Deskripsi     string                `json:"deskripsi"`
	Toko          GetTokoByIDResponse   `json:"toko"`
	Category      CategoryWithID        `json:"category"`
	FotoProduk    []FotoProdukGetProduk `json:"foto_produk"`
}

type UpdateProdukRequest struct {
	ID            uint
	NamaProduk    string `validate:"reuqired"`
	CategoryID    uint   `validate:"reuqired"`
	TokoID        uint   `validate:"reuqired"`
	HargaReseller uint   `validate:"reuqired"`
	HargaKonsumen uint   `validate:"reuqired"`
	Stok          uint   `validate:"reuqired"`
	Deskripsi     string `validate:"reuqired"`
	Photos        []Photos
}

type FilterProduk struct {
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
	NamaProduk string `query:"nama_produk"`
	CategoryID uint   `query:"category_id"`
	TokoID     uint   `query:"toko_id"`
	MaxHarga   uint   `query:"max_harga"`
	MinHarga   uint   `query:"min_harga"`
}
