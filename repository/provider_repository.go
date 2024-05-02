package repository

import (
	"gorm.io/gorm"
	"order-management/entity"
	"time"
)

type ProviderRepository interface {
	GetOrdersOfLastSevenDays(providerId uint) ([]entity.Order, error)
	GetOrdersOfProviderWithProviderUrl(providerUrl string) ([]entity.Order, error)
}

type providerRepositoryImpl struct {
	db *gorm.DB
}

func (repository *providerRepositoryImpl) GetOrdersOfLastSevenDays(providerId uint) ([]entity.Order, error) {
	var orders []entity.Order
	startDate := time.Now().AddDate(0, 0, -7) // Seven days ago
	endDate := time.Now()
	err := repository.db.Where("id = ? AND date BETWEEN ? AND ?", providerId, startDate, endDate).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (repository *providerRepositoryImpl) GetOrdersOfProviderWithProviderUrl(providerUrl string) ([]entity.Order, error) {
	var orders []entity.Order
	err := repository.db.Where("provider_url = ?", providerUrl).Preload("Orders").Find(&orders).Error
	return orders, err
}

func NewProviderRepository(db *gorm.DB) ProviderRepository {
	return &providerRepositoryImpl{db: db}
}
