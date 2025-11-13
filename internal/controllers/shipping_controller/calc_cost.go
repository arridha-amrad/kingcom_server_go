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
	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.CalcCost](c, ctrl.validator)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	form := url.Values{}
	form.Set("origin", strconv.Itoa(body.OriginID))
	form.Set("destination", strconv.Itoa(body.DestinationID))
	form.Set("weight", strconv.FormatFloat(body.Weight, 'f', 2, 64))
	form.Set("courier", courier)
	form.Set("price", price)

	req, err := http.NewRequest("POST", calcCostUrl, strings.NewReader(form.Encode()))
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("key", ctrl.env.RajaOngkirAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	var cr CostResponse
	if err := json.Unmarshal(respBody, &cr); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	max := min(len(cr.Data), 10)
	costs := cr.Data[:max]

	c.JSON(http.StatusOK, gin.H{"costs": costs})
}
