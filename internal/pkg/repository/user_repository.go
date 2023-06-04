package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"net/http"
	"strings"

	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx context.Context, data daos.User) (errHelper *helper.ErrorStruct)
	Login(ctx context.Context, noTelp string) (response daos.User, errHelper *helper.ErrorStruct)
	GetMyProfile(ctx context.Context, userID uint) (response daos.User, errHelper *helper.ErrorStruct)
	UpdateProfile(ctx context.Context, data daos.User) (errHelper *helper.ErrorStruct)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur *UserRepositoryImpl) Register(ctx context.Context, data daos.User) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ur.db

	// create user and toko record at the same time with gorm transaction
	errDb := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		toko := daos.Toko{
			ID:       data.ID,
			UserID:   data.ID,
			NamaToko: fmt.Sprintf("Toko %s", data.Nama),
			UrlFoto:  fmt.Sprintf("example.com/%s/toko", strings.ToLower(data.Nama)),
		}
		if err := tx.Create(&toko).Error; err != nil {
			return err
		}

		return nil
	})
	// error checking
	if errDb != nil {
		// check if error is duplicate notelp or email
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

func (ur *UserRepositoryImpl) Login(ctx context.Context, noTelp string) (response daos.User, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ur.db

	// get user data by notelp
	if errDb := db.Model(&daos.User{}).First(&response, "notelp = ?", noTelp).Error; errDb != nil {
		// check if error is user not found
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("No Telp atau kata sandi salah"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}
		//response another error
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

func (ur *UserRepositoryImpl) GetMyProfile(ctx context.Context, userID uint) (response daos.User, errHelper *helper.ErrorStruct) {
	db := ur.db

	// get user data from database with specified id
	if errDb := db.Model(&daos.User{}).First(&response, "id = ?", userID).Error; errDb != nil {
		// check if error is record not found
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

func (ur *UserRepositoryImpl) UpdateProfile(ctx context.Context, data daos.User) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := ur.db
	// update record with primary key id in data
	if errDb := db.Model(&data).Updates(data).Error; errDb != nil {
		// check if error is duplicate notelp or email
		var mysqlErr *mysql.MySQLError
		if errors.As(errDb, &mysqlErr) && mysqlErr.Number == 1062 {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusBadRequest,
			}
			return errHelper
		}
		// response other error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return errHelper
	}
	//success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return errHelper
}
