package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ReferenceId         uint        `json:"reference_id"`
	Provider            uint        `json:"provider" json:"providerId"`
	Status              OrderStatus `json:"status"`
	CreatedAt           time.Time   `json:"created_at"`
	LastUpdated         time.Time   `json:"-"`
	NotifiedReceiver    bool        `json:"-" gorm:"default:false"`
	PickUpDate          time.Time   `json:"-"`
	DeliveryDate        time.Time   `json:"-"`
	DeliveryAddress     string      `json:"delivery_address"`
	SendersPhoneNumber  int32       `json:"senders_phone_number"`
	ReceiverName        string      `json:"receiver_name"`
	ReceiverPhoneNumber int32       `json:"receiver_phone_number"`
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
