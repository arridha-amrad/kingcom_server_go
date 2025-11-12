package authservice

import (
	"context"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/repositories"
	cacheservice "kingcom_api/internal/services/cache_service"
)

type authService struct {
	env      *lib.Env
	db       *lib.Database
	cacheSvc cacheservice.CacheService
	userRepo repositories.UserRepository
}

type AuthService interface {
	CreatePwdResetToken(ctx context.Context, userId string) (string, error)
	CreateAuthTokens(ctx context.Context, userId, jwtVersion, role string) (*AuthTokens, error)
	DeleteAuthTokens(ctx context.Context, refToken, jti string) error
	CreateAndStoreRefToken(ctx context.Context, userId, jti string) (string, error)
	CreateVerificationToken(ctx context.Context, userId string) (*cacheservice.VerificationTokenPayload, string, error)
	CreateJWT(payload JWTPayload, secret, issuer string) (string, error)
	VerifyJwt(token string) (*JWTPayload, error)
	HashPassword(plainPassword string) (string, error)
	VerifyPassword(hashedPassword string, plainPassword string) error
	CreateAndStoreAccessToken(ctx context.Context, jti, userId, jwtVersion, role string) (string, error)
}

func New(
	db *lib.Database,
	env *lib.Env,
	cache cacheservice.CacheService,
	repo repositories.UserRepository,
) AuthService {
	return &authService{
		env:      env,
		cacheSvc: cache,
		db:       db,
		userRepo: repo,
	}
}
