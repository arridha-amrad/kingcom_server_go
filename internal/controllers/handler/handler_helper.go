package handler

import (
	"kingcom_api/internal/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerHelper struct {
	c      *gin.Context
	logger *lib.Logger
}

func NewHandlerHelper(
	c *gin.Context,
	logger *lib.Logger,
) *HandlerHelper {
	return &HandlerHelper{
		c:      c,
		logger: logger,
	}
}

func (h *HandlerHelper) ResErr(statusCode int, err error, message string) {
	h.logger.Error(err)
	if message == "" {
		message = http.StatusText(statusCode)
	}
	h.c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}

func (h *HandlerHelper) ResErrValidation(errors map[string]string) {
	h.logger.Error("== Validation Error ==", h.c.FullPath(), h.c.Request.Method)
	h.c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"validationErrors": errors})
}

func (h *HandlerHelper) ResInternalServerErr(err error) {
	h.logger.Error(err)
	code := http.StatusInternalServerError
	h.c.AbortWithStatusJSON(code, gin.H{"message": http.StatusText(code)})
}

func (h *HandlerHelper) ResErrUnauthorized(err error) {
	h.logger.Error(err)
	code := http.StatusUnauthorized
	h.c.AbortWithStatusJSON(code, gin.H{"message": http.StatusText(code)})
}

func (h *HandlerHelper) ResOk(data any) {
	h.c.JSON(http.StatusOK, data)
}

func (h *HandlerHelper) ResCreated(message string) {
	h.c.JSON(http.StatusCreated, gin.H{"message": message})
}
