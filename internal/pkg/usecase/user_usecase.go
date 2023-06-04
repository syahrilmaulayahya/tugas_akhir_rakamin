package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(ctx context.Context, data dto.UserRegisterAndUpdate) (errHelper *helper.ErrorStruct)
	Login(ctx context.Context, data dto.UserLogin) (response dto.UserResponse, errHelper *helper.ErrorStruct)
	GetMyProfile(ctx context.Context, userID uint) (response dto.UserResponse, errHelper *helper.ErrorStruct)
	UpdateProfile(ctx context.Context, ID uint, data dto.UserRegisterAndUpdate) (errHelper *helper.ErrorStruct)
}

type UserUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepository: userRepository,
	}
}

func (uc *UserUseCaseImpl) Register(ctx context.Context, data dto.UserRegisterAndUpdate) (errHelper *helper.ErrorStruct) {
	// validate user input
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errValidate,
		}
		return errHelper
	}

	// hash password before insert to database
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(data.KataSandi), bcrypt.DefaultCost)

	// call register function in user repository and insert register data to get error information
	if errRepo := uc.userRepository.Register(ctx, daos.User{
		Nama:         data.Nama,
		KataSandi:    string(passwordByte),
		Notelp:       data.Notelp,
		TanggalLahir: utils.ParseStringToTime(data.TanggalLahir),
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IDProvinsi:   data.IDProvinsi,
		IDKota:       data.IDKota,
	}); errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Code: errRepo.Code,
			Err:  errRepo.Err,
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

func (uc *UserUseCaseImpl) Login(ctx context.Context, data dto.UserLogin) (response dto.UserResponse, errHelper *helper.ErrorStruct) {

	// validate user input
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errValidate,
		}
		return response, errHelper
	}

	// call login function in user repository and insert login data to get user data and error information
	responseRepo, errRepo := uc.userRepository.Login(ctx, data.Notelp)
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Code: errRepo.Code,
			Err:  errRepo.Err,
		}
		return response, errHelper
	}

	// compare password from input with password from database
	if err := bcrypt.CompareHashAndPassword([]byte(responseRepo.KataSandi), []byte(data.KataSandi)); err != nil {
		errHelper = &helper.ErrorStruct{
			Code: http.StatusUnauthorized,
			Err:  errors.New("No Telp atau kata sandi salah"),
		}
		return response, errHelper
	}

	// mapping user data from database to UserResponse
	response = dto.UserResponse{
		ID:           responseRepo.ID,
		Nama:         responseRepo.Nama,
		NoTelp:       responseRepo.Notelp,
		TanggalLahir: utils.ParseTimeToString(responseRepo.TanggalLahir),
		Tentang:      responseRepo.Tentang,
		Pekerjaan:    responseRepo.Pekerjaan,
		Email:        responseRepo.Email,
		IDProvinsi:   responseRepo.IDProvinsi,
		IDKota:       responseRepo.IDKota,
		IsAdmin:      responseRepo.IsAdmin,
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper
}

func (uc *UserUseCaseImpl) GetMyProfile(ctx context.Context, userID uint) (response dto.UserResponse, errHelper *helper.ErrorStruct) {
	// call GetMyProfile from user repository to  get user data and error information
	responseRepo, errRepo := uc.userRepository.GetMyProfile(ctx, userID)
	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}

	// mapping response from repository to usecase
	response = dto.UserResponse{
		ID:           responseRepo.ID,
		Nama:         responseRepo.Nama,
		NoTelp:       responseRepo.Notelp,
		TanggalLahir: utils.ParseTimeToString(responseRepo.TanggalLahir),
		Tentang:      responseRepo.Tentang,
		Pekerjaan:    responseRepo.Pekerjaan,
		Email:        responseRepo.Email,
		IDProvinsi:   responseRepo.IDProvinsi,
		IDKota:       responseRepo.IDKota,
		IsAdmin:      responseRepo.IsAdmin,
	}
	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return response, errHelper

}
func (uc *UserUseCaseImpl) UpdateProfile(ctx context.Context, ID uint, data dto.UserRegisterAndUpdate) (errHelper *helper.ErrorStruct) {
	//// validate user input
	//if errValidate := helper.Validate.Struct(data); errValidate != nil {
	//	errHelper = &helper.ErrorStruct{
	//		Code: http.StatusBadRequest,
	//		Err:  errValidate,
	//	}
	//	return errHelper
	//}

	// hash password before insert to database
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(data.KataSandi), bcrypt.DefaultCost)

	// call UpdateProfile from user useCase to get error information
	if errRepo := uc.userRepository.UpdateProfile(ctx, daos.User{
		ID:           ID,
		Nama:         data.Nama,
		KataSandi:    string(passwordByte),
		Notelp:       data.Notelp,
		TanggalLahir: utils.ParseStringToTime(data.TanggalLahir),
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IDProvinsi:   data.IDProvinsi,
		IDKota:       data.IDKota,
	}); errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Code: errRepo.Code,
			Err:  errRepo.Err,
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
