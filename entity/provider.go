package entity

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	ProviderName string  `json:"provider_name"`
	URL          string  `json:"url"`
	Orders       []Order `json:"orders" gorm:"foreignKey:ID"`
}
