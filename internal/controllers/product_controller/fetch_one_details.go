package productcontroller

import (
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) FetchOneWithDetails(c *gin.Context) {
	hh := response.New(c, ctrl.logger)
	slug := c.Param("slug")

	product, err := ctrl.productService.FindBySlug(slug)

	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
