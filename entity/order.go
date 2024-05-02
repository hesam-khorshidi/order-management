package entity

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ReferenceId         uint        `json:"reference_id"`
	Provider            uint        `json:"provider" json:"providerId"`
	Status              OrderStatus `json:"status"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"-"`
	NotifiedReceiver    bool        `json:"-" gorm:"default:false"`
	PickUpDate          time.Time   `json:"-"`
	DeliveryDate        time.Time   `json:"-"`
	DeliveryAddress     string      `json:"delivery_address"`
	SendersPhoneNumber  uint        `json:"senders_phone_number"`
	ReceiverName        string      `json:"receiver_name"`
	ReceiverPhoneNumber uint        `json:"receiver_phone_number"`
}

// OrderStatus
// Defining order status enum
type OrderStatus string

const (
	PENDING       = "PENDING"
	IN_PROGRESS   = "IN_PROGRESS"
	PROVIDER_SEEN = "PROVIDER_SEEN"
	PICKED_UP     = "PICKED_UP"
	DELIVERED     = "DELIVERED"
)

type UpdateOrderRequest struct {
	OrderReferenceId uint
	Status           OrderStatus
}

// BeforeUpdate
// This hook checks if the order is updated in the last 24 hours or if it's in delivered or pending status
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	if time.Now().Sub(order.UpdatedAt).Hours() < 24 {
		err = errors.New("this order has already been updated in the last 24 hours")
	}
	if order.Status == PENDING {
		err = errors.New("this order is in pending state and cannot be updated")
	}
	if order.Status == DELIVERED {
		err = errors.New("this order is in delivered state and cannot be updated")
	}
	return
}

// AfterSave
// This hook updates updated_at when order is created
func (order *Order) AfterSave(tx *gorm.DB) (err error) {
	tx.Model(&Order{}).Where("reference_id = ?", order.ReferenceId).Update("updated_at", time.Now())
	return
}

// AfterUpdate
// This hook updates updated_at field everytime the order is updated
func (order *Order) AfterUpdate(tx *gorm.DB) (err error) {
	tx.Model(&Order{}).Where("reference_id = ?", order.ReferenceId).Update("updated_at", time.Now())
	return
}
