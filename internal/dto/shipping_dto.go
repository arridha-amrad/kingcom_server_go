package dto

type CalcCost struct {
	OriginID      int     `json:"originId" validate:"required,numeric,gt=0"`
	DestinationID int     `json:"destinationId" validate:"required,numeric,gt=0"`
	Weight        float64 `json:"weight" validate:"required,numeric,gt=0"`
}
