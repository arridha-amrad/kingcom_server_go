package response

import (
	"kingcom_api/internal/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResErr(
	c *gin.Context,
	logger *lib.Logger,
	statusCode int,
	err error,
	message string,
) {
	logger.Error(err)
	if message == "" {
		message = http.StatusText(statusCode)
	}
	c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}

func ResValidationErr(
	c *gin.Context,
	logger *lib.Logger,
	errors map[string]string,
) {
	logger.Error("== Validation Error ==", c.FullPath(), c.Request.Method)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"validationErrors": errors})
}
