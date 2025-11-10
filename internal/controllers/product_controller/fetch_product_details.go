package productcontroller

import (
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) FetchProductDetails(c *gin.Context) {
	slug := c.Param("slug")
	product, err := ctrl.productService.FindBySlug(slug)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
