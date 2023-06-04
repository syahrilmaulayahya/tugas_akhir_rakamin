package usecase

import (
	"context"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"net/http"
)

type AlamatUseCase interface {
	CreateAlamat(ctx context.Context, userID uint, data dto.CreateAlamatRequest) (ID uint, errHelper *helper.ErrorStruct)
	GetMyAlamat(ctx context.Context, userID uint, params dto.FilterAlamat) (response []dto.Alamat, errHelper *helper.ErrorStruct)
	GetAlamatByID(ctx context.Context, ID uint) (response dto.Alamat, errHelper *helper.ErrorStruct)
	UpdateAlamatByID(ctx context.Context, userID, ID uint, data dto.UpdateAlamatRequest) (errHelper *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, userID, ID uint) (errHelper *helper.ErrorStruct)
}

type AlamatUseCaseImpl struct {
	alamatRepository repository.AlamatRepository
}

func NewAlamatUseCase(alamatRepository repository.AlamatRepository) AlamatUseCase {
	return &AlamatUseCaseImpl{alamatRepository: alamatRepository}
}

func (au *AlamatUseCaseImpl) CreateAlamat(ctx context.Context, userID uint, data dto.CreateAlamatRequest) (ID uint, errHelper *helper.ErrorStruct) {
	// validate user input
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errValidate,
			Code: http.StatusBadRequest,
		}
		return ID, errHelper
	}

	// call CreateAlamat from alamat repository to get id new inserted alamat and error information
	IDRepository, errRepository := au.alamatRepository.CreateAlamat(ctx, daos.Alamat{
		UserID:       userID,
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NoTelp,
		DetailAlamat: data.DetailAlamat,
	})

	// error checking
	if errRepository.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepository.Err,
			Code: errRepository.Code,
		}
		return ID, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return IDRepository, errHelper
}

func (au *AlamatUseCaseImpl) GetMyAlamat(ctx context.Context, userID uint, params dto.FilterAlamat) (response []dto.Alamat, errHelper *helper.ErrorStruct) {
	// call GetMyAlamat from alamat repository to get list alamat and error information
	responseRepository, errRepository := au.alamatRepository.GetMyAlamat(ctx, userID, daos.FilterAlamat{JudulAlamat: params.JudulAlamat})
	if errRepository.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errHelper.Err,
			Code: errHelper.Code,
		}
		return response, errHelper
	}

	// mapping from daos alamat to dto alamat
	if len(responseRepository) > 0 {
		for _, v := range responseRepository {
			alamat := dto.Alamat{
				ID:           v.ID,
				JudulAlamat:  v.JudulAlamat,
				NamaPenerima: v.NamaPenerima,
				NoTelp:       v.NoTelp,
				DetailAlamat: v.NoTelp,
			}
			response = append(response, alamat)
		}
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (au *AlamatUseCaseImpl) GetAlamatByID(ctx context.Context, ID uint) (response dto.Alamat, errHelper *helper.ErrorStruct) {
	// call GetAlamatByID from alamat repository to get alamat record and error information
	responseRepo, errRepo := au.alamatRepository.GetAlamatByID(ctx, ID)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	response = dto.Alamat{
		ID:           responseRepo.ID,
		JudulAlamat:  responseRepo.JudulAlamat,
		NamaPenerima: responseRepo.NamaPenerima,
		NoTelp:       responseRepo.NoTelp,
		DetailAlamat: responseRepo.DetailAlamat,
	}
	return response, errHelper
}

func (au *AlamatUseCaseImpl) UpdateAlamatByID(ctx context.Context, userID, ID uint, data dto.UpdateAlamatRequest) (errHelper *helper.ErrorStruct) {
	// validate user input
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errValidate,
			Code: http.StatusBadRequest,
		}
		return errHelper
	}
	// call UpdateAlamatByID from alamat repository to update alamat record and get error information
	if errUseCase := au.alamatRepository.UpdateAlamatByID(ctx, daos.Alamat{
		ID:           ID,
		UserID:       userID,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NoTelp,
		DetailAlamat: data.DetailAlamat,
	}); errUseCase.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errUseCase.Err,
			Code: errUseCase.Code,
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

func (au *AlamatUseCaseImpl) DeleteAlamatByID(ctx context.Context, userID, ID uint) (errHelper *helper.ErrorStruct) {
	// call DeleteAlamatByID from alamat repository to delete alamat record and get error information
	if errRepo := au.alamatRepository.DeleteAlamatByID(ctx, userID, ID); errRepo.Err != nil {
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
