package repository

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
	"net/http"
)

type AlamatRepository interface {
	CreateAlamat(ctx context.Context, data daos.Alamat) (ID uint, errHelper *helper.ErrorStruct)
	GetMyAlamat(ctx context.Context, userID uint, params daos.FilterAlamat) (response []daos.Alamat, errHelper *helper.ErrorStruct)
	GetAlamatByID(ctx context.Context, ID uint) (response daos.Alamat, errHelper *helper.ErrorStruct)
	UpdateAlamatByID(ctx context.Context, data daos.Alamat) (errHelper *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, userID, ID uint) (errHelper *helper.ErrorStruct)
}

type AlamatRepositoryImpl struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &AlamatRepositoryImpl{db: db}
}

func (ar *AlamatRepositoryImpl) CreateAlamat(ctx context.Context, data daos.Alamat) (ID uint, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ar.db

	// create alamat record to database and get error information
	if errDb := db.Create(&data).Error; errDb != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return ID, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return data.ID, errHelper
}

func (ar *AlamatRepositoryImpl) GetMyAlamat(ctx context.Context, userID uint, params daos.FilterAlamat) (response []daos.Alamat, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ar.db

	// get list alamat with specified user_id from database and get error information
	if errDb := db.Where("user_id = ?", userID).Find(&response).Error; errDb != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return response, errHelper
	}
	if params.JudulAlamat != "" {
		if errDb := db.Where("user_id = ?", userID).Where("judul_alamat LIKE ?", "%"+params.JudulAlamat+"%").Find(&response).Error; errDb != nil {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusInternalServerError,
			}
			return response, errHelper
		}
	}
	if len(response) <= 0 {
		errHelper = &helper.ErrorStruct{
			Err:  errors.New("alamat not found"),
			Code: http.StatusNotFound,
		}
		return []daos.Alamat{}, errHelper
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (ar *AlamatRepositoryImpl) GetAlamatByID(ctx context.Context, ID uint) (response daos.Alamat, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ar.db

	// get alamat record with specified userID and ID and get error information
	errDb := db.Where("id = ?", ID).First(&response).Error
	// error checking
	if errDb != nil {
		// check error if record not found
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
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

func (ar *AlamatRepositoryImpl) UpdateAlamatByID(ctx context.Context, data daos.Alamat) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ar.db

	// variable to store response from database
	alamat := daos.Alamat{}
	// update alamat with specified user_id and id
	if errDb := db.Model(&data).Where("user_id = ? AND id = ?", data.UserID, data.ID).First(&alamat).Updates(data).Error; errDb != nil {
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

func (ar *AlamatRepositoryImpl) DeleteAlamatByID(ctx context.Context, userID, ID uint) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ar.db

	// variable to store response from database
	alamat := daos.Alamat{}

	// delete alamat record with specified userID and ID and get error information
	errDb := db.Where("user_id = ? AND id = ?", userID, ID).First(&alamat).Delete(&alamat).Error
	if errDb != nil {
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
