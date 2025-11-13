package services

import (
	"kingcom_api/internal/models"
	"kingcom_api/internal/repositories"

	"github.com/google/uuid"
)

type cartService struct {
	cartRepo repositories.CartRepository
}

type CartService interface {
	Add(cart *models.Cart) error
	FindWithProduct(userId uuid.UUID) (*[]models.Cart, error)
	Delete(id uuid.UUID) error
	FindOne(id uuid.UUID) (*models.Cart, error)
}

func NewCartService(
	cartRepo repositories.CartRepository,
) CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

func (c *cartService) Add(cart *models.Cart) error {
	return c.cartRepo.Add(cart)
}

func (c *cartService) FindWithProduct(userId uuid.UUID) (*[]models.Cart, error) {
	return c.cartRepo.FindWithProduct(userId)
}

func (c *cartService) Delete(id uuid.UUID) error {
	return c.cartRepo.Delete(id)
}

func (c *cartService) FindOne(id uuid.UUID) (*models.Cart, error) {
	return c.cartRepo.FindById(id)
}
