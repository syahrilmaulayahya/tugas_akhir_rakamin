package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"strconv"
)

type TRXController interface {
	GetALlTRX(ctx *fiber.Ctx) (err error)
	GetTRXByID(ctx *fiber.Ctx) (err error)
	CreateTRX(ctx *fiber.Ctx) (err error)
}

type TRXControllerImpl struct {
	trxUseCase usecase.TRXUseCase
}

func NewTRXController(trxUseCase usecase.TRXUseCase) TRXController {
	return &TRXControllerImpl{trxUseCase: trxUseCase}
}
func (trxc *TRXControllerImpl) GetALlTRX(ctx *fiber.Ctx) (err error) {
	// get userID from middleware
	userIDMiddleware := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))
	// get limit and page from query parameter url
	params := new(dto.FilterTRX)
	if errQuery := ctx.QueryParser(params); errQuery != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errQuery.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call GetTRXByID from trx useCase
	c := ctx.Context()
	trxUsecase, errUseCase := trxc.trxUseCase.GetAllTRX(c, uint(userID), *params)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	type listTRXResponse struct {
		Data []dto.TRXGetResponse `json:"data"`
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    listTRXResponse{Data: trxUsecase},
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (trxc *TRXControllerImpl) GetTRXByID(ctx *fiber.Ctx) (err error) {
	// get userID from middleware
	userIDMiddleware := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))

	// get id trx from url parameter
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
	// call GetTRXByID from trx useCase
	c := ctx.Context()
	trxUsecase, errUseCase := trxc.trxUseCase.GetTRXByID(c, uint(userID), uint(IDParam))
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
		Data:    trxUsecase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (trxc *TRXControllerImpl) CreateTRX(ctx *fiber.Ctx) (err error) {
	// get user id from middleware
	userIDMiddleware := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))

	data := new(dto.TRX)
	// get user input
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed To POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call CreateTRX from TRXUseCase
	c := ctx.Context()
	data.UserID = uint(userID)
	IDUsecase, errUsecase := trxc.trxUseCase.CreateTRX(c, *data)
	// error checking
	if errUsecase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{errUsecase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUsecase.Code).JSON(response)
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    IDUsecase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
