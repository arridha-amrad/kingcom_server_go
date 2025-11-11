package productcontroller

import (
	"kingcom_api/internal/response"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) FetchProduct(c *gin.Context) {
	res := response.New(c, ctrl.logger)
	slug := c.Param("slug")

	product, err := ctrl.productService.FindBySlug(slug)

	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{"product": product})
}
