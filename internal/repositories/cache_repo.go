package repositories

import (
	"context"
	"kingcom_api/internal/lib"
	"time"
)

type CacheRepository struct {
	*lib.Redis
	logger *lib.Logger
}

func NewCacheRepository(logger *lib.Logger, rdb *lib.Redis) *CacheRepository {
	return &CacheRepository{logger: logger, Redis: rdb}
}

func (r *CacheRepository) Save(ctx context.Context, params SaveParams) error {
	return r.Redis.Set(ctx, params.Key, params.Value, params.Expiry).Err()
}

func (r *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.Redis.Get(ctx, key).Result()
}

func (r *CacheRepository) Delete(ctx context.Context, key string) error {
	return r.Redis.Del(ctx, key).Err()
}

func (r *CacheRepository) WithLog(err error) error {
	if err != nil {
		r.logger.Errorf("error occurs in cache repository. %v", err)
	}
	return err
}

type VerificationData struct {
	Code   string
	UserId string
}

type SaveParams struct {
	Key    string
	Value  string
	Expiry time.Duration
}
