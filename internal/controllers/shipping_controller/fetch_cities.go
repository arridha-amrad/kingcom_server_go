package shippingcontroller

import (
	"encoding/json"
	"fmt"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) FetchCities(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	provinceId := c.Param("provinceID")

	url := fmt.Sprintf("%s/city/%s", destinationUrl, provinceId)
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
	c.JSON(http.StatusOK, gin.H{"cities": ror.Data})
}
