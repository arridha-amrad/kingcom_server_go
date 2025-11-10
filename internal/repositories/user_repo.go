package repositories

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	*lib.Database
	logger *lib.Logger
}

type UserRepository interface {
	Create(user *models.User) error
	FindBy(field string, value any) (*models.User, error)
	FindById(userId uuid.UUID) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByEmailOrUsername(identity string) (*models.User, error)
	UpdatePassword(userId uuid.UUID, password string, jwtVersion string) error
	VerifyUser(userId uuid.UUID) error
}

func NewUserRepository(
	db *lib.Database,
	logger *lib.Logger,
) UserRepository {
	return &userRepository{
		Database: db,
		logger:   logger,
	}
}

func (r *userRepository) Create(user *models.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindBy(field string, value any) (*models.User, error) {
	var user models.User
	err := r.DB.Where(field+" = ?", value).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
}

func (r *userRepository) FindById(userId uuid.UUID) (*models.User, error) {
	return r.FindBy("id", userId)
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	return r.FindBy("username", username)
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	return r.FindBy("email", email)
}

func (r *userRepository) FindByEmailOrUsername(identity string) (*models.User, error) {
	if strings.Contains(identity, "@") {
		return r.FindByEmail(identity)
	}
	return r.FindByUsername(identity)

}

func (r *userRepository) UpdatePassword(userId uuid.UUID, password string, jwtVersion string) error {
	return r.DB.Model(models.User{}).
		Where("id = ?", userId).
		Updates(map[string]any{
			"password":    password,
			"jwt_version": jwtVersion,
		}).Error
}

func (r userRepository) VerifyUser(userId uuid.UUID) error {
	return r.DB.Model(models.User{}).
		Where("id = ?", userId).
		Update("is_verified", true).Error
}
