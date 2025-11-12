package authservice

import (
	"kingcom_api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	Hashed string
	Raw    string
}

type CustomClaims struct {
	JWTPayload
	jwt.RegisteredClaims
}

type RefTokenPayload struct {
	HashedToken string
	RawToken    string
	UserId      string
	Jti         string
}

type AuthTokens struct {
	RefreshToken string
	AccessToken  string
}

type VerificationTokenPayload struct {
	UserId string
	Code   string
}

type JWTPayload struct {
	UserId     string
	Jti        string
	JwtVersion string
	Role       models.Role
}
