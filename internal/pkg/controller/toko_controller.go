package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"github.com/valyala/fasthttp"
	"strconv"
)

type TokoController interface {
	GetTokoByID(ctx *fiber.Ctx) (err error)
	GetMyToko(ctx *fiber.Ctx) (err error)
	GetAllToko(ctx *fiber.Ctx) (err error)
	UpdateToko(ctx *fiber.Ctx) (err error)
}

type TokoControllerImpl struct {
	tokoUseCase usecase.TokoUseCase
}

func NewTokoController(tokoUseCase usecase.TokoUseCase) TokoController {
	return &TokoControllerImpl{tokoUseCase: tokoUseCase}
}

func (tc *TokoControllerImpl) GetTokoByID(ctx *fiber.Ctx) (err error) {
	ID := 0
	// get id_toko from url parameters
	GetParams := ctx.Params("id_toko")

	// convert id string to int
	ID, err = strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetTokoByID function from toko useCase to get toko data and error information
	c := ctx.Context()
	responseUseCase, errUseCase := tc.tokoUseCase.GetTokoByID(c, uint(ID))
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET DATA",
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

func (tc *TokoControllerImpl) GetMyToko(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()
	ID := 0

	// get local context from middleware
	getLocalContext := ctx.Locals("userID")
	ID, _ = strconv.Atoi(fmt.Sprintf("%v", getLocalContext))

	// call GetTokoByID function from toko useCase
	responseUseCase, errUseCase := tc.tokoUseCase.GetTokoByUserID(c, uint(ID))
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET DATA",
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

func (tc *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) (err error) {
	// get limit and page from query parameter url
	filter := new(dto.TokoFilter)
	if errQuery := ctx.QueryParser(filter); errQuery != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errQuery.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call GetAllToko from toko useCase to get all toko and error information
	c := ctx.Context()
	responseUseCase, errUseCase := tc.tokoUseCase.GetAllToko(c, dto.TokoFilter{
		Limit: filter.Limit,
		Page:  filter.Page,
		Nama:  filter.Nama,
	})
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
	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (tc *TokoControllerImpl) UpdateToko(ctx *fiber.Ctx) (err error) {
	userID := 0

	// get local context from middleware
	getLocalContext := ctx.Locals("userID")
	userID, _ = strconv.Atoi(fmt.Sprintf("%v", getLocalContext))

	// get new nama toko from user input
	namaToko := ctx.FormValue("nama_toko")
	// get photo file
	file, errFile := ctx.FormFile("photo")
	var filename string
	if errFile != nil {
		if errFile == fasthttp.ErrMissingFile {
			filename = ""

		} else {
			var response = BaseResponse{
				Status:  false,
				Message: "Failed to POST data",
				Error:   []string{errFile.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}
	}
	if file != nil {
		filename = file.Filename
		if errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/toko/%s", filename)); errSaveFile != nil {
			response := BaseResponse{
				Status:  false,
				Message: "Failed to POST data",
				Error:   []string{errSaveFile.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	// call UpdateToko from toko useCase to update toko record and get error information
	c := ctx.Context()
	if errUseCase := tc.tokoUseCase.UpdateToko(c, uint(userID), dto.UpdateTokoRequest{
		NamaToko: namaToko,
		Photo:    filename,
	}); errUseCase.Err != nil {
		response := BaseResponse{
			Status:  true,
			Message: "Failed to POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to UPDATE data",
		Error:   nil,
		Data:    "Update toko succeed",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
