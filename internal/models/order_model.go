package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
)

type Order struct {
	ID             uuid.UUID     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	OrderNumber    string        `gorm:"column:order_number;unique;not null" json:"orderNumber"`
	Status         OrderStatus   `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Total          int64         `gorm:"not null" json:"total"`
	PaymentMethod  string        `gorm:"column:payment_method;type:varchar(50)" json:"paymentMethod"`
	BillingAddress string        `gorm:"column:billing_address;type:text" json:"billingAddress"`
	CreatedAt      time.Time     `gorm:"column:created_at;not null" json:"createdAt"`
	PaidAt         time.Time     `gorm:"column:paid_at" json:"paidAt"`
	ShippedAt      time.Time     `gorm:"column:shipped_at" json:"shippedAt"`
	DeliveredAt    time.Time     `gorm:"column:delivered_at" json:"deliveredAt"`
	ShippingID     uint          `gorm:"column:shipping_id;not null" json:"-"`
	Shipping       OrderShipping `gorm:"foreignKey:ShippingID;constraint:onDelete:CASCADE" json:"shipping"`
	UserID         uuid.UUID     `gorm:"column:user_id;type:uuid;not null" json:"-"`
	User           User          `gorm:"foreignKey:UserID" json:"-"`
	OrderItems     []OrderItem   `gorm:"foreignKey:OrderID;constraint:onDelete:CASCADE;" json:"items"`
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	ProductID uuid.UUID `gorm:"column:product_id;not null" json:"-"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	OrderID   uuid.UUID `gorm:"column:order_id;not null;constraint:onDelete:CASCADE;" json:"-"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"-"`
}

type OrderShipping struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Service     string  `json:"service"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	Etd         string  `json:"etd"`
	Address     string  `json:"address"`
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.OrderNumber = generateOrderNumber()
	return
}

func generateOrderNumber() string {
	timestamp := time.Now().UnixNano()
	randomPart := uuid.New().String()[0:8]
	return fmt.Sprintf("ORD-%d-%s", timestamp, randomPart)
}
