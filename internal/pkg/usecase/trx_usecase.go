package usecase

import (
	"context"
	"fmt"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"math/rand"
	"net/http"
	"sort"
)

type TRXUseCase interface {
	GetAllTRX(ctx context.Context, userID uint, params dto.FilterTRX) (trx []dto.TRXGetResponse, errHelper *helper.ErrorStruct)
	GetTRXByID(ctx context.Context, userID, ID uint) (trx dto.TRXGetResponse, errHelper *helper.ErrorStruct)
	CreateTRX(ctx context.Context, trx dto.TRX) (ID uint, errHelper *helper.ErrorStruct)
}

type TRXUseCaseImpl struct {
	trxRepository repository.TRXRepository
}

func NewTRXUseCase(trxRepository repository.TRXRepository) TRXUseCase {
	return &TRXUseCaseImpl{trxRepository: trxRepository}
}

func (trxu *TRXUseCaseImpl) GetAllTRX(ctx context.Context, userID uint, params dto.FilterTRX) (trx []dto.TRXGetResponse, errHelper *helper.ErrorStruct) {
	// setup pagination
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	// call GetTRXByID from trx repository
	trxRepo, errRepo := trxu.trxRepository.GetAllTRX(ctx, userID, daos.FilterTRX{
		Limit:  params.Limit,
		Offset: params.Page,
	})
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return trx, errHelper
	}

	for _, t := range trxRepo {
		var listDetailTRX []dto.DetailTRXGetResponse
		for _, v := range t.DetailTRX {
			var listFoto []dto.LogFotoProdukResponse
			for _, f := range v.LogProduk.LogFotoProduk {
				foto := dto.LogFotoProdukResponse{
					ID:          f.ID,
					LogProdukID: f.LogProdukID,
					URL:         f.URL,
				}
				listFoto = append(listFoto, foto)
			}
			logProduk := dto.LogProdukGetResponse{
				ID:            v.LogProduk.ID,
				ProdukID:      v.LogProduk.ProdukID,
				NamaProduk:    v.LogProduk.NamaProduk,
				Slug:          v.LogProduk.Slug,
				HargaReseller: v.LogProduk.HargaReseller,
				HargaKonsumen: v.LogProduk.HargaKonsumen,
				Deskripsi:     v.LogProduk.Deskripsi,

				Toko: dto.TokoTRX{
					NamaToko: v.LogProduk.Toko.NamaToko,
					UrlFoto:  v.LogProduk.Toko.UrlFoto,
				},
				Category: dto.CategoryWithID{
					ID:           v.LogProduk.Category.ID,
					NamaCategory: v.LogProduk.Category.NamaCategory,
				},
				Photos: listFoto,
			}
			detailTRX := dto.DetailTRXGetResponse{
				Product: logProduk,
				Toko: dto.GetTokoByIDResponse{
					ID:       v.Toko.ID,
					NamaToko: v.Toko.NamaToko,
					UrlFoto:  v.Toko.UrlFoto,
				},
				Kuantitas:  v.Kuantitas,
				HargaTotal: v.HargaTotal,
			}
			listDetailTRX = append(listDetailTRX, detailTRX)
		}
		alamat := dto.AlamatTRX{
			ID:           t.Alamat.ID,
			JudulAlamat:  t.Alamat.JudulAlamat,
			NamaPenerima: t.Alamat.NamaPenerima,
			NoTelp:       t.Alamat.NoTelp,
			DetailAlamat: t.Alamat.DetailAlamat,
		}
		transaction := dto.TRXGetResponse{
			ID:          t.ID,
			HargaTotal:  t.HargaTotal,
			KodeInvoice: t.KodeInvoice,
			MethodBayar: t.MethodBayar,
			Alamat:      alamat,
			DetailTRX:   listDetailTRX,
		}
		trx = append(trx, transaction)
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return trx, errHelper
}

