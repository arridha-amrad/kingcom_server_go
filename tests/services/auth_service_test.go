package services_test

import (
	"context"
	"testing"

	"kingcom_api/internal/lib"
	authservice "kingcom_api/internal/services/authService"
	cacheservice "kingcom_api/internal/services/cache_service"
	"kingcom_api/internal/utils"
	"kingcom_api/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AuthServiceSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockCache *mocks.MockCacheService
	mockRepo  *mocks.MockUserRepository
	env       *lib.Env
	db        *lib.Database
	service   authservice.AuthService
	ctx       context.Context
}

func (s *AuthServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockCache = mocks.NewMockCacheService(s.ctrl)
	s.mockRepo = mocks.NewMockUserRepository(s.ctrl)
	s.env = &lib.Env{
		JwtSecret: "secret",
	}
	s.db = &lib.Database{}
	s.service = authservice.New(s.db, s.env, s.mockCache, s.mockRepo)
	s.ctx = context.Background()
}

func (s *AuthServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

// --- TEST CASES ---

func (s *AuthServiceSuite) TestCreatePwdResetToken() {
	userId := "123"

	s.mockCache.EXPECT().
		SavePasswordResetToken(
			s.ctx,
			gomock.Any(),
			gomock.AssignableToTypeOf(cacheservice.PasswordResetTokenPayload{}),
		).
		DoAndReturn(func(_ context.Context, _ string, payload cacheservice.PasswordResetTokenPayload) error {
			s.Equal(userId, payload.UserId)
			return nil
		})

	token, err := s.service.CreatePwdResetToken(s.ctx, userId)
	s.NoError(err)
	s.NotEmpty(token)
}

func (s *AuthServiceSuite) TestCreateAuthTokens() {
	userId := "123"
	jwtVer := "v1"

	s.mockCache.EXPECT().SaveRefreshToken(s.ctx, gomock.Any(), gomock.Any()).Return(nil)
	s.mockCache.EXPECT().SaveAccessToken(s.ctx, gomock.Any(), gomock.Any()).Return(nil)

	tokens, err := s.service.CreateAuthTokens(s.ctx, userId, jwtVer)
	s.Nil(err)
	s.NotNil(tokens)
}

func (s *AuthServiceSuite) TestDeleteAuthTokens() {
	refToken := "some-refresh-token"
	jti := "jti-uuid"

	s.mockCache.EXPECT().
		DeleteAccessToken(s.ctx, jti).
		Return(nil)

	s.mockCache.EXPECT().
		DeleteRefreshToken(s.ctx, utils.HashWithSHA256(refToken)).
		Return(nil)

	err := s.service.DeleteAuthTokens(s.ctx, refToken, jti)
	s.NoError(err)
}

func (s *AuthServiceSuite) TestCreateAndStoreRefToken() {
	userId := uuid.New().String()
	jti := uuid.New().String()
	s.mockCache.EXPECT().
		SaveRefreshToken(
			s.ctx,
			gomock.Any(),
			gomock.AssignableToTypeOf(cacheservice.RefreshTokenPayload{}),
		).
		DoAndReturn(func(_ context.Context, _ string, payload cacheservice.RefreshTokenPayload) error {
			s.Equal(userId, payload.UserId)
			s.Equal(jti, payload.Jti)
			return nil
		})

	token, err := s.service.CreateAndStoreRefToken(s.ctx, userId, jti)
	s.NoError(err)
	s.NotEmpty(token)
}

func (s *AuthServiceSuite) TestCreateVerificationToken() {
	userId := "user123"

	s.mockCache.EXPECT().
		SaveVerificationToken(s.ctx, gomock.Any(), gomock.Any()).
		Return(nil)

	payload, token, err := s.service.CreateVerificationToken(s.ctx, userId)
	s.NoError(err)
	s.NotNil(payload)
	s.NotEmpty(token)
	s.Equal(userId, payload.UserId)
}

func (s *AuthServiceSuite) TestCreateJWT() {
	payload := authservice.JWTPayload{
		UserId:     "user123",
		Jti:        "uuid-1",
		JwtVersion: "v1",
	}

	token, err := s.service.CreateJWT(payload, s.env.JwtSecret, "issuer")
	s.NoError(err)
	s.NotEmpty(token)
}

func (s *AuthServiceSuite) TestVerifyJwt() {
	token := "invalid.jwt.token"
	_, err := s.service.VerifyJwt(token)
	s.Error(err)
}

func (s *AuthServiceSuite) TestHashAndVerifyPassword() {
	hashed, err := s.service.HashPassword("password123")
	s.NoError(err)
	s.NotEmpty(hashed)

	err = s.service.VerifyPassword(hashed, "password123")
	s.NoError(err)

	err = s.service.VerifyPassword(hashed, "wrongpass")
	s.Error(err)
}

func (s *AuthServiceSuite) TestCreateAndStoreAccessToken() {
	s.mockCache.EXPECT().
		SaveAccessToken(s.ctx, gomock.Any(), gomock.Any()).
		Return(nil)

	token, err := s.service.CreateAndStoreAccessToken(s.ctx, "jti-1", "user123", "v1")
	s.NoError(err)
	s.NotEmpty(token)
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceSuite))
}
