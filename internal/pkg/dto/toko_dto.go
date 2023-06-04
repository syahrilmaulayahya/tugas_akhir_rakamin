package dto

type GetTokoResponse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_Foto"`
	UserID   uint   `json:"user_id,omitempty"`
}

type GetTokoByUserIDResponse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   uint   `json:"user_id"`
}

type GetTokoByIDResponse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}
type TokoFilter struct {
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Nama  string `query:"nama"`
}

type GetAllTokoResponse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}

type UpdateTokoRequest struct {
	NamaToko string `json:"nama_toko"`
	Photo    string `json:"photo"`
}

type TokoTRX struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}
