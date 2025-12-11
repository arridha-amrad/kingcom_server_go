package services

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"
	"kingcom_api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderService struct {
	orderRepo repositories.OrderRepository
	cartRepo  repositories.CartRepository
	*lib.Database
}

type OrderService interface {
	PlaceOrder(order *models.Order, cartsIds []uuid.UUID) error
	FindByUserId(userId uuid.UUID) (*[]models.Order, error)
	FindById(id uuid.UUID) (*models.Order, error)
	Update(order *models.Order) error
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	cartRepo repositories.CartRepository,
	db *lib.Database,
) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		Database:  db,
		cartRepo:  cartRepo,
	}
}

func (o *orderService) Update(order *models.Order) error {
	return o.orderRepo.Update(order)
}

func (o *orderService) FindById(id uuid.UUID) (*models.Order, error) {
	return o.orderRepo.FindById(id)
}

func (o *orderService) PlaceOrder(order *models.Order, cartsIds []uuid.UUID) error {
	err := o.DB.Transaction(func(tx *gorm.DB) error {
		orderRepoTrx := o.orderRepo.WithTrx(tx)
		cartRepoTrx := o.cartRepo.WithTrx(tx)
		if err := orderRepoTrx.Create(order); err != nil {
			return err
		}
		if err := cartRepoTrx.DeleteMany(cartsIds); err != nil {
			return err
		}
		return nil
	})
	return err

}

func (o *orderService) FindByUserId(userId uuid.UUID) (*[]models.Order, error) {
	return o.orderRepo.FindByUserId(userId)
}
