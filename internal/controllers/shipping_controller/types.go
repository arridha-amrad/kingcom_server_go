package shippingcontroller

type RajaOngkirResponse struct {
	MetaResponse
	Data []ResponseData
}

type ResponseData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MetaResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
}

type CostResponse struct {
	MetaResponse
	Data []ShippingCost
}

type ShippingCost struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Service     string  `json:"service"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	Etd         string  `json:"etd"`
}
