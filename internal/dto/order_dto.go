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
	CartID            uuid.UUID `json:"cartId"`
	ProductID         uuid.UUID `json:"productId"`
	Quantity          int       `json:"quantity"`
	PriceAtOrder      float64   `json:"priceAtOrder"`
	DiscountAtOrder   int       `json:"discountAtOrder"`
	FinalPriceAtOrder float64   `json:"finalPriceAtOrder"`
}

type CreateTransactionTokenParam struct {
	OrderId string `json:"orderId" validate:"required"`
}

type MidtransNotificationParams struct {
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	SignatureKey      string `json:"signature_key"`
	SettlementTime    string `json:"settlement_time"`
}
