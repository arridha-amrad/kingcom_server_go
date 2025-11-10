package cacheservice

import (
	"context"
	"kingcom_api/internal/repositories"
)

type cacheService struct {
	cacheRepo *repositories.CacheRepository
}

type CacheService interface {
	FindAccessToken(ctx context.Context, jti string) (*AccessTokenPayload, error)
	SaveAccessToken(ctx context.Context, jti string, payload AccessTokenPayload) error
	DeleteAccessToken(ctx context.Context, jti string) error
	SetAccessTokenKey(jti string) string

	FindVerificationToken(ctx context.Context, hashedToken string) (*VerificationTokenPayload, error)
	SaveVerificationToken(ctx context.Context, hashedToken string, payload VerificationTokenPayload) error
	DeleteVerificationToken(ctx context.Context, hashedToken string) error
	SetVerificationTokenKey(hashedToken string) string

	FindPasswordResetToken(ctx context.Context, hashedToken string) (*PasswordResetTokenPayload, error)
	SavePasswordResetToken(ctx context.Context, hashedToken string, payload PasswordResetTokenPayload) error
	DeletePasswordResetToken(ctx context.Context, hashedToken string) error
	SetPasswordResetTokenKey(hashedToken string) string

	FindRefreshToken(ctx context.Context, hashedToken string) (*RefreshTokenPayload, error)
	SaveRefreshToken(ctx context.Context, hashedToken string, payload RefreshTokenPayload) error
	DeleteRefreshToken(ctx context.Context, hashedToken string) error
	SetRefreshTokenKey(hashedToken string) string

	SaveShippingProvinces(ctx context.Context, params SaveProvincesData) error
	FindShippingProvinces(ctx context.Context) (*SaveProvincesData, error)
	SetShippingProvincesKey() string
}

func New(
	cacheRepo *repositories.CacheRepository,
) CacheService {
	return &cacheService{
		cacheRepo: cacheRepo,
	}
}