func (trxu *TRXUseCaseImpl) GetTRXByID(ctx context.Context, userID, ID uint) (trx dto.TRXGetResponse, errHelper *helper.ErrorStruct) {
	// call GetTRXByID from trx repository
	trxRepo, errRepo := trxu.trxRepository.GetTRXByID(ctx, userID, ID)
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return trx, errHelper
	}

	var listDetailTRX []dto.DetailTRXGetResponse
	for _, v := range trxRepo.DetailTRX {
		var listFoto []dto.LogFotoProdukResponse
		for _, f := range v.LogProduk.LogFotoProduk {
			foto := dto.LogFotoProdukResponse{
				ID:          f.ID,
				LogProdukID: f.LogProdukID,
				URL:         f.URL,
			}
			listFoto = append(listFoto, foto)
		}
		logProduk := dto.LogProdukGetResponse{
			ID:            v.LogProduk.ID,
			ProdukID:      v.LogProduk.ProdukID,
			NamaProduk:    v.LogProduk.NamaProduk,
			Slug:          v.LogProduk.Slug,
			HargaReseller: v.LogProduk.HargaReseller,
			HargaKonsumen: v.LogProduk.HargaKonsumen,
			Deskripsi:     v.LogProduk.Deskripsi,

			Toko: dto.TokoTRX{
				NamaToko: v.LogProduk.Toko.NamaToko,
				UrlFoto:  v.LogProduk.Toko.UrlFoto,
			},
			Category: dto.CategoryWithID{
				ID:           v.LogProduk.Category.ID,
				NamaCategory: v.LogProduk.Category.NamaCategory,
			},
			Photos: listFoto,
		}
		detailTRX := dto.DetailTRXGetResponse{
			Product: logProduk,
			Toko: dto.GetTokoByIDResponse{
				ID:       v.Toko.ID,
				NamaToko: v.Toko.NamaToko,
				UrlFoto:  v.Toko.UrlFoto,
			},
			Kuantitas:  v.Kuantitas,
			HargaTotal: v.HargaTotal,
		}
		listDetailTRX = append(listDetailTRX, detailTRX)
	}
	alamat := dto.AlamatTRX{
		ID:           trxRepo.Alamat.ID,
		JudulAlamat:  trxRepo.Alamat.JudulAlamat,
		NamaPenerima: trxRepo.Alamat.NamaPenerima,
		NoTelp:       trxRepo.Alamat.NoTelp,
		DetailAlamat: trxRepo.Alamat.DetailAlamat,
	}
	trx = dto.TRXGetResponse{
		ID:          trxRepo.ID,
		HargaTotal:  trxRepo.HargaTotal,
		KodeInvoice: trxRepo.KodeInvoice,
		MethodBayar: trxRepo.MethodBayar,
		Alamat:      alamat,
		DetailTRX:   listDetailTRX,
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return trx, errHelper
}

func (trxu *TRXUseCaseImpl) CreateTRX(ctx context.Context, trx dto.TRX) (ID uint, errHelper *helper.ErrorStruct) {
	// validate user input
	if errValidate := helper.Validate.Struct(trx); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errValidate,
		}
		return ID, errHelper
	}

	// sort detail
	sort.SliceStable(trx.DetailTRX, func(i, j int) bool {
		return trx.DetailTRX[i].ProductID < trx.DetailTRX[j].ProductID
	})
	// create kode invoice
	randInt := 1000000000 + rand.Intn((9999999999 - 1000000000))
	kodeInvoice := fmt.Sprintf("INV-%d", randInt)
	// convert DetailTRX dto to ProdukIDKuantitas daos
	var listProdukIDKuantitas []daos.ProdukIDKuantitas
	for _, v := range trx.DetailTRX {
		produkIDKuantitas := daos.ProdukIDKuantitas{
			ProdukID:  v.ProductID,
			Kuantitas: v.Kuantitas,
		}
		listProdukIDKuantitas = append(listProdukIDKuantitas, produkIDKuantitas)
	}
	// call CreateTRX from trx repository
	IDRepo, errRepo := trxu.trxRepository.CreateTRX(ctx, daos.TRX{

		UserID:      trx.UserID,
		AlamatID:    trx.AlamatID,
		KodeInvoice: kodeInvoice,
		MethodBayar: trx.MethodBayar,
	}, listProdukIDKuantitas)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return ID, errHelper
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return IDRepo, errHelper
}
