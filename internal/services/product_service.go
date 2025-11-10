package services

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"kingcom_api/internal/repositories"

	"github.com/google/uuid"
)

type productService struct {
	prodRepo repositories.ProductRepository
	db       *lib.Database
}

type ProductService interface {
	FindMany() (*[]models.Product, error)
	FindBySlug(slug string) (*models.Product, error)
	InsertProduct(product *models.Product) error
	FindById(id uuid.UUID) (*models.Product, error)
}

func NewProductService(
	logger *lib.Logger,
	db *lib.Database,
	prodRepo repositories.ProductRepository,
) ProductService {
	return &productService{
		db:       db,
		prodRepo: prodRepo,
	}
}

func (p *productService) FindMany() (*[]models.Product, error) {
	return p.prodRepo.FindMany()
}

func (p *productService) FindBySlug(slug string) (*models.Product, error) {
	return p.prodRepo.FindBySlug(slug)
}

func (p *productService) InsertProduct(product *models.Product) error {
	return p.prodRepo.InsertProduct(product)
}

func (p *productService) FindById(id uuid.UUID) (*models.Product, error) {
	return p.prodRepo.FindById(id)
}
