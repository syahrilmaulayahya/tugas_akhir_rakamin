package usecase

import (
	"context"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"net/http"
)

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, data dto.CreateAndUpdateCategoryRequest) (id uint, errHelper *helper.ErrorStruct)
	GetAllCategory(ctx context.Context) (response []dto.CategoryWithID, errHelper *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, ID uint) (response dto.CategoryWithID, errHelper *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, ID uint, data dto.CreateAndUpdateCategoryRequest) (errHelper *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, data dto.CategoryIDOnly) (errHelper *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (cu *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CreateAndUpdateCategoryRequest) (id uint, errHelper *helper.ErrorStruct) {

	// validate user input
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errValidate,
			Code: http.StatusBadRequest,
		}
		return id, errHelper
	}

	// call CreateCategory function from category repository to get id new category record and error information
	idRepo, errRepo := cu.categoryRepository.CreateCategory(ctx, daos.Category{NamaCategory: data.NamaCategory})

	// check error from repository
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return id, errHelper
	}

	// success response
	id = idRepo
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: http.StatusOK,
	}
	return id, errHelper
}

func (cu *CategoryUseCaseImpl) GetAllCategory(ctx context.Context) (response []dto.CategoryWithID, errHelper *helper.ErrorStruct) {

	// call GetAllCategory from category repository to get all category
	responseRepo, errRepo := cu.categoryRepository.GetAllCategory(ctx)

	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}

	// check if responseRepo not nil
	if len(responseRepo) > 0 {
		// mapping response from repository to response useCase
		for _, v := range responseRepo {
			category := dto.CategoryWithID{
				ID:           v.ID,
				NamaCategory: v.NamaCategory,
			}
			response = append(response, category)
		}
	}

	// success response
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: errRepo.Code,
	}
	return response, errHelper
}

func (cu *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, ID uint) (response dto.CategoryWithID, errHelper *helper.ErrorStruct) {
	// call GetCategoryByID function from category repository to get category record and error information
	responseRepo, errRepo := cu.categoryRepository.GetCategoryByID(ctx, ID)

	// error checking
	if errRepo.Err != nil {
		errHelper = &helper.ErrorStruct{
			Err:  errRepo.Err,
			Code: errRepo.Code,
		}
		return response, errHelper
	}

	// success response
	response = dto.CategoryWithID{
		ID:           responseRepo.ID,
		NamaCategory: responseRepo.NamaCategory,
	}
	errHelper = &helper.ErrorStruct{
		Err:  nil,
		Code: errRepo.Code,
	}
	return response, errHelper
}

func (cu *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, ID uint, data dto.CreateAndUpdateCategoryRequest) (errHelper *helper.ErrorStruct) {

	// call UpdateCategoryByID function from repository to get error from repository and check if there's error
	if errRepo := cu.categoryRepository.UpdateCategoryByID(ctx, daos.Category{
		ID:           ID,
		NamaCategory: data.NamaCategory,
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

func (cu *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, data dto.CategoryIDOnly) (errHelper *helper.ErrorStruct) {

	// call DeleteCategoryByID from category repository to get error from repository
	if errRepo := cu.categoryRepository.DeleteCategoryByID(ctx, daos.Category{ID: data.ID}); errRepo.Err != nil {
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
