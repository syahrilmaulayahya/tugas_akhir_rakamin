package usecase

import (
	"context"
	"github.com/kataras/jwt"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
)

type Middleware interface {
	CreateJwt(ctx context.Context, claim any) (string, error)
	VerifyJwt(ctx context.Context, token string) (claim dto.ClaimJwt, err error)
}

type Config struct {
	SharedKey string
}

type MiddlewareImpl struct {
	config Config
}

func NewMiddleware(config Config) Middleware {
	return &MiddlewareImpl{config: config}
}

func (m *MiddlewareImpl) CreateJwt(ctx context.Context, claim any) (string, error) {
	jwtKey := m.config.SharedKey
	token, err := jwt.Sign(jwt.HS256, []byte(jwtKey), claim)
	if err != nil {
		return "", err
	}
	return string(token), err
}

func (m *MiddlewareImpl) VerifyJwt(ctx context.Context, token string) (claim dto.ClaimJwt, err error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(m.config.SharedKey), []byte(token))
	if err != nil {
		return claim, err
	}
	err = verifiedToken.Claims(&claim)
	if err != nil {
		return claim, err
	}
	return claim, err
}
