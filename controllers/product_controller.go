package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taepunphu/go-rest-api-structure/models/common"
	"github.com/taepunphu/go-rest-api-structure/models/features/product/request"
	"github.com/taepunphu/go-rest-api-structure/services"
	utils "github.com/taepunphu/go-rest-api-structure/utils/errors"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (ctrl *ProductController) Create(ctx *gin.Context) {
	createProductRequest := request.CreateProductRequest{}
	err := ctx.ShouldBindJSON(&createProductRequest)
	utils.ErrorPanic(err)

	ctrl.productService.Create(createProductRequest)
	resp := common.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, resp)
}

func (ctrl *ProductController) Update(ctx *gin.Context) {
	updateProductRequest := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	utils.ErrorPanic(err)

	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)
	updateProductRequest.Id = id

	ctrl.productService.Update(updateProductRequest)
	resp := common.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (ctrl *ProductController) Delete(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)

	ctrl.productService.Delete(id)

	resp := common.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (ctrl *ProductController) FindById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)

	productResponse := ctrl.productService.FindById(id)

	resp := common.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   productResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (ctrl *ProductController) FindAll(ctx *gin.Context) {
	productResponse := ctrl.productService.FindAll()
	resp := common.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   productResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}
