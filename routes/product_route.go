package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taepunphu/go-rest-api-structure/controllers"
)

func NewProductRoute(ctrl *controllers.ProductController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api/v1")
	productRouter := baseRouter.Group("/product")
	productRouter.GET("", ctrl.FindAll)
	productRouter.GET("/:productId", ctrl.FindById)
	productRouter.POST("", ctrl.Create)
	productRouter.PUT("/:productId", ctrl.Update)
	productRouter.DELETE("/:productId", ctrl.Delete)

	return router
}