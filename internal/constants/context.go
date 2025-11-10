package constants

import (
	"time"
)

const (
	// DBTransaction is database transaction handle set at router context
	ACCESS_TOKEN_PAYLOAD      = "accessTokenPayload"
	COOKIE_REFRESH_TOKEN      = "kingcom-refresh-token"
	JWT_VERSION_LENGTH        = 8
	REFRESH_TOKEN_MAX_AGE     = 3600 * 24 * 7
	RAJA_ONGKIR_PROVINCES_KEY = "rajaOngkir:provinces"
)

var TTL = struct {
	AccessToken time.Duration
}{
	AccessToken: 1 * time.Hour,
}
