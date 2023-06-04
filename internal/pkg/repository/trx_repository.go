package repository

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"time"
)

type TRXRepository interface {
	GetAllTRX(ctx context.Context, userID uint, params daos.FilterTRX) (trx []daos.TRXResponse, errHelper *helper.ErrorStruct)
	GetTRXByID(ctx context.Context, userID, ID uint) (trx daos.TRXResponse, errHelper *helper.ErrorStruct)
	CreateTRX(ctx context.Context, trx daos.TRX, listKuantitasProdukID []daos.ProdukIDKuantitas) (ID uint, errHelper *helper.ErrorStruct)
}

type TRXRepositoryImpl struct {
	db *gorm.DB
}

func NewTRXRepository(db *gorm.DB) TRXRepository {
	return &TRXRepositoryImpl{db: db}
}

func (tr *TRXRepositoryImpl) GetAllTRX(ctx context.Context, userID uint, params daos.FilterTRX) (trx []daos.TRXResponse, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db
	var trxDB []daos.TRX
	// get trx by id
	if errDb := db.Limit(params.Limit).Offset(params.Offset).Preload("DetailTRX").Where("user_id = ?", userID).Find(&trxDB).Error; errDb != nil {

		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return trx, errHelper
	}
	if len(trxDB) <= 0 {
		errHelper = &helper.ErrorStruct{
			Err:  errors.New("no trx found"),
			Code: http.StatusNotFound,
		}
		return []daos.TRXResponse{}, errHelper
	}

	for _, t := range trxDB {
		var alamat daos.Alamat
		if errDb := db.First(&alamat, t.AlamatID).Error; errDb != nil {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusInternalServerError,
			}
			return trx, errHelper
		}
		var listDetailTRX []daos.DetailTRXResponse

		for _, v := range t.DetailTRX {
			var logProduk daos.LogProduk
			if err := db.Where("id = ?", v.LogProdukID).Preload("Toko").Preload("Category").Preload("LogFotoProduk").First(&logProduk).Error; err != nil {
				errHelper = &helper.ErrorStruct{
					Err:  err,
					Code: http.StatusInternalServerError,
				}
				return trx, errHelper
			}
			detailTRXResponse := daos.DetailTRXResponse{
				TRXID:       v.TRXID,
				LogProdukID: v.LogProdukID,
				TokoID:      v.TokoID,
				Toko:        logProduk.Toko,
				Kuantitas:   v.Kuantitas,
				HargaTotal:  v.HargaTotal,
				UpdatedAt:   time.Time{},
				CreatedAt:   time.Time{},
				LogProduk:   logProduk,
			}
			listDetailTRX = append(listDetailTRX, detailTRXResponse)
		}
		transaction := daos.TRXResponse{
			ID:          t.ID,
			UserID:      t.UserID,
			AlamatID:    t.AlamatID,
			Alamat:      alamat,
			HargaTotal:  t.HargaTotal,
			KodeInvoice: t.KodeInvoice,
			MethodBayar: t.MethodBayar,
			UpdatedAt:   time.Time{},
			CreatedAt:   time.Time{},
			DetailTRX:   listDetailTRX,
		}
		trx = append(trx, transaction)
	}

	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return trx, errHelper
}

func (tr *TRXRepositoryImpl) GetTRXByID(ctx context.Context, userID, ID uint) (trx daos.TRXResponse, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db
	var trxDB daos.TRX
	// get trx by id
	if errDb := db.Preload("DetailTRX").Where("user_id = ? AND id = ?", userID, ID).First(&trxDB).Error; errDb != nil {
		if errDb == gorm.ErrRecordNotFound {
			errHelper = &helper.ErrorStruct{
				Err:  errDb,
				Code: http.StatusNotFound,
			}
			return trx, errHelper
		}
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return trx, errHelper
	}
	var alamat daos.Alamat
	if errDb := db.First(&alamat, trxDB.AlamatID).Error; errDb != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errDb,
			Code: http.StatusInternalServerError,
		}
		return trx, errHelper
	}
	var listDetailTRX []daos.DetailTRXResponse

	for _, v := range trxDB.DetailTRX {
		var logProduk daos.LogProduk
		if err := db.Where("id = ?", v.LogProdukID).Preload("Toko").Preload("Category").Preload("LogFotoProduk").First(&logProduk).Error; err != nil {
			errHelper = &helper.ErrorStruct{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
			return trx, errHelper
		}
		detailTRXResponse := daos.DetailTRXResponse{
			TRXID:       v.TRXID,
			LogProdukID: v.LogProdukID,
			TokoID:      v.TokoID,
			Toko:        logProduk.Toko,
			Kuantitas:   v.Kuantitas,
			HargaTotal:  v.HargaTotal,
			UpdatedAt:   time.Time{},
			CreatedAt:   time.Time{},
			LogProduk:   logProduk,
		}
		listDetailTRX = append(listDetailTRX, detailTRXResponse)
	}
	trx = daos.TRXResponse{
		ID:          trxDB.ID,
		UserID:      trxDB.UserID,
		AlamatID:    trxDB.AlamatID,
		Alamat:      alamat,
		HargaTotal:  trxDB.HargaTotal,
		KodeInvoice: trxDB.KodeInvoice,
		MethodBayar: trxDB.MethodBayar,
		UpdatedAt:   time.Time{},
		CreatedAt:   time.Time{},
		DetailTRX:   listDetailTRX,
	}
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return trx, errHelper
}

