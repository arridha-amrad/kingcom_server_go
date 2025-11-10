package cacheservice

type SaveProvincesData struct {
	Provinces []Province
}

type Province struct {
	ID   int
	Name string
}

type AccessTokenPayload struct {
	UserId, JwtVersion string
}

type PasswordResetTokenPayload struct {
	UserId string
}

type RefreshTokenPayload struct {
	UserId, Jti string
}

type VerificationTokenPayload struct {
	Code, UserId string
}
