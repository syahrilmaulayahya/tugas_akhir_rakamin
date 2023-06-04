package repository

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
	"net/http"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, data daos.Category) (id uint, errHelper *helper.ErrorStruct)
	GetAllCategory(ctx context.Context) (response []daos.Category, errHelper *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, ID uint) (response daos.Category, errHelper *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, data daos.Category) (errHelper *helper.ErrorStruct)
	DeleteCategoryByID(cxx context.Context, data daos.Category) (errHelper *helper.ErrorStruct)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (cr *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data daos.Category) (id uint, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := cr.db

	// create new categories record in database
	errDb := db.Create(&data).Error

	// error checking
	if errDb != nil {
		// check if new category name are duplicate
		var mysqlErr *mysql.MySQLError
		if errors.As(errDb, &mysqlErr) && mysqlErr.Number == 1062 {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusBadRequest,
			}
			return id, errHelper
		}

		// response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return id, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	id = data.ID
	return id, errHelper
}

func (cr *CategoryRepositoryImpl) GetAllCategory(ctx context.Context) (response []daos.Category, errHelper *helper.ErrorStruct) {
	db := cr.db

	// get all categories from database
	errDb := db.Find(&response).Error

	// error checking
	if errDb != nil {
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

func (cr *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, ID uint) (response daos.Category, errHelper *helper.ErrorStruct) {
	db := cr.db

	// get category record from database with id as condition
	errDb := db.First(&response, "id = ?", ID).Error
	// error checking
	if errDb != nil {
		// check if record not found
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("No Data Category"),
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

func (cr *CategoryRepositoryImpl) UpdateCategoryByID(ctx context.Context, data daos.Category) (errHelper *helper.ErrorStruct) {
	db := cr.db

	// update category name in database
	errDb := db.First(&data).Update("nama_category", data.NamaCategory).Error

	if errDb != nil {
		// check if id available
		if errDb == gorm.ErrRecordNotFound {
			err := errDb
			errHelper = &helper.ErrorStruct{
				Err:  err,
				Code: http.StatusNotFound,
			}
			return errHelper
		}

		// check if name duplicate
		var mysqlErr *mysql.MySQLError
		if errors.As(errDb, &mysqlErr) && mysqlErr.Number == 1062 {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusBadRequest,
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

func (cr *CategoryRepositoryImpl) DeleteCategoryByID(ctx context.Context, data daos.Category) (errHelper *helper.ErrorStruct) {
	db := cr.db

	// delete record in categories with id as identifier
	errDb := db.First(&data).Delete(&data).Error

	// check another error
	if errDb != nil {
		// check if id available
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusBadRequest,
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
