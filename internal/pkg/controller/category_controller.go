package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"strconv"
)

type CategoryController interface {
	CreateCategory(ctx *fiber.Ctx) (err error)
	GetAllCategory(ctx *fiber.Ctx) (err error)
	GetCategoryByID(ctx *fiber.Ctx) (err error)
	UpdateCategoryByID(ctx *fiber.Ctx) (err error)
	DeleteCategoryByID(ctx *fiber.Ctx) (err error)
}

type CategoryControllerImpl struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(categoryUseCase usecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		categoryUseCase: categoryUseCase,
	}
}

func (cc *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) (err error) {

	// get user input
	data := new(dto.CreateAndUpdateCategoryRequest)
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed To POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call CreateCategory from category useCase to get new inserted data id and error information
	c := ctx.Context()
	idUseCase, errUseCase := cc.categoryUseCase.CreateCategory(c, *data)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed To POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    idUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (cc *CategoryControllerImpl) GetAllCategory(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	// call GetAllCategory function from category UseCase to get all categories and check if there's error
	responseUseCase, errUseCase := cc.categoryUseCase.GetAllCategory(c)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (cc *CategoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	// get id from url parameter
	GetParams := ctx.Params("id")

	// check if id valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetCategoryByID function from category useCase to get category record and error information
	responseUseCase, errUseCase := cc.categoryUseCase.GetCategoryByID(c, uint(ID))
	// error checking
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(errUseCase.Code).JSON(response)
}

func (cc *CategoryControllerImpl) UpdateCategoryByID(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	// get url parameter
	GetParams := ctx.Params("id")

	// check if id is valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	data := new(dto.CreateAndUpdateCategoryRequest)

	// get user input
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call UpdateCategoryByID from user useCase to get error information and check if there's error
	if errUseCase := cc.categoryUseCase.UpdateCategoryByID(c, uint(ID), dto.CreateAndUpdateCategoryRequest{
		NamaCategory: data.NamaCategory,
	}); errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to PUT data",
		Error:   nil,
		Data:    "",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (cc *CategoryControllerImpl) DeleteCategoryByID(ctx *fiber.Ctx) (err error) {
	// get id from url parameter
	GetParams := ctx.Params("id")
	// check if id is valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// get DeleteCategoryByID from category useCase to get error information and check if there's error
	c := ctx.Context()
	if errUseCase := cc.categoryUseCase.DeleteCategoryByID(c, dto.CategoryIDOnly{
		ID: uint(ID),
	}); errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to DELETE data",
		Error:   nil,
		Data:    "",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
