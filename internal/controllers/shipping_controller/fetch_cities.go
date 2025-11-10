package shippingcontroller

import (
	"encoding/json"
	"fmt"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) FetchCities(c *gin.Context) {
	provinceId := c.Param("provinceID")
	url := fmt.Sprintf("%s/city/%s", destinationUrl, provinceId)
	body, err := ctrl.RunRequestToRajaOngkir(url, ctrl.env.RajaOngkirAPIKey)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "Failed to fetch cities")
		return
	}
	var ror RajaOngkirResponse
	if err := json.Unmarshal(body, &ror); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"cities": ror.Data})
}
