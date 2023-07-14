package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/taepunphu/go-rest-api-structure/entities"
	"github.com/taepunphu/go-rest-api-structure/models/features/product/request"
	"github.com/taepunphu/go-rest-api-structure/models/features/product/response"
	"github.com/taepunphu/go-rest-api-structure/repositories"
	"github.com/taepunphu/go-rest-api-structure/utils"
)

type ProductService interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productId int)
	FindById(productId int) response.ProductDetailResponse
	FindAll() []response.ProductResponse
}

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	Validate          *validator.Validate
}

// Create implements ProductService.
func (ps *ProductServiceImpl) Create(product request.CreateProductRequest) {
	err := ps.Validate.Struct(product)
	utils.ErrorPanic(err)
	productModel := entities.ProductEntity{
		ProductName: product.ProductName,
		AvailableStock: product.AvailableStock,
	}
	ps.ProductRepository.Save(productModel)
}

// Delete implements ProductService.
func (ps *ProductServiceImpl) Delete(productId int) {
	ps.ProductRepository.Delete(productId)
}

// FindAll implements ProductService.
func (ps *ProductServiceImpl) FindAll() []response.ProductResponse {
	result := ps.ProductRepository.FindAll()

	var products []response.ProductResponse
	for _, value := range result{
		p := response.ProductResponse{
			Id: value.Id,
			ProductName: value.ProductName,
			AvailableStock: value.AvailableStock,
		}
		products = append(products, p)
	}

	return products
}

// FindById implements ProductService.
func (ps *ProductServiceImpl) FindById(productId int) response.ProductDetailResponse {
	product, err := ps.ProductRepository.FindById(productId)
	utils.ErrorPanic(err)

	productResult := response.ProductDetailResponse{
		Id: product.Id,
		ProductName: product.ProductName,
		AvailableStock: product.AvailableStock,
	}

	return productResult
}

// Update implements ProductService.
func (ps *ProductServiceImpl) Update(product request.UpdateProductRequest) {
	productData, err := ps.ProductRepository.FindById(product.Id)
	utils.ErrorPanic(err)
	productData.ProductName = product.ProductName
	productData.AvailableStock = product.AvailableStock
	ps.ProductRepository.Update(productData)
}

func NewProductservice(productRepository repositories.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate: validate,
	}
}
