package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-management/controller"
	"order-management/repository"
	"order-management/service"
)

func InitializeLogisticsApiRoute(engine *gin.Engine, db *gorm.DB) {
	//Initializing provider repository, service and controller
	providerController := initializeProviderController(db)

	//Initializing order repository, service ad controller
	orderController := initializeOrderController(db)

	baseApi := engine.Group("/api/v1")
	{
		providerApi := baseApi.Group("/provider")
		{
			providerApi.PUT("/order", orderController.UpdateOrderStatus)
			providerApi.GET("/order/averageDelivery/:providerId", providerController.AverageDeliveryTimeOfSevenDays)
		}
		ordersApi := baseApi.Group("/orders")
		{
			ordersApi.GET("/tracking/:providerId/:orderReferenceId", providerController.OrderByUrlAndReference)
		}
	}
}

func initializeProviderController(db *gorm.DB) controller.ProviderController {
	providerRepository := repository.NewProviderRepository(db)
	providerService := service.NewProviderService(providerRepository)
	return controller.NewProviderController(providerService)
}
func initializeOrderController(db *gorm.DB) controller.OrderController {
	orderRepository := repository.NewOrderRepository(db)
	smsService := service.NewSmsService()
	orderService := service.NewOrderService(orderRepository, smsService)
	return controller.NewOrderController(orderService)
}
