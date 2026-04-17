package repository

import (
	db2 "order/db"
	"order/models"
)

type OrderRepository struct {
	Database *db2.Db
}

func NewOrderRepository(database *db2.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *models.Order) (*models.Order, error) {
	result := repo.Database.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetAll() ([]*models.Order, error) {
	var orders []*models.Order
	result := repo.Database.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (repo *OrderRepository) GetById(id uint) (*models.Order, error) {
	var order models.Order
	result := repo.Database.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
