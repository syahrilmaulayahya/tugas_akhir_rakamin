package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"strconv"
)

type AlamatController interface {
	CreateAlamat(ctx *fiber.Ctx) (err error)
	GetMyAlamat(ctx *fiber.Ctx) (err error)
	GetAlamatByID(ctx *fiber.Ctx) (err error)
	UpdateAlamatByID(ctx *fiber.Ctx) (err error)
	DeleteAlamatByID(ctx *fiber.Ctx) (err error)
}

type AlamatControllerImpl struct {
	alamatUseCase usecase.AlamatUseCase
}

func NewAlamatController(alamatUseCase usecase.AlamatUseCase) AlamatController {
	return &AlamatControllerImpl{
		alamatUseCase: alamatUseCase,
	}
}

func (ac *AlamatControllerImpl) CreateAlamat(ctx *fiber.Ctx) (err error) {

	// get user id from middleware
	userIDMiddleware := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))

	// get user input and error checking
	data := new(dto.CreateAlamatRequest)
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call CreateAlamat from alamat useCae to get id new inserted user and error information
	c := ctx.Context()
	IDUseCase, errUseCase := ac.alamatUseCase.CreateAlamat(c, uint(userID), *data)
	// error checking
	if errUseCase.Err != nil {
		var response = BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
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
		Data:    IDUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (ac *AlamatControllerImpl) GetMyAlamat(ctx *fiber.Ctx) (err error) {

	// get user_id from middleware
	userIDString := ctx.Locals("userID")
	userIDInt, _ := strconv.Atoi(fmt.Sprintf("%v", userIDString))
	// get limit and page from query parameter url
	filter := new(dto.FilterAlamat)
	if errQuery := ctx.QueryParser(filter); errQuery != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errQuery.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call GetMyAlamat from alamat useCase to get alamat records and error information
	c := ctx.Context()
	responseUseCase, errUseCase := ac.alamatUseCase.GetMyAlamat(c, uint(userIDInt), *filter)
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

func (ac *AlamatControllerImpl) GetAlamatByID(ctx *fiber.Ctx) (err error) {

	// get id from url parameter
	getParam := ctx.Params("id")
	IDParam, errParam := strconv.Atoi(getParam)
	if errParam != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{errParam.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetAlamatByID from alamat useCase to get alamat data and error information
	c := ctx.Context()
	responseUseCase, errUseCase := ac.alamatUseCase.GetAlamatByID(c, uint(IDParam))
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

func (ac *AlamatControllerImpl) UpdateAlamatByID(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()
	// get userID from middleware
	userIDMiddleware := ctx.Locals("userID")
	userIDInt, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))

	// get id alamat from url parameter
	getParam := ctx.Params("id")
	IDParam, errParam := strconv.Atoi(getParam)
	if errParam != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{errParam.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	data := new(dto.UpdateAlamatRequest)
	// get user input and error checking
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call UpdateAlamatByID from alamat useCase to update alamat record and get error information
	if errUseCase := ac.alamatUseCase.UpdateAlamatByID(c, uint(userIDInt), uint(IDParam), *data); errUseCase.Err != nil {
		var response = BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
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

func (ac *AlamatControllerImpl) DeleteAlamatByID(ctx *fiber.Ctx) (err error) {
	// get userID from middleware
	getUserIDMiddleware := ctx.Locals("userID")
	// convert user id from middleware (interface type) to int
	UserID, _ := strconv.Atoi(fmt.Sprintf("%v", getUserIDMiddleware))

	// get ID from url parameter
	getParams := ctx.Params("id")
	// convert ID from parameters (string type) to int
	IDParams, errParams := strconv.Atoi(getParams)
	if errParams != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{errParams.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call DeleteAlamatByID from alamat useCase to delete alamat record and get error information
	c := ctx.Context()
	if errUseCase := ac.alamatUseCase.DeleteAlamatByID(c, uint(UserID), uint(IDParams)); errUseCase.Err != nil {
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
