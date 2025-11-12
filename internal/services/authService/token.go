package authservice

import (
	"context"
	"errors"
	"fmt"
	"kingcom_api/internal/models"
	cacheservice "kingcom_api/internal/services/cache_service"
	"kingcom_api/internal/utils"

	"github.com/google/uuid"
)

func (s *authService) CreatePwdResetToken(ctx context.Context, userId string) (string, error) {
	tokenPair, err := createPairToken()
	if err != nil {
		return "", err
	}
	if err := s.cacheSvc.SavePasswordResetToken(
		ctx,
		tokenPair.Hashed,
		cacheservice.PasswordResetTokenPayload{UserId: userId},
	); err != nil {
		return "", err
	}
	return tokenPair.Raw, nil
}

func (s *authService) CreateAuthTokens(ctx context.Context, userId, jwtVersion, role string) (*AuthTokens, error) {
	jti := uuid.New().String()
	refresh, err := s.CreateAndStoreRefToken(ctx, userId, jti)
	if err != nil {
		return nil, err
	}
	access, err := s.CreateAndStoreAccessToken(ctx, jti, userId, jwtVersion, role)
	if err != nil {
		return nil, err
	}
	return &AuthTokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *authService) DeleteAuthTokens(ctx context.Context, refToken, jti string) error {
	if err := s.cacheSvc.DeleteAccessToken(ctx, jti); err != nil {
		return err
	}
	return s.cacheSvc.DeleteRefreshToken(ctx, utils.HashWithSHA256(refToken))
}

func (s *authService) CreateAndStoreRefToken(ctx context.Context, userId, jti string) (string, error) {
	pair, err := createPairToken()
	if err != nil {
		return "", err
	}
	payload := cacheservice.RefreshTokenPayload{UserId: userId, Jti: jti}
	if err := s.cacheSvc.SaveRefreshToken(ctx, pair.Hashed, payload); err != nil {
		return "", fmt.Errorf("save refresh token: %w", err)
	}
	return pair.Raw, nil
}

func (s *authService) CreateAndStoreAccessToken(ctx context.Context, jti, userId, jwtVersion, role string) (string, error) {
	jwtPayload := JWTPayload{UserId: userId, JwtVersion: jwtVersion, Jti: jti, Role: models.Role(role)}
	token, err := s.CreateJWT(jwtPayload, s.env.JwtSecret, s.env.AppTitle)
	if err != nil {
		return "", err
	}
	if err := s.cacheSvc.SaveAccessToken(
		ctx,
		jti,
		cacheservice.AccessTokenPayload{
			UserId:     userId,
			JwtVersion: jwtVersion,
		},
	); err != nil {
		return "", fmt.Errorf("save access token: %w", err)
	}
	return token, nil
}

func (s *authService) CreateVerificationToken(ctx context.Context, userId string) (*cacheservice.VerificationTokenPayload, string, error) {
	pairToken, err := createPairToken()
	if err != nil {
		return nil, "", err
	}
	code, err := utils.GenerateRandomBytes(4)
	if err != nil {
		return nil, "", err
	}
	if err := s.cacheSvc.SaveVerificationToken(
		ctx,
		pairToken.Hashed,
		cacheservice.VerificationTokenPayload{
			Code:   code,
			UserId: userId,
		}); err != nil {
		return nil, "", err
	}
	return &cacheservice.VerificationTokenPayload{
		UserId: userId,
		Code:   code,
	}, pairToken.Raw, nil
}

// Stateless helpers
func createPairToken() (*TokenPair, error) {
	rawToken, err := utils.GenerateRandomBytes(32)
	if err != nil {
		return nil, errors.New("failure on generating random bytes")
	}
	hashedToken := utils.HashWithSHA256(rawToken)
	return &TokenPair{
		Raw:    rawToken,
		Hashed: hashedToken,
	}, nil
}
