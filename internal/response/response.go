package response

import (
	"kingcom_api/internal/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	c      *gin.Context
	logger *lib.Logger
}

func New(
	c *gin.Context,
	logger *lib.Logger,
) *Response {
	return &Response{
		c:      c,
		logger: logger,
	}
}

func (h *Response) ResErr(statusCode int, err error, message string) {
	h.logger.Error(err)
	if message == "" {
		message = http.StatusText(statusCode)
	}
	h.c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}

func (h *Response) ResErrValidation(errors map[string]string) {
	h.logger.Error("== Validation Error ==", h.c.FullPath(), h.c.Request.Method)
	h.c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"validationErrors": errors})
}

func (h *Response) ResInternalServerErr(err error) {
	h.logger.Error(err)
	code := http.StatusInternalServerError
	h.c.AbortWithStatusJSON(code, gin.H{"message": http.StatusText(code)})
}

func (h *Response) ResErrUnauthorized(err error) {
	h.logger.Error(err)
	code := http.StatusUnauthorized
	h.c.AbortWithStatusJSON(code, gin.H{"message": http.StatusText(code)})
}

func (h *Response) ResOk(data any) {
	h.c.JSON(http.StatusOK, data)
}

func (h *Response) ResCreated(message string) {
	h.c.JSON(http.StatusCreated, gin.H{"message": message})
}

type PaginatedResponse[T any] struct {
	Page       int `json:"page"`
	PrevPage   int `json:"prevPage"`
	NextPage   int `json:"nextPage"`
	TotalPage  int `json:"totalPage"`
	TotalItems int `json:"totalItems"`
	Items      []T `json:"items"`
}
