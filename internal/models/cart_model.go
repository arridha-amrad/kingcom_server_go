package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Quantity  int       `gorm:"not null;check:quantity > 0" json:"quantity"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;not null;uniqueIndex:idx_user_product" json:"-"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	ProductID uuid.UUID `gorm:"column:product_id;type:uuid;not null;uniqueIndex:idx_user_product" json:"-"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product"`
}

func (p *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
