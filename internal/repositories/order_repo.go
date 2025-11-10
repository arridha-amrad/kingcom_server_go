package repositories

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	*lib.Database
	logger *lib.Logger
}

type OrderRepository interface {
	Create(order *models.Order) error
	WithTrx(tx *gorm.DB) OrderRepository
	FindByUserId(userId uuid.UUID) (*[]models.Order, error)
}

func NewOrderRepository(
	db *lib.Database,
	logger *lib.Logger,
) OrderRepository {
	return &orderRepository{
		Database: db,
		logger:   logger,
	}
}

func (o *orderRepository) Create(order *models.Order) error {
	return o.DB.Create(order).Error
}

func (o *orderRepository) WithTrx(tx *gorm.DB) OrderRepository {
	if tx == nil {
		o.logger.Error("transaction db for order repo is missing")
		return o
	}
	return &orderRepository{
		Database: &lib.Database{DB: tx},
		logger:   o.logger,
	}
}

func (o *orderRepository) FindByUserId(userId uuid.UUID) (*[]models.Order, error) {
	var orders []models.Order
	if err := o.DB.
		Where("user_id = ?", userId).
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Product.Images").
		Preload("Shipping").
		Find(&orders).
		Error; err != nil {
		return nil, nil
	}
	return &orders, nil
}
