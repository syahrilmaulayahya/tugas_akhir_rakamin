package repository

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"os"
)

type ProdukRepository interface {
	UploadProduk(ctx context.Context, data daos.Produk) (ID uint, errHelper *helper.ErrorStruct)
	GetProdukByID(ctx context.Context, ID uint) (response daos.Produk, errHelper *helper.ErrorStruct)
	UpdateProdukByID(ctx context.Context, data daos.Produk) (errHelper *helper.ErrorStruct)
	DeleteProdukByID(ctx context.Context, tokoID, ID uint) (errHelper *helper.ErrorStruct)
	GetAllProduk(ctx context.Context, params daos.FilterProduk) (response []daos.Produk, errHelper *helper.ErrorStruct)
}

type ProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &ProdukRepositoryImpl{db: db}
}

func (pr *ProdukRepositoryImpl) UploadProduk(ctx context.Context, dataProduk daos.Produk) (ID uint, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := pr.db

	//create produk and foto_produks record in database and get error information
	if errDb := db.Create(&dataProduk).Error; errDb != nil {
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
	return dataProduk.ID, errHelper
}

func (pr *ProdukRepositoryImpl) GetProdukByID(ctx context.Context, ID uint) (response daos.Produk, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := pr.db

	// get produk record from database and error information
	errDb := db.Preload("FotoProduk").Preload("Category").Preload("Toko").First(&response, ID)
	// error handle if record not found
	if errDb.Error != nil {
		if errDb.Error == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("No Data Product"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}
		//response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb.Error,
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

func (pr *ProdukRepositoryImpl) UpdateProdukByID(ctx context.Context, data daos.Produk) (errHelper *helper.ErrorStruct) {
	// get gorm client
	var responseDb daos.Produk

	db := pr.db

	// update with transaction
	errDb := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("toko_id = ? AND id = ?", data.TokoID, data.ID).First(&responseDb).Updates(data).Error; err != nil {
			// return any error will roll back
			return err
		}
		if len(data.FotoProduk) > 0 {
			if err := tx.Create(&data.FotoProduk).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	// error checking
	if errDb != nil {
		// check if error is record not found
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusNotFound,
			}
			return errHelper
		}
		// check another error
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

func (pr *ProdukRepositoryImpl) DeleteProdukByID(ctx context.Context, tokoID, ID uint) (errHelper *helper.ErrorStruct) {
	// get gorm client
	db := pr.db
	var fotoDb daos.FotoProduk
	var produkDb daos.Produk
	var listFoto []daos.FotoProduk
	// update with transaction
	errDb := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		getFoto := tx.Where("produk_id = ?", ID).Find(&listFoto)
		if getFoto.Error == nil {
			for _, v := range listFoto {
				if _, err := os.Stat(v.URL); err == nil {
					if e := os.Remove(v.URL); e != nil {
						helper.Logger("produk_repository", helper.LoggerLevelInfo, "Failed to DELETE foto")
					}
				}
			}
		}
		if err := tx.Where("produk_id = ?", ID).Delete(&fotoDb).Error; err != nil {
			// return any error will roll back
			return err
		}
		if err := tx.Where("toko_id = ? AND id = ?", tokoID, ID).First(&produkDb).Delete(&produkDb).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	// check if error is record not found
	if errDb != nil {
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

func (pr *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, params daos.FilterProduk) (response []daos.Produk, errHelper *helper.ErrorStruct) {

	// get gorm client
	db := pr.db
	filter := daos.Produk{
		TokoID:     params.TokoID,
		CategoryID: params.CategoryID,
	}
	// get produk records from database
	if errDb := db.Where(&filter).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response).Error; errDb != nil {
		// response another error
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return response, errHelper
	}
	if params.NamaProduk != "" {
		errDb := db.Where(&filter).Where("slug LIKE ?", "%"+params.NamaProduk+"%").Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		if params.MaxHarga != 0 {
			errDb = db.Where(&filter).Where("slug LIKE ?", "%"+params.NamaProduk+"%").Where("harga_konsumen <= ?", params.MaxHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		}
		if params.MinHarga != 0 {
			errDb = db.Where(&filter).Where("slug LIKE ?", "%"+params.NamaProduk+"%").Where("harga_konsumen >= ?", params.MinHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		}
		if params.MinHarga != 0 && params.MaxHarga != 0 {
			errDb = db.Where(&filter).Where("slug LIKE ?", "%"+params.NamaProduk+"%").Where("harga_konsumen >= ?", params.MinHarga).Where("harga_konsumen <= ?", params.MaxHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		}
		// check if error is record not found
		if len(response) <= 0 {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("no product found"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}
		// response another error
		if errDb.Error != nil {

			errHelper = &helper.ErrorStruct{
				Err:  errDb.Error,
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
	if params.MaxHarga != 0 {
		errDb := db.Where(&filter).Where("harga_konsumen <= ?", params.MaxHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		if params.MinHarga != 0 {
			errDb = db.Where(&filter).Where("slug LIKE ?", "%"+params.NamaProduk+"%").Where("harga_konsumen >= ?", params.MinHarga).Where("harga_konsumen <= ?", params.MaxHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)
		}
		// check if error is record not found
		if len(response) <= 0 {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("no product found"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}
		// response another error
		if errDb.Error != nil {
			errHelper = &helper.ErrorStruct{
				Err:  errDb.Error,
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
	if params.MinHarga != 0 {
		errDb := db.Where(&filter).Where("harga_konsumen >= ?", params.MinHarga).Limit(params.Limit).Offset(params.Offset).Preload("FotoProduk").Preload("Category").Preload("Toko").Find(&response)

		// check if error is record not found
		if len(response) <= 0 {
			errHelper = &helper.ErrorStruct{
				Err:  errors.New("no product found"),
				Code: http.StatusNotFound,
			}
			return response, errHelper
		}
		// response another error
		if errDb.Error != nil {
			errHelper = &helper.ErrorStruct{
				Err:  errDb.Error,
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
	// check if error is record not found
	if len(response) <= 0 {
		errHelper = &helper.ErrorStruct{
			Err:  errors.New("no product found"),
			Code: http.StatusNotFound,
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
