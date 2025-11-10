package dto

type CalcCost struct {
	OriginID      int `json:"originId" validate:"required,numeric,gt=0"`
	DestinationID int `json:"destinationId" validate:"required,numeric,gt=0"`
	Weight        int `json:"weight" validate:"required,numeric,gt=0"`
}
