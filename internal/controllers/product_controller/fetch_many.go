package productcontroller

import (
	"kingcom_api/internal/models"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/response"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) FetchAllProducts(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	name := c.DefaultQuery("name", "")
	limitStr := c.DefaultQuery("limit", "9")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 9
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	filter := repositories.FindManyFilter{
		Name:   name,
		Limit:  limit,
		Offset: limit * (page - 1),
	}

	products, err := ctrl.productService.FindMany(filter)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	data := response.PaginatedResponse[models.Product]{
		Page:       page,
		PrevPage:   page - 1,
		NextPage:   page + 1,
		TotalPage:  int(math.Ceil(float64(products.Total) / float64(limit))),
		TotalItems: products.Total,
		Items:      products.Products,
	}

	res.ResOk(data)
}
