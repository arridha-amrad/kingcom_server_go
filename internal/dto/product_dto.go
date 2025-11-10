package dto

type CreateProduct struct {
	Name          string   `json:"name" validate:"required"`
	Weight        float64  `json:"weight" validate:"required,numeric,gt=0"`
	Price         float64  `json:"price" validate:"required,numeric,gt=0"`
	Description   string   `json:"description" validate:"required"`
	Stock         uint     `json:"stock" validate:"required,numeric,gte=0"`
	Discount      int      `json:"discount" validate:"omitempty,numeric,gte=0"`
	Specification *string  `json:"specification" validate:"omitempty"`
	VideoUrl      *string  `json:"videoUrl" validate:"omitempty,url"`
	Images        []string `json:"images" validate:"required,min=1,dive,required,url"`
}
