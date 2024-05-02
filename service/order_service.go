package service

import (
	"fmt"
	"order-management/entity"
	"order-management/repository"
)

type OrderService interface {
	UpdateOrderStatus(request entity.UpdateOrderRequest) error
}

type OrderServiceImpl struct {
	repository repository.OrderRepository
	smsService SmsService
}

func (service *OrderServiceImpl) UpdateOrderStatus(request entity.UpdateOrderRequest) error {
	// checking if updated status is PICKED_UP
	if request.Status == entity.PICKED_UP {
		notified, err := service.repository.CheckNotifiedStatus(request.OrderReferenceId)
		if err != nil {
			return err
		}
		//If user is notified before then we skip the notification step
		if notified {
			return service.repository.UpdateOrderStatus(request)
		}
		//if not we notify them and update the record
		receiverNumber, err := service.repository.GetOrderReceiverNumber(request.OrderReferenceId)
		if err != nil {
			return err
		}
		service.smsService.SendSms(fmt.Sprintf("Order number %d status changed to %s", request.OrderReferenceId, request.Status), receiverNumber)
		return service.repository.UpdateOrderStatus(request)
	} else {
		// Get customer phone number from database
		receiverNumber, err := service.repository.GetOrderReceiverNumber(request.OrderReferenceId)
		if err != nil {
			return err
		}
		// Notify User and update record
		service.smsService.SendSms(fmt.Sprintf("Order number %d status changed to %s", request.OrderReferenceId, request.Status), receiverNumber)
		return service.repository.UpdateOrderStatus(request)
	}
}

func NewOrderService(r repository.OrderRepository, smsService SmsService) OrderService {
	return &OrderServiceImpl{repository: r, smsService: smsService}
}