func (tr *TRXRepositoryImpl) CreateTRX(ctx context.Context, trx daos.TRX, listProdukIDKuantitas []daos.ProdukIDKuantitas) (ID uint, errHelper *helper.ErrorStruct) {
	// get gorm client
	db := tr.db

	// start transaction
	errTrans := db.Transaction(func(tx *gorm.DB) error {

		var listNewDetailTRX []daos.DetailTRX
		var hargaTotalTRX uint
		for _, v := range listProdukIDKuantitas {
			produk := daos.Produk{}

			// update stok produk
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", v.ProdukID).Preload("FotoProduk").First(&produk).Error; err != nil {
				return err
			}
			if produk.TokoID == trx.UserID {
				return errors.New("user cannot buy their own items")
			}
			if int(produk.Stok)-int(v.Kuantitas) <= 0 {
				return errors.New("not enough stock")
			}
			if err := tx.Model(&daos.Produk{}).Where("id=?", v.ProdukID).Update("stok", produk.Stok-v.Kuantitas).Error; err != nil {
				return err
			}

			newLogProduk := daos.LogProduk{
				ProdukID:      v.ProdukID,
				NamaProduk:    produk.NamaProduk,
				Slug:          produk.Slug,
				HargaReseller: produk.HargaReseller,
				HargaKonsumen: produk.HargaKonsumen,
				Deskripsi:     produk.Deskripsi,
				TokoID:        produk.TokoID,
				CategoryID:    produk.CategoryID,
			}

			if err := tx.Create(&newLogProduk).Error; err != nil {
				return err
			}
			var listLogFotoProduk []daos.LogFotoProduk
			if len(produk.FotoProduk) > 0 {
				for _, f := range produk.FotoProduk {
					logFotoProduk := daos.LogFotoProduk{
						LogProdukID: newLogProduk.ID,
						URL:         f.URL,
					}
					listLogFotoProduk = append(listLogFotoProduk, logFotoProduk)
				}
				if err := tx.Create(&listLogFotoProduk).Error; err != nil {
					return err
				}
			}

			newDetailTRX := daos.DetailTRX{
				LogProdukID: newLogProduk.ID,
				TokoID:      newLogProduk.TokoID,
				Kuantitas:   v.Kuantitas,
				HargaTotal:  produk.HargaKonsumen * v.Kuantitas,
			}
			hargaTotalTRX += newDetailTRX.HargaTotal
			listNewDetailTRX = append(listNewDetailTRX, newDetailTRX)
		}
		newTRX := daos.TRX{
			UserID:      trx.UserID,
			AlamatID:    trx.AlamatID,
			HargaTotal:  hargaTotalTRX,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
		}
		if err := tx.Create(&newTRX).Error; err != nil {
			return err
		}
		ID = newTRX.ID

		for i, _ := range listNewDetailTRX {
			listNewDetailTRX[i].TRXID = newTRX.ID
		}
		if err := tx.Create(&listNewDetailTRX).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	// error checking
	if errTrans != nil {
		if errTrans.Error() == "not enough stock" || errTrans.Error() == "user cannot buy their own items" {
			errHelper = &helper.ErrorStruct{
				Err:  errTrans,
				Code: http.StatusBadRequest,
			}
			return ID, errHelper
		}
		errHelper = &helper.ErrorStruct{
			Err:  errTrans,
			Code: http.StatusInternalServerError,
		}
		return ID, errHelper
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return ID, errHelper
}
