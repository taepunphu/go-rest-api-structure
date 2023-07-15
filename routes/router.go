package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/taepunphu/go-rest-api-structure/controllers"
)

func NewProductRoute(ctrl *controllers.ProductController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		productRouter := v1.Group("/product")
		{
			productRouter.GET("", ctrl.FindAll)
			productRouter.GET("/:productId", ctrl.FindById)
			productRouter.POST("", ctrl.Create)
			productRouter.PUT("/:productId", ctrl.Update)
			productRouter.DELETE("/:productId", ctrl.Delete)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return router
}
