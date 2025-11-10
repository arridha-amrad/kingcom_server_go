package repositories

import (
	"errors"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	*lib.Database
	logger *lib.Logger
}

type ProductRepository interface {
	InsertProduct(product *models.Product) error
	FindMany() (*[]models.Product, error)
	FindById(id uuid.UUID) (*models.Product, error)
	FindBySlug(slug string) (*models.Product, error)
}

func NewProductRepository(
	db *lib.Database,
	logger *lib.Logger,
) ProductRepository {
	return &productRepository{
		Database: db,
		logger:   logger,
	}
}

func (p *productRepository) InsertProduct(product *models.Product) error {
	return p.Create(product).Error
}

func (p productRepository) FindMany() (*[]models.Product, error) {
	var products []models.Product
	if err := p.Preload("Images").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (p productRepository) FindById(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := p.DB.Create(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) FindBySlug(slug string) (*models.Product, error) {
	var product models.Product
	if err := p.DB.Where("slug = ?", slug).
		Preload("Images").
		First(&product).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
