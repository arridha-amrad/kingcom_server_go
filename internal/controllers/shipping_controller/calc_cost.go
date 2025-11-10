package shippingcontroller

import (
	"encoding/json"
	"io"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctrl *ShippingController) CalcCost(c *gin.Context) {
	body, errValidation := request.GetBody[dto.CalcCost](c, ctrl.validator)
	if errValidation != nil {
		response.ResValidationErr(c, ctrl.logger, errValidation)
		return
	}

	form := url.Values{}
	form.Set("origin", strconv.Itoa(body.OriginID))
	form.Set("destination", strconv.Itoa(body.DestinationID))
	form.Set("weight", strconv.Itoa(body.Weight))
	form.Set("courier", courier)
	form.Set("price", price)

	req, err := http.NewRequest("POST", calcCostUrl, strings.NewReader(form.Encode()))
	if err != nil {
		ctrl.logger.Info("failed to create post req to calc the cost")
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("key", ctrl.env.RajaOngkirAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	var cr CostResponse
	if err := json.Unmarshal(respBody, &cr); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{"costs": cr.Data})
}
