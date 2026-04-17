package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Owner  string `json:"owner"`
	Amount string `json:"amount"`
}

func NewOrder(owner string, amount string) Order {
	order := Order{
		Owner:  owner,
		Amount: amount,
	}
	return order
}
