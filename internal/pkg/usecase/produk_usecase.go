package usecase

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"net/http"
	"strings"
)

type ProdukUseCase interface {
	UploadProduk(ctx context.Context, data dto.UploadProdukRequest) (ID uint, errHelper *helper.ErrorStruct)
	GetProdukByID(ctx context.Context, ID uint) (response dto.GetProduk, errHelper *helper.ErrorStruct)
	UpdateProdukByID(ctx context.Context, data dto.UpdateProdukRequest) (errHelper *helper.ErrorStruct)
	DeleteProdukByID(ctx context.Context, tokoID, ID uint) (errHelper *helper.ErrorStruct)
	GetAllProduk(ctx context.Context, params dto.FilterProduk) (response []dto.GetProduk, errHelper *helper.ErrorStruct)
}

type ProdukUseCaseImpl struct {
	produkRepository repository.ProdukRepository
}

func NewProdukUseCase(produkRepository repository.ProdukRepository) ProdukUseCase {
	return &ProdukUseCaseImpl{produkRepository: produkRepository}
}

func (pu *ProdukUseCaseImpl) UploadProduk(ctx context.Context, data dto.UploadProdukRequest) (ID uint, errHelper *helper.ErrorStruct) {
	// validate user input
	if data.NamaProduk == "" || data.Deskripsi == "" {
		errHelper = &helper.ErrorStruct{
			Err:  errors.New("nama_produk and deskripsi are required"),
			Code: http.StatusBadRequest,
		}
		return ID, errHelper
	}

	// mapping foto url
	var listPhotos []daos.FotoProduk
	for _, v := range data.Photos {
		listPhotos = append(listPhotos, daos.FotoProduk{
			URL: v.URL,
		})
	}

	// format slug
	slug := strings.ToLower(data.NamaProduk)
	listSlug := strings.Split(slug, " ")
	slug = strings.Join(listSlug, "-")

	// Call UploadProduk function from produk repository to create new record in database and get ID new inserted record and error information
	IDUseCase, errUseCase := pu.produkRepository.UploadProduk(ctx, daos.Produk{
		CategoryID:    data.CategoryID,
		TokoID:        data.TokoID,
		NamaProduk:    data.NamaProduk,
		Slug:          slug,
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		FotoProduk:    listPhotos,
	})
	// error checking UploadProduk useCase
	if errUseCase.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errUseCase.Err,
			Code: errUseCase.Code,
		}
		return ID, errUseCase
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return IDUseCase, errHelper

}

func (pu *ProdukUseCaseImpl) GetProdukByID(ctx context.Context, ID uint) (response dto.GetProduk, errHelper *helper.ErrorStruct) {
	// call GetProdukByID from user repository
	responseRepo, errRepo := pu.produkRepository.GetProdukByID(ctx, ID)
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}

	// success response

	// mapping foto produk from daos to dto
	var listFoto []dto.FotoProdukGetProduk
	for _, v := range responseRepo.FotoProduk {
		foto := dto.FotoProdukGetProduk{
			ID:       v.ID,
			ProdukID: v.ProdukID,
			URL:      v.URL,
		}
		listFoto = append(listFoto, foto)
	}
	// mapping toko from daos to dto
	toko := dto.GetTokoByIDResponse{
		ID:       responseRepo.Toko.ID,
		NamaToko: responseRepo.Toko.NamaToko,
		UrlFoto:  responseRepo.Toko.UrlFoto,
	}
	// mapping category from daos to dto
	category := dto.CategoryWithID{
		ID:           responseRepo.Category.ID,
		NamaCategory: responseRepo.Category.NamaCategory,
	}
	// mapping response from db to local struct
	response = dto.GetProduk{
		ID:            responseRepo.ID,
		NamaProduk:    responseRepo.NamaProduk,
		Slug:          responseRepo.Slug,
		HargaReseller: responseRepo.HargaReseller,
		HargaKonsumen: responseRepo.HargaKonsumen,
		Stok:          responseRepo.Stok,
		Deskripsi:     responseRepo.Deskripsi,
		Toko:          toko,
		Category:      category,
		FotoProduk:    listFoto,
	}
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (pu *ProdukUseCaseImpl) UpdateProdukByID(ctx context.Context, data dto.UpdateProdukRequest) (errHelper *helper.ErrorStruct) {
	// edit slug
	slug := strings.ToLower(data.NamaProduk)
	listSlug := strings.Split(slug, " ")
	slug = strings.Join(listSlug, "-")

	// mapping foto data from dto to daos
	var listFoto []daos.FotoProduk
	for _, v := range data.Photos {
		foto := daos.FotoProduk{
			ProdukID: data.ID,
			URL:      v.URL,
		}
		listFoto = append(listFoto, foto)
	}

	// call GetProdukByID from user repository
	errRepo := pu.produkRepository.UpdateProdukByID(ctx, daos.Produk{
		ID:            data.ID,
		NamaProduk:    data.NamaProduk,
		Slug:          slug,
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		TokoID:        data.TokoID,
		CategoryID:    data.CategoryID,
		FotoProduk:    listFoto,
	})
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
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

func (pu *ProdukUseCaseImpl) DeleteProdukByID(ctx context.Context, tokoID, ID uint) (errHelper *helper.ErrorStruct) {
	// call DeleteProdukByID form ProdukRepository to get error information
	errRepo := pu.produkRepository.DeleteProdukByID(ctx, tokoID, ID)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
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

func (pu *ProdukUseCaseImpl) GetAllProduk(ctx context.Context, params dto.FilterProduk) (response []dto.GetProduk, errHelper *helper.ErrorStruct) {
	// setup pagination
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	// call GetAllProduk from produk repository to get all produk
	responseRepo, errRepo := pu.produkRepository.GetAllProduk(ctx, daos.FilterProduk{
		Limit:      params.Limit,
		Offset:     params.Page,
		NamaProduk: params.NamaProduk,
		CategoryID: params.CategoryID,
		TokoID:     params.TokoID,
		MaxHarga:   params.MaxHarga,
		MinHarga:   params.MinHarga,
	})
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}
	for _, v := range responseRepo {
		var listFoto []dto.FotoProdukGetProduk

		for _, f := range v.FotoProduk {
			foto := dto.FotoProdukGetProduk{
				ID:       f.ID,
				ProdukID: f.ProdukID,
				URL:      f.URL,
			}
			listFoto = append(listFoto, foto)
		}
		produk := dto.GetProduk{
			ID:            v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseller: v.HargaReseller,
			HargaKonsumen: v.HargaKonsumen,
			Stok:          v.Stok,
			Deskripsi:     v.Deskripsi,
			Toko: dto.GetTokoByIDResponse{
				ID:       v.Toko.ID,
				NamaToko: v.Toko.NamaToko,
				UrlFoto:  v.Toko.UrlFoto,
			},
			Category: dto.CategoryWithID{
				ID:           v.Category.ID,
				NamaCategory: v.Category.NamaCategory,
			},
			FotoProduk: listFoto,
		}
		response = append(response, produk)
	}
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}
