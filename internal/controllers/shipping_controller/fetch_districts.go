package shippingcontroller

import (
	"encoding/json"
	"fmt"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) GetDistricts(c *gin.Context) {
	cityId := c.Param("cityID")
	url := fmt.Sprintf("%s/district/%s", destinationUrl, cityId)
	body, err := ctrl.RunRequestToRajaOngkir(url, ctrl.env.RajaOngkirAPIKey)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "failed to fetch districts")
		return
	}
	var ror RajaOngkirResponse
	if err := json.Unmarshal(body, &ror); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"districts": ror.Data})
}
