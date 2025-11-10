package services_test

import (
	"errors"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"kingcom_api/internal/services"
	"kingcom_api/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockRepo *mocks.MockUserRepository
	service  services.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mocks.NewMockUserRepository(s.ctrl)
	s.service = services.NewUserService(&lib.Logger{}, s.mockRepo, &lib.Database{})
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

// --- TESTS ---

func (s *UserServiceTestSuite) TestCreate() {
	user := &models.User{ID: uuid.New(), Username: "ari"}

	s.Run("success", func() {
		s.mockRepo.EXPECT().Create(user).Return(nil)
		err := s.service.Create(user)
		s.NoError(err)
	})

	s.Run("error", func() {
		s.mockRepo.EXPECT().Create(user).Return(errors.New("db error"))
		err := s.service.Create(user)
		s.EqualError(err, "db error")
	})
}

func (s *UserServiceTestSuite) TestFindById() {
	id := uuid.New()
	expected := &models.User{ID: id}

	s.mockRepo.EXPECT().FindById(id).Return(expected, nil)
	user, err := s.service.FindById(id)

	s.NoError(err)
	s.Equal(expected, user)
}

func (s *UserServiceTestSuite) TestFindByEmail() {
	email := "ari@example.com"
	expected := &models.User{Email: email}

	s.mockRepo.EXPECT().FindByEmail(email).Return(expected, nil)
	user, err := s.service.FindByEmail(email)

	s.NoError(err)
	s.Equal(expected, user)
}

func (s *UserServiceTestSuite) TestFindByUsername() {
	username := "ari"
	expected := &models.User{Username: username}

	s.mockRepo.EXPECT().FindByUsername(username).Return(expected, nil)
	user, err := s.service.FindByUsername(username)

	s.NoError(err)
	s.Equal(expected, user)
}

func (s *UserServiceTestSuite) TestFindByEmailOrUsername() {
	identity := "ari"
	expected := &models.User{Username: "ari"}

	s.mockRepo.EXPECT().FindByEmailOrUsername(identity).Return(expected, nil)
	user, err := s.service.FindByEmailOrUsername(identity)

	s.NoError(err)
	s.Equal(expected, user)
}

func (s *UserServiceTestSuite) TestVerifyUser() {
	id := uuid.New()

	s.Run("success", func() {
		s.mockRepo.EXPECT().VerifyUser(id).Return(nil)
		err := s.service.VerifyUser(id)
		s.NoError(err)
	})

	s.Run("error", func() {
		s.mockRepo.EXPECT().VerifyUser(id).Return(errors.New("verify failed"))
		err := s.service.VerifyUser(id)
		s.EqualError(err, "verify failed")
	})
}

func (s *UserServiceTestSuite) TestUpdatePassword() {
	id := uuid.New()
	pwd := "hashed-password"
	jwtVer := "v1"

	s.Run("success", func() {
		s.mockRepo.EXPECT().UpdatePassword(id, pwd, jwtVer).Return(nil)
		err := s.service.UpdatePassword(id, pwd, jwtVer)
		s.NoError(err)
	})

	s.Run("error", func() {
		s.mockRepo.EXPECT().UpdatePassword(id, pwd, jwtVer).Return(errors.New("update failed"))
		err := s.service.UpdatePassword(id, pwd, jwtVer)
		s.EqualError(err, "update failed")
	})
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
