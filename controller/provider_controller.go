package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-management/entity"
	"order-management/service"
	"strconv"
)

type ProviderController interface {
	AverageDeliveryTimeOfSevenDays(ctx *gin.Context)
	OrderByUrlAndReference(ctx *gin.Context)
}

type providerControllerImpl struct {
	providerService service.ProviderService
}

func (controller *providerControllerImpl) OrderByUrlAndReference(ctx *gin.Context) {
	// checking order reference and provider api address
	providerUrl := ctx.Param("providerUrl")
	reference := ctx.Param("orderReferenceId")
	referenceId, err := strconv.Atoi(reference)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: "Invalid order reference"})
		return
	}
	if providerUrl == "" {
		ctx.JSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: "Invalid provider api"})
		return
	}
	// fetching order from database
	order, err := controller.providerService.GetOrdersStatusWithProviderUrlAndReferenceId(providerUrl, uint(referenceId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: err.Error()})
		return
	}

	// formatting the final response
	ctx.JSON(http.StatusOK, entity.OrderTrackingResponse{Code: 200, Message: "Success", Data: entity.OrderOverview{Id: order.ID, ReferenceId: order.ReferenceId, Status: order.Status, CreatedAt: order.CreatedAt}})
}

func (controller *providerControllerImpl) AverageDeliveryTimeOfSevenDays(ctx *gin.Context) {
	// checking if providerId is present
	providerId := ctx.Param("providerId")
	if providerId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: "Missing providerId"})
		return
	}
	// checking if providerId is in right format
	id, err := strconv.Atoi(providerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: "bad providerId"})
		return
	}

	// calculating list of delivery times of the provider in the last 7 days
	average, err := controller.providerService.CalculateAverageDeliveryTimeOfLastSevenDays(uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.GeneralResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entity.GeneralResponse{Status: 200, Message: fmt.Sprintf("Provider average delivery time is %s", average)})
}

func NewProviderController(providerService service.ProviderService) ProviderController {
	return &providerControllerImpl{providerService: providerService}
}
