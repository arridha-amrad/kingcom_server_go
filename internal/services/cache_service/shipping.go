package cacheservice

import (
	"context"
	"encoding/json"
	"errors"
	"kingcom_api/internal/repositories"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	shippingProvincesKey = "shipping-provinces-key"
	shippingProvinceTTL  = 24 * time.Hour
)

func (s *cacheService) FindShippingProvinces(ctx context.Context) (*SaveProvincesData, error) {
	key := s.SetShippingProvincesKey()
	jsonStr, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var result SaveProvincesData
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *cacheService) SaveShippingProvinces(ctx context.Context, params SaveProvincesData) error {
	key := s.SetShippingProvincesKey()
	value, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return s.cacheRepo.Save(ctx, repositories.SaveParams{
		Key:    key,
		Expiry: shippingProvinceTTL,
		Value:  string(value),
	})
}

func (s *cacheService) SetShippingProvincesKey() string {
	return shippingProvincesKey
}
