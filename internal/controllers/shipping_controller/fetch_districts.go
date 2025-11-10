package shippingcontroller

import (
	"encoding/json"
	"fmt"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) GetDistricts(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	cityId := c.Param("cityID")

	url := fmt.Sprintf("%s/district/%s", destinationUrl, cityId)

	body, err := ctrl.RunRequestToRajaOngkir(url, ctrl.env.RajaOngkirAPIKey)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	var ror RajaOngkirResponse
	if err := json.Unmarshal(body, &ror); err != nil {
		res.ResInternalServerErr(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"districts": ror.Data})
}
