package repository

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
	"net/http"
)

type TokoRepository interface {
	GetTokoByID(ctx context.Context, ID uint) (response daos.Toko, errHelper *helper.ErrorStruct)
	GetTokoByUserID(ctx context.Context, userID uint) (response daos.Toko, errHelper *helper.ErrorStruct)
	GetAllToko(ctx context.Context, params daos.FilterToko) (response []daos.Toko, errHelper *helper.ErrorStruct)
	UpdateToko(ctx context.Context, data daos.Toko) (errHelper *helper.ErrorStruct)
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}

func (tr *TokoRepositoryImpl) GetTokoByID(ctx context.Context, ID uint) (response daos.Toko, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db

	// get toko data by id
	errDb := db.First(&response, "id = ?", ID).Error
	if errDb != nil {
		// check if toko with specified id available
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("Toko tidak ditemukan"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}

		// response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return response, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (tr *TokoRepositoryImpl) GetTokoByUserID(ctx context.Context, userID uint) (response daos.Toko, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db

	// get toko data by user_id
	errDb := db.First(&response, "user_id = ?", userID).Error
	if errDb != nil {
		// check if toko with specified user_id available
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("Toko tidak ditemukan"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}

		// response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return response, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (tr *TokoRepositoryImpl) GetAllToko(ctx context.Context, params daos.FilterToko) (response []daos.Toko, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db

	// get all toko from database and error information
	if errDb := db.Limit(params.Limit).Offset(params.Offset).Find(&response).Error; errDb != nil {
		if len(response) <= 0 {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("no product found"),
				Code: http.StatusNotFound,
			}
			return []daos.Toko{}, errHelper
		}
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return response, errHelper
	}
	if params.Nama != "" {
		if errDb := db.Limit(params.Limit).Offset(params.Offset).Where("nama_toko LIKE ?", "%"+params.Nama+"%").Find(&response).Error; errDb != nil {
			if len(response) <= 0 {
				errHelper = &helper.ErrorStruct{
					Err:  errors.New("no product found"),
					Code: http.StatusNotFound,
				}
				return []daos.Toko{}, errHelper
			}
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusInternalServerError,
			}
			return response, errHelper
		}
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (tr *TokoRepositoryImpl) UpdateToko(ctx context.Context, data daos.Toko) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db
	var responseDb daos.Toko
	// update toko data
	if errDb := db.Where("user_id = ?", data.UserID).First(&responseDb).Updates(data).Error; errDb != nil {
		// check if error is record not found
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusNotFound,
			}
			return errHelper
		}
		// response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return errHelper
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return errHelper
}
