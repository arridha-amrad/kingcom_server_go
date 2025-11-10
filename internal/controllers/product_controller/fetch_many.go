package productcontroller

import (
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) FetchAllProducts(c *gin.Context) {
	products, err := ctrl.productService.FindMany()
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
