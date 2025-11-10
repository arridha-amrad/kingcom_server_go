package dto

import "github.com/google/uuid"

type Shipping struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Service     string  `json:"service"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	Etd         string  `json:"etd"`
	Address     string  `json:"address"`
}

type CreateOrder struct {
	Total    int64    `json:"total"`
	Items    []Item   `json:"items"`
	Shipping Shipping `json:"shipping"`
}

type Item struct {
	CartID    uuid.UUID `json:"cartId"`
	ProductID uuid.UUID `json:"productId"`
	Quantity  int       `json:"quantity"`
}
