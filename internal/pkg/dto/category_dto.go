package dto

type CategoryWithID struct {
	ID           uint   `json:"id"`
	NamaCategory string `json:"nama_category"`
}

type CategoryIDOnly struct {
	ID uint `json:"id"`
}

type CreateAndUpdateCategoryRequest struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}
