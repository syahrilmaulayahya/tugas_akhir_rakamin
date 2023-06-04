package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/utils/apicall"
	"strconv"
)

type ProvinceCityController interface {
	GetAllProvince(ctx *fiber.Ctx) (err error)
	GetListCities(ctx *fiber.Ctx) (err error)
	GetDetailProvince(ctx *fiber.Ctx) (err error)
}

type ProvinceCityControllerImpl struct {
	apiCall apicall.ProvinceCity
}

func NewProvincCityController(apiCall apicall.ProvinceCity) ProvinceCityController {
	return &ProvinceCityControllerImpl{
		apiCall: apiCall,
	}
}

func (pcci *ProvinceCityControllerImpl) GetAllProvince(ctx *fiber.Ctx) (err error) {
	// call GetAllProvince function from apicall package to get all province available and error information
	result, err := pcci.apiCall.GetAllProvince()
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to get data",
		Error:   nil,
		Data:    result,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pcci *ProvinceCityControllerImpl) GetListCities(ctx *fiber.Ctx) (err error) {

	// get id from url parameter
	GetParams := ctx.Params("prov_id")

	// check if id valid
	provinceID, err := strconv.Atoi(GetParams)
	if err != nil && provinceID <= 11 && provinceID >= 94 {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer >= 11 and <= 94",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetAllProvince function from apicall package to get regency with specified provinceId and error information
	result, err := pcci.apiCall.GetListCity(uint(provinceID))
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to get data",
		Error:   nil,
		Data:    result,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pcci *ProvinceCityControllerImpl) GetDetailProvince(ctx *fiber.Ctx) (err error) {

	// get id from url parameter
	GetParams := ctx.Params("prov_id")

	// check if id valid
	provinceID, err := strconv.Atoi(GetParams)
	if err != nil && provinceID <= 11 && provinceID >= 94 {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer >= 11 and <= 94",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetAllProvince function from apicall package to get regency with specified provinceId and error information
	result, err := pcci.apiCall.GetDetailProvince(uint(provinceID))
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to get data",
		Error:   nil,
		Data:    result,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
