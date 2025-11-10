package shippingcontroller

import (
	"encoding/json"
	"fmt"
	"kingcom_api/internal/response"
	cacheservice "kingcom_api/internal/services/cache_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) FetchProvinces(c *gin.Context) {
	data, err := ctrl.cacheService.FindShippingProvinces(c.Request.Context())
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	if data == nil {
		var ror RajaOngkirResponse
		// make request to raja ongkir api
		url := fmt.Sprintf("%s/province", destinationUrl)
		body, err := ctrl.RunRequestToRajaOngkir(url, ctrl.env.RajaOngkirAPIKey)
		if err != nil {
			response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "Failed to fetch provinces")
			return
		}
		if err := json.Unmarshal(body, &ror); err != nil {
			response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
			return
		}

		if ror.Meta.Code == 200 {
			// save provinces into cache
			provinces := make([]cacheservice.Province, 0, len(ror.Data))
			for _, data := range ror.Data {
				provinces = append(provinces, cacheservice.Province{ID: data.ID, Name: data.Name})
			}
			params := cacheservice.SaveProvincesData{
				Provinces: provinces,
			}
			if err := ctrl.cacheService.SaveShippingProvinces(
				c.Request.Context(),
				params,
			); err != nil {
				response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"provinces": ror.Data})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provinces": data.Provinces})

}
