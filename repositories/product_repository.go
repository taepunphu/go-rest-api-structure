package repositories

import (
	"errors"

	"github.com/taepunphu/go-rest-api-structure/entities"
	"github.com/taepunphu/go-rest-api-structure/models/features/product/request"
	"github.com/taepunphu/go-rest-api-structure/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product entities.ProductEntity)
	Update(product entities.ProductEntity)
	Delete(productId int)
	FindById(productId int) (product entities.ProductEntity, err error)
	FindAll() []entities.ProductEntity
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements ProductRepository.
func (pr *ProductRepositoryImpl) Delete(productId int) {
	var product entities.ProductEntity
	result := pr.Db.Where("id = ?", productId).Delete(product)
	utils.ErrorPanic(result.Error)
}

// FindAll implements ProductRepository.
func (pr *ProductRepositoryImpl) FindAll() []entities.ProductEntity {
	var product []entities.ProductEntity
	result := pr.Db.Find(&product)
	utils.ErrorPanic(result.Error)
	return product
}

// FindById implements ProductRepository.
func (pr *ProductRepositoryImpl) FindById(productId int) (product entities.ProductEntity, err error) {
	var productEntity entities.ProductEntity
	result := pr.Db.Find(&productEntity, productId)
	if result != nil {
		return productEntity, nil
	} else {
		return productEntity, errors.New("product is not found")
	}
}

// Save implements ProductRepository.
func (pr *ProductRepositoryImpl) Save(product entities.ProductEntity) {
	result := pr.Db.Create(&product)
	utils.ErrorPanic(result.Error)
}

// Update implements ProductRepository.
func (pr *ProductRepositoryImpl) Update(product entities.ProductEntity) {
	var updateProduct = request.UpdateProductRequest{
		Id : product.Id,
		ProductName: product.ProductName,
		AvailableStock: product.AvailableStock,
	}

	result := pr.Db.Model(&product).Updates(updateProduct)
	utils.ErrorPanic(result.Error)
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: DB,
	}
}
