package cacheservice

import (
	"context"
	"fmt"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/utils"
	"time"
)

const (
	ttlRefToken    = 24 * 7 * time.Hour
	prefixRefToken = "refreshToken:%s"
)

func (s *cacheService) FindRefreshToken(ctx context.Context, hashedToken string) (*RefreshTokenPayload, error) {
	key := s.SetRefreshTokenKey(hashedToken)
	dataStr, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	data, err := utils.JsonStringToMap(dataStr)
	if err != nil {
		return nil, err
	}
	return &RefreshTokenPayload{
		UserId: data["userId"],
		Jti:    data["jti"],
	}, nil
}

func (s *cacheService) SaveRefreshToken(ctx context.Context, hashedToken string, payload RefreshTokenPayload) error {
	key := s.SetRefreshTokenKey(hashedToken)
	value, err := utils.MapToJsonString(map[string]string{
		"userId": payload.UserId,
		"jti":    payload.Jti,
	})
	if err != nil {
		return err
	}
	params := repositories.SaveParams{
		Key:    key,
		Expiry: ttlRefToken,
		Value:  value,
	}
	return s.cacheRepo.Save(ctx, params)
}

func (s *cacheService) DeleteRefreshToken(ctx context.Context, hashedToken string) error {
	key := s.SetRefreshTokenKey(hashedToken)
	return s.cacheRepo.Delete(ctx, key)
}

func (s *cacheService) SetRefreshTokenKey(hashedToken string) string {
	return fmt.Sprintf(prefixRefToken, hashedToken)
}
