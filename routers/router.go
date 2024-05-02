package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeLogisticsApiRoute(engine *gin.Engine, db *gorm.DB) {
	baseApi := engine.Group("/api/v1")
	{
		providerApi := baseApi.Group("/provider")
		{
			providerApi.PUT("/order/:orderReferenceId/:status")
		}
		ordersApi := baseApi.Group("/orders")
		{
			ordersApi.GET("/tracking/:providerId/:orderReferenceId")
		}
	}
}
