package cacheservice

import (
	"context"
	"fmt"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/utils"
	"time"
)

const (
	prefixAccToken = "accessToken:%s"
	ttlAccToken    = time.Hour
)

func (s *cacheService) FindAccessToken(ctx context.Context, jti string) (*AccessTokenPayload, error) {
	key := s.SetAccessTokenKey(jti)
	dataStr, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	data, err := utils.JsonStringToMap(dataStr)
	if err != nil {
		return nil, err
	}
	return &AccessTokenPayload{
		UserId:     data["userId"],
		JwtVersion: data["jwtVersion"],
	}, nil
}

func (s *cacheService) SaveAccessToken(ctx context.Context, jti string, payload AccessTokenPayload) error {
	key := s.SetAccessTokenKey(jti)
	value, err := utils.MapToJsonString(map[string]string{
		"userId":     payload.UserId,
		"jwtVersion": payload.JwtVersion,
	})
	if err != nil {
		return err
	}
	params := repositories.SaveParams{
		Key:    key,
		Expiry: ttlAccToken,
		Value:  value,
	}
	return s.cacheRepo.Save(ctx, params)
}

func (s *cacheService) DeleteAccessToken(ctx context.Context, jti string) error {
	key := s.SetAccessTokenKey(jti)
	return s.cacheRepo.Delete(ctx, key)
}

func (s *cacheService) SetAccessTokenKey(jti string) string {
	return fmt.Sprintf(prefixAccToken, jti)
}
