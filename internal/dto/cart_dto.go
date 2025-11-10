package dto

import "github.com/google/uuid"

type AddToCart struct {
	ProductID uuid.UUID `json:"productId" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
}
