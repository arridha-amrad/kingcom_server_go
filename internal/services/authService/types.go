package authservice

import "github.com/golang-jwt/jwt/v5"

type TokenPair struct {
	Hashed string
	Raw    string
}

type CustomClaims struct {
	UserID     string `json:"userId"`
	JTI        string `json:"jti"`
	JwtVersion string `json:"jwtVersion"`
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
}
