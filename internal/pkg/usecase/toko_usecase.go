package usecase

import (
	"context"
	"errors"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"net/http"
)

type TokoUseCase interface {
	GetTokoByID(ctx context.Context, ID uint) (response dto.GetTokoByIDResponse, errHelper *helper.ErrorStruct)
	GetTokoByUserID(ctx context.Context, userID uint) (response dto.GetTokoByUserIDResponse, errHelper *helper.ErrorStruct)
	GetAllToko(ctx context.Context, params dto.TokoFilter) (response []dto.GetAllTokoResponse, errHelper *helper.ErrorStruct)
	UpdateToko(ctx context.Context, userID uint, data dto.UpdateTokoRequest) (errHelper *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokoRepository repository.TokoRepository
}

func NewTokoUseCase(tokoRepository repository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{tokoRepository: tokoRepository}
}

func (tu *TokoUseCaseImpl) GetTokoByID(ctx context.Context, ID uint) (response dto.GetTokoByIDResponse, errHelper *helper.ErrorStruct) {

	// call GetTokoByID function from toko repository to get toko record and error information
	responseRepo, errRepo := tu.tokoRepository.GetTokoByID(ctx, ID)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Code: errRepo.Code,
			Err:  errRepo.Err,
		}
		return response, errHelper
	}

	// success response
	// mapping response from repository
	response = dto.GetTokoByIDResponse{
		ID:       responseRepo.ID,
		NamaToko: responseRepo.NamaToko,
		UrlFoto:  responseRepo.UrlFoto,
	}
	errHelper = &helper.ErrorStruct{
		Err:  errRepo.Err,
		Code: errRepo.Code,
	}
	return response, errHelper
}

func (tu *TokoUseCaseImpl) GetTokoByUserID(ctx context.Context, userID uint) (response dto.GetTokoByUserIDResponse, errHelper *helper.ErrorStruct) {

	// call GetTokoByID function from toko repository
	responseRepo, errRepo := tu.tokoRepository.GetTokoByUserID(ctx, userID)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Code: errRepo.Code,
			Err:  errRepo.Err,
		}
		return response, errHelper
	}

	// success response
	// mapping response from repository
	response = dto.GetTokoByUserIDResponse{
		ID:       responseRepo.ID,
		NamaToko: responseRepo.NamaToko,
		UrlFoto:  responseRepo.UrlFoto,
		UserID:   responseRepo.UserID,
	}
	errHelper = &helper.ErrorStruct{
		Err:  errRepo.Err,
		Code: errRepo.Code,
	}
	return response, errHelper
}

func (tu *TokoUseCaseImpl) GetAllToko(ctx context.Context, params dto.TokoFilter) (response []dto.GetAllTokoResponse, errHelper *helper.ErrorStruct) {
	// setup pagination
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	// call GetAllToko function from toko repository to get all toko and error information
	responseRepo, errRepo := tu.tokoRepository.GetAllToko(ctx, daos.FilterToko{
		Limit:  params.Limit,
		Offset: params.Page,
		Nama:   params.Nama,
	})
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}
	// check if length of toko data from repository > 0
	if len(responseRepo) > 0 {
		for _, v := range responseRepo {
			response = append(response, dto.GetAllTokoResponse{
				ID:       v.ID,
				NamaToko: v.NamaToko,
				UrlFoto:  v.UrlFoto,
			})
		}
	} else {
		// check if toko not found
		errHelper = &helper.ErrorStruct{
			Err:  errors.New("belum ada toko"),
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

func (tu *TokoUseCaseImpl) UpdateToko(ctx context.Context, userID uint, data dto.UpdateTokoRequest) (errHelper *helper.ErrorStruct) {
	// call UpdateToko function from toko repository to update data and get err information
	if errRepo := tu.tokoRepository.UpdateToko(ctx, daos.Toko{

		UserID:   userID,
		NamaToko: data.NamaToko,
		UrlFoto:  data.Photo,
	}); errRepo.Err != nil {
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
