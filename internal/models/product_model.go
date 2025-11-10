package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID      `gorm:"column:id" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
	Name          string         `gorm:"column:name" json:"name"`
	Weight        float64        `gorm:"column:weight" json:"weight"`
	Slug          string         `gorm:"column:slug" json:"slug"`
	Price         float64        `gorm:"column:price" json:"price"`
	Description   string         `gorm:"column:description" json:"description"`
	Specification *string        `gorm:"column:specification" json:"specification"`
	Stock         uint           `gorm:"column:stock" json:"stock"`
	VideoUrl      *string        `gorm:"column:video_url" json:"videoUrl"`
	Discount      int            `gorm:"column:discount" json:"discount"`
	Images        []ProductImage `json:"images"`
}

type ProductImage struct {
	ID        uint      `gorm:"column:id" json:"id"`
	Url       string    `gorm:"url" json:"url"`
	ProductID uuid.UUID `gorm:"column:product_id" json:"-"`
	Product   Product   `gorm:"foreignKey:ProductID;not null;constraint:onDelete:CASCADE" json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
