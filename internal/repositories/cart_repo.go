package repositories

import (
	"errors"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartRepository struct {
	*lib.Database
	logger *lib.Logger
}

type CartRepository interface {
	WithTrx(tx *gorm.DB) CartRepository
	Add(cart *models.Cart) error
	FindWithProduct(userId uuid.UUID) (*[]models.Cart, error)
	DeleteMany(ids []uuid.UUID) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*models.Cart, error)
}

func NewCartRepository(
	db *lib.Database,
	logger *lib.Logger,
) CartRepository {
	return initRepo(db, logger)
}

func initRepo(db *lib.Database, logger *lib.Logger) CartRepository {
	return &cartRepository{
		Database: db,
		logger:   logger,
	}
}

func (c *cartRepository) FindById(id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	if err := c.DB.Where("id = ?", id).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepository) Delete(id uuid.UUID) error {
	return c.DB.Where("id = ?", id).Delete(&models.Cart{}).Error
}

func (c *cartRepository) DeleteMany(ids []uuid.UUID) error {
	if len(ids) == 0 {
		err := errors.New("ids is empty")
		c.logger.Error(err.Error())
		return err
	}
	return c.DB.Where("id IN ?", ids).Delete(&models.Cart{}).Error
}

func (c *cartRepository) WithTrx(tx *gorm.DB) CartRepository {
	if tx == nil {
		c.logger.Error("transaction db is missing")
		return c
	}
	return initRepo(&lib.Database{DB: tx}, c.logger)
}

func (c *cartRepository) Add(cart *models.Cart) error {
	return c.DB.Create(cart).Error
}

func (c *cartRepository) FindWithProduct(userId uuid.UUID) (*[]models.Cart, error) {
	var carts []models.Cart
	if err := c.DB.
		Where("user_id = ?", userId).
		Preload("Product").
		Preload("Product.Images").
		Find(&carts).Error; err != nil {
		return nil, err
	}
	return &carts, nil
}
