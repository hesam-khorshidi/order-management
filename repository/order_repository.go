package repository

import (
	"gorm.io/gorm"
	"order-management/entity"
)

type OrderRepository interface {
	UpdateOrderStatus(updateRequest entity.UpdateOrderRequest) error
	CheckNotifiedStatus(referenceId uint) (bool, error)
	GetOrderReceiverNumber(referenceId uint) (uint, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func (orderRepository *OrderRepositoryImpl) GetOrderReceiverNumber(referenceId uint) (uint, error) {
	var order entity.Order
	err := orderRepository.db.Model(&entity.Order{}).Where("referenceId = ?", referenceId).First(&order).Error
	if err != nil {
		return 0, err
	}
	return order.ReferenceId, nil
}

func (orderRepository *OrderRepositoryImpl) UpdateOrderStatus(updateRequest entity.UpdateOrderRequest) error {
	if updateRequest.Status == entity.PICKED_UP {
		return orderRepository.db.Model(&entity.Order{}).Where("order_reference_id = ?", updateRequest.OrderReferenceId).Update("status", updateRequest.Status).Update("notified_receiver", true).Error
	}
	return orderRepository.db.Model(&entity.Order{}).Where("order_reference_id = ?", updateRequest.OrderReferenceId).Update("status", updateRequest.Status).Error
}

func (orderRepository *OrderRepositoryImpl) CheckNotifiedStatus(referenceId uint) (bool, error) {
	var order entity.Order
	if err := orderRepository.db.Model(&entity.Order{}).Where("order_reference_id = ?", referenceId).First(&order).Error; err != nil {
		return false, err
	} else {
		return order.NotifiedReceiver, nil
	}
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}
