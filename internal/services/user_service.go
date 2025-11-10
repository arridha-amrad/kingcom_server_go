package services

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"kingcom_api/internal/repositories"

	"github.com/google/uuid"
)

type userService struct {
	userRepo repositories.UserRepository
	db       *lib.Database
}

type UserService interface {
	Create(user *models.User) error
	FindById(userId uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmailOrUsername(identity string) (*models.User, error)
	VerifyUser(userId uuid.UUID) error
	UpdatePassword(id uuid.UUID, pwd string, jwtVersion string) error
}

func NewUserService(
	logger *lib.Logger,
	userRepo repositories.UserRepository,
	db *lib.Database,
) UserService {
	return &userService{
		userRepo: userRepo,
		db:       db,
	}
}

func (s *userService) Create(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) FindById(userId uuid.UUID) (*models.User, error) {
	return s.userRepo.FindById(userId)
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *userService) FindByUsername(username string) (*models.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *userService) FindByEmailOrUsername(identity string) (*models.User, error) {
	return s.userRepo.FindByEmailOrUsername(identity)
}

func (s *userService) VerifyUser(userId uuid.UUID) error {
	return s.userRepo.VerifyUser(userId)
}

func (s *userService) UpdatePassword(id uuid.UUID, pwd string, jwtVersion string) error {
	return s.userRepo.UpdatePassword(id, pwd, jwtVersion)
}
