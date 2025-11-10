package cacheservice

import (
	"context"
	"fmt"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/utils"
	"time"
)

const (
	prefix = "accountVerification:%s"
	ttl    = 30 * time.Minute
)

func (s *cacheService) FindVerificationToken(ctx context.Context, hashedToken string) (*VerificationTokenPayload, error) {
	key := s.SetVerificationTokenKey(hashedToken)
	dataStr, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	data, err := utils.JsonStringToMap(dataStr)
	if err != nil {
		return nil, err
	}
	return &VerificationTokenPayload{
		Code:   data["code"],
		UserId: data["userId"],
	}, nil
}

func (s *cacheService) SaveVerificationToken(ctx context.Context, hashedToken string, payload VerificationTokenPayload) error {
	key := s.SetVerificationTokenKey(hashedToken)
	value, err := utils.MapToJsonString(map[string]string{
		"code":   payload.Code,
		"userId": payload.UserId,
	})
	if err != nil {
		return err
	}
	params := repositories.SaveParams{
		Key:    key,
		Value:  value,
		Expiry: ttl,
	}
	return s.cacheRepo.Save(ctx, params)
}

func (s *cacheService) DeleteVerificationToken(ctx context.Context, hashedToken string) error {
	key := s.SetVerificationTokenKey(hashedToken)
	return s.cacheRepo.Delete(ctx, key)
}

func (s *cacheService) SetVerificationTokenKey(hashedToken string) string {
	key := fmt.Sprintf(prefix, hashedToken)
	return key
}
