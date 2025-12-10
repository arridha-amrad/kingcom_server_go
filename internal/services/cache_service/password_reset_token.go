package cacheservice

import (
	"context"
	"fmt"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/utils"
	"time"
)

const (
	prefixPwdResetToken = "pwdReset:%s"
	ttlPwdReset         = time.Hour * 24
)

func (s *cacheService) FindPasswordResetToken(ctx context.Context, hashedToken string) (*PasswordResetTokenPayload, error) {
	key := s.SetPasswordResetTokenKey(hashedToken)
	jsonStr, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	data, err := utils.JsonStringToMap(jsonStr)
	if err != nil {
		return nil, err
	}
	return &PasswordResetTokenPayload{UserId: data["userId"]}, nil
}

func (s *cacheService) SavePasswordResetToken(ctx context.Context, hashedToken string, payload PasswordResetTokenPayload) error {
	key := s.SetPasswordResetTokenKey(hashedToken)
	value, err := utils.MapToJsonString(map[string]string{
		"userId": payload.UserId,
	})
	if err != nil {
		return err
	}
	params := repositories.SaveParams{
		Key:    key,
		Value:  value,
		Expiry: ttlPwdReset,
	}
	return s.cacheRepo.Save(ctx, params)
}

func (s *cacheService) DeletePasswordResetToken(ctx context.Context, hashedToken string) error {
	key := s.SetPasswordResetTokenKey(hashedToken)
	return s.cacheRepo.Delete(ctx, key)
}

func (s *cacheService) SetPasswordResetTokenKey(hashedToken string) string {
	return fmt.Sprintf(prefixPwdResetToken, hashedToken)
}
