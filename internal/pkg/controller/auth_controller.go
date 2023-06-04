package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"strconv"
	"time"
)

type Auth interface {
	CheckJwt(ctx *fiber.Ctx) error
	CheckJwtUser(ctx *fiber.Ctx) error
	CheckJwtAdmin(ctx *fiber.Ctx) error
}

type AuthImpl struct {
	middleware usecase.Middleware
}

func NewAuthImpl(middleware usecase.Middleware) Auth {
	return &AuthImpl{middleware: middleware}
}

func (a *AuthImpl) CheckJwt(ctx *fiber.Ctx) error {

	// get token from header request
	token := ctx.Get("token")

	// verify jwt with VerifyJWt function from middleware useCase
	c := ctx.Context()
	claim, err := a.middleware.VerifyJwt(c, token)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "invalid token",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// check jwt content
	switch {

	case claim.Issuer != "syahril":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Type != "ACCESS_TOKEN":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case !time.Unix(claim.NotValidBefore, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case time.Unix(claim.ExpiredAT, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	}

	// send userID to next controller
	userID := strconv.Itoa(int(claim.Subject))
	ctx.Locals("userID", userID)
	if err = ctx.Next(); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "context error",
			Error:   []string{"error while setting context value"},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}
	return err
}

func (a *AuthImpl) CheckJwtUser(ctx *fiber.Ctx) error {

	// get token from header request
	token := ctx.Get("token")

	// verify jwt with VerifyJWt function from middleware useCase
	c := ctx.Context()
	claim, err := a.middleware.VerifyJwt(c, token)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "invalid token",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// check jwt content
	switch {
	case claim.Issuer != "syahril":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Audience != "user":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Scope != "user":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Type != "ACCESS_TOKEN":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case !time.Unix(claim.NotValidBefore, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp("Invalid Valid Time Token"))
	case time.Unix(claim.ExpiredAT, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp("Invalid Expire Time Token"))
	}

	// send userID to next controller
	userID := strconv.Itoa(int(claim.Subject))
	ctx.Locals("userID", userID)
	if err = ctx.Next(); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "context error",
			Error:   []string{"error while setting context value"},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}
	return err
}

func (a *AuthImpl) CheckJwtAdmin(ctx *fiber.Ctx) error {

	// get token from header request
	token := ctx.Get("token")

	// verify jwt with VerifyJWt function from middleware useCase
	c := ctx.Context()
	claim, err := a.middleware.VerifyJwt(c, token)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "invalid token",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// check jwt content
	switch {
	case claim.Issuer != "syahril":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Audience != "admin":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Scope != "admin":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case claim.Type != "ACCESS_TOKEN":
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case !time.Unix(claim.NotValidBefore, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	case time.Unix(claim.ExpiredAT, 0).Before(time.Now()):
		response := BaseResponse{}
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.TokenErrorResp(ctx.Method()))
	}

	// send userID to next controller
	userID := strconv.Itoa(int(claim.Subject))
	ctx.Locals("userID", userID)
	if err = ctx.Next(); err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "context error",
			Error:   []string{"error while setting context value"},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)

	}
	return err
}
