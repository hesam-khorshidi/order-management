package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerId   int     `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	Orders       []Order `json:"orders" gorm:"foreignKey:ID"`
}
