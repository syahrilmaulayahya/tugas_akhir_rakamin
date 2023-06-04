package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/utils"
	"strconv"
	"sync"
	"time"
)

type UserController interface {
	Register(ctx *fiber.Ctx) (err error)
	Login(ctx *fiber.Ctx) (err error)
	GetMyProfile(ctx *fiber.Ctx) (err error)
	UpdateProfile(ctx *fiber.Ctx) (err error)
}

type UserControllerImpl struct {
	userUseCase usecase.UserUseCase
	middleware  usecase.Middleware
}

func NewUserController(userUseCase usecase.UserUseCase, middleware usecase.Middleware) UserController {
	return &UserControllerImpl{
		userUseCase: userUseCase,
		middleware:  middleware,
	}
}

func (uc *UserControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	data := new(dto.UserRegisterAndUpdate)

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

	// call register function from user useCase and insert the data
	if errUseCase := uc.userUseCase.Register(c, *data); errUseCase.Err != nil {
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
		Data:    "Register Succeed",
	}
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (uc *UserControllerImpl) Login(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	data := new(dto.UserLogin)

	// get user input
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call login function from user useCase and insert the data
	responseUseCase, errUseCase := uc.userUseCase.Login(c, *data)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}

	// insert data for creating jwt token
	claimAccess := dto.ClaimJwt{
		JwtID:          uuid.New(),
		Subject:        responseUseCase.ID,
		Issuer:         "syahril",
		Audience:       "user",
		Scope:          "user",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       time.Now().Unix(),
		NotValidBefore: time.Now().Unix(),
		ExpiredAT:      time.Now().Add(24 * time.Hour).Unix(),
	}
	if responseUseCase.IsAdmin {
		claimAccess = dto.ClaimJwt{
			JwtID:          uuid.New(),
			Subject:        responseUseCase.ID,
			Issuer:         "syahril",
			Audience:       "admin",
			Scope:          "admin",
			Type:           "ACCESS_TOKEN",
			IssuedAt:       time.Now().Unix(),
			NotValidBefore: time.Now().Unix(),
			ExpiredAT:      time.Now().Add(24 * time.Hour).Unix(),
		}
	}
	// create jwt token
	accessToken, errToken := uc.middleware.CreateJwt(c, claimAccess)
	if errToken != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Error While Creating Token",
			Error:   []string{errToken.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}

	// give token to user
	responseUseCase.Token = accessToken

	// fetching data from 3rd party api
	var wg sync.WaitGroup
	// get province data from 3rd party api
	var getProvince dto.Province
	var errGetProvince error
	wg.Add(1)
	go func() {
		defer wg.Done()
		getProvince, errGetProvince = utils.GetProvince(responseUseCase.IDProvinsi)
	}()
	// get regency data from 3rd party api
	var getRegency dto.Regency
	var errGetRegency error
	wg.Add(1)
	go func() {
		defer wg.Done()
		getRegency, errGetRegency = utils.GetRegency(responseUseCase.IDKota)

	}()
	wg.Wait()

	// catch error from GetProvince api call
	if errGetProvince != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Error while fetching province data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}
	responseUseCase.Province = getProvince

	// catch error from GetRegency api call
	if errGetRegency != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Error while fetching regency data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	responseUseCase.Regency = getRegency

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()
	// get user id from middleware
	userIDString := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDString))

	// call GetMyProfile with user id as argument from user useCase to get user data and error information
	responseUseCase, errUseCase := uc.userUseCase.GetMyProfile(c, uint(userID))
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
	// fetch data from 3rd party api
	var wg sync.WaitGroup
	// get province data from 3rd party api
	var getProvince dto.Province
	var errGetProvince error
	wg.Add(1)
	go func() {
		defer wg.Done()
		getProvince, err = utils.GetProvince(responseUseCase.IDProvinsi)

	}()
	// get regency data from 3rd party api
	var getRegency dto.Regency
	var errGetRegency error
	wg.Add(1)
	go func() {
		defer wg.Done()
		getRegency, errGetRegency = utils.GetRegency(responseUseCase.IDKota)

	}()
	wg.Wait()

	// catch error from GetProvince error
	if errGetProvince != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Error while fetching province data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}
	responseUseCase.Province = getProvince

	// catch error from GetRegency api call
	if errGetRegency != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Error while fetching regency data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	responseUseCase.Regency = getRegency

	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserControllerImpl) UpdateProfile(ctx *fiber.Ctx) (err error) {
	c := ctx.Context()

	// get id from middleware token
	userIDMiddleware := ctx.Locals("userID")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDMiddleware))

	data := new(dto.UserRegisterAndUpdate)

	// get user input
	if err = ctx.BodyParser(data); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call UpdateProfile from user useCase to get error information
	if errUseCase := uc.userUseCase.UpdateProfile(c, uint(userID), *data); errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	// success response
	var response = BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    "",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
