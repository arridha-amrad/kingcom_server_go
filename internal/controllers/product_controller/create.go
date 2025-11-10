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

	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.CreateProduct](c, ctrl.validator)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	user, err := ctrl.userService.FindById(userId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user == nil {
		res.ResErr(http.StatusNotFound, err, "")
		return
	}
	if user.Role != models.RoleAdmin {
		res.ResErr(http.StatusForbidden, err, "")
		return
	}

	images := make([]models.ProductImage, 0, len(body.Images))
	for _, img := range body.Images {
		images = append(images, models.ProductImage{Url: img})
	}

	newProductSlug := utils.ToSlug(body.Name)
	existingProduct, err := ctrl.productService.FindBySlug(newProductSlug)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if existingProduct != nil {
		err = errors.New("slug has been used")
		res.ResErr(http.StatusConflict, err, err.Error())
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
		res.ResInternalServerErr(err)
		return
	}

	res.ResCreated("New product created successfully")

}
