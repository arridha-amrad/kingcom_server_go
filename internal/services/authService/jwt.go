package authservice

import (
	"errors"
	"kingcom_api/internal/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *authService) CreateJWT(payload JWTPayload, secret, issuer string) (string, error) {
	claims := CustomClaims{
		UserID:     payload.UserId,
		JTI:        payload.Jti,
		JwtVersion: payload.JwtVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.TTL.AccessToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *authService) VerifyJwt(token string) (*JWTPayload, error) {
	claims := &CustomClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.env.JwtSecret), nil
	})

	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid token")
	}

	return &JWTPayload{
		UserId:     claims.UserID,
		Jti:        claims.JTI,
		JwtVersion: claims.JwtVersion,
	}, nil

}
