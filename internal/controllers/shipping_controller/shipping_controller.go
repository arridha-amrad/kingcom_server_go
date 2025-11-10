package shippingcontroller

import (
	"io"
	"kingcom_api/internal/lib"
	cacheservice "kingcom_api/internal/services/cache_service"
	"net/http"
)

const (
	destinationUrl = "https://rajaongkir.komerce.id/api/v1/destination"
	calcCostUrl    = "https://rajaongkir.komerce.id/api/v1/calculate/district/domestic-cost"
	courier        = "jne:sicepat:ide:sap:jnt:ninja:tiki:lion:anteraja:pos:ncs:rex:rpx:sentral:star:wahana:dse"
	price          = "lowest"
)

type ShippingController struct {
	cacheService cacheservice.CacheService
	env          *lib.Env
	logger       *lib.Logger
	validator    *lib.Validator
}

func New(
	cacheService cacheservice.CacheService,
	env *lib.Env,
	logger *lib.Logger,
	validator *lib.Validator,
) *ShippingController {
	return &ShippingController{
		cacheService: cacheService,
		env:          env,
		logger:       logger,
		validator:    validator,
	}
}

func (ctrl *ShippingController) RunRequestToRajaOngkir(url, key string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctrl.logger.Error(err)
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ctrl.logger.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctrl.logger.Error(err)
		return nil, err
	}
	return body, nil
}
