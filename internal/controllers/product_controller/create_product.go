package productcontroller

import (
	"errors"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	body, errValidation := request.GetBody[dto.CreateProduct](c, ctrl.validator)
	if errValidation != nil {
		response.ResValidationErr(c, ctrl.logger, errValidation)
		return
	}

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, err, "")
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	user, err := ctrl.userService.FindById(userId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if user == nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "")
		return
	}
	if user.Role != models.RoleAdmin {
		response.ResErr(c, ctrl.logger, http.StatusForbidden, err, "You are not an admin")
		return
	}

	images := make([]models.ProductImage, 0, len(body.Images))
	for _, img := range body.Images {
		images = append(images, models.ProductImage{Url: img})
	}

	newProductSlug := utils.ToSlug(body.Name)
	existingProduct, err := ctrl.productService.FindBySlug(newProductSlug)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if existingProduct != nil {
		err = errors.New("slug has been used")
		response.ResErr(c, ctrl.logger, http.StatusConflict, err, err.Error())
		return
	}

	product := models.Product{
		Name:          body.Name,
		Weight:        body.Weight,
		Price:         body.Price,
		Description:   body.Description,
		Stock:         body.Stock,
		Specification: body.Specification,
		VideoUrl:      body.VideoUrl,
		Slug:          newProductSlug,
		Discount:      body.Discount,
		Images:        images,
	}

	if err := ctrl.productService.InsertProduct(&product); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New product created successfully"})
}
