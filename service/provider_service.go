package service

import (
	"errors"
	"order-management/entity"
	"order-management/repository"
	"sort"
	"time"
)

type ProviderService interface {
	CalculateAverageDeliveryTimeOfLastSevenDays(providerId uint) ([]time.Duration, error)
	GetOrdersStatusWithProviderUrlAndReferenceId(providersUrl string, referenceId uint) (entity.Order, error)
}

type providerServiceImpl struct {
	providerRepository repository.ProviderRepository
}

func (service *providerServiceImpl) GetOrdersStatusWithProviderUrlAndReferenceId(providersUrl string, referenceId uint) (entity.Order, error) {
	//fetch all orders of
	orders, err := service.providerRepository.GetOrdersOfProviderWithProviderUrl(providersUrl)
	if err != nil {
		return entity.Order{}, err
	}
	for _, order := range orders {
		if order.ReferenceId == referenceId {
			return order, nil
		}
	}
	return entity.Order{}, errors.New("order not found")
}

func (service *providerServiceImpl) CalculateAverageDeliveryTimeOfLastSevenDays(providerId uint) ([]time.Duration, error) {
	//fetching all orders of last 7 days
	orders, err := service.providerRepository.GetOrdersOfLastSevenDays(providerId)
	ordersDeliveryTime := make([]time.Duration, 0)
	if err != nil {
		return ordersDeliveryTime, err
	}

	//calculating delivery time
	for _, order := range orders {
		ordersDeliveryTime = append(ordersDeliveryTime, order.DeliveryDate.Sub(order.PickUpDate))
	}

	//Sorting the orders delivery time
	sort.Slice(ordersDeliveryTime, func(i, j int) bool {
		return ordersDeliveryTime[i] > ordersDeliveryTime[j]
	})

	return ordersDeliveryTime, nil
}

func NewProviderService(providerRepository repository.ProviderRepository) ProviderService {
	return &providerServiceImpl{providerRepository: providerRepository}
}
