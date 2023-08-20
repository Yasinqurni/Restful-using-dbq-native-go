package entity

import (
	"errors"
	"restful-api/pkg"
)

type Item struct {
	name     string
	category string
	price    uint64
	quantity uint32

	//relation for user
	user *User
}

type ItemDTO struct {
	Name     string
	Category string
	Price    uint64
	Quantity uint32

	//relation for user
	User *User
}

func NewItem(dto ItemDTO) (*Item, error) {
	if dto.Name == "" {
		return nil, errors.New(pkg.EMPTY_NAME)
	}
	return &Item{
		name:     dto.Name,
		category: dto.Category,
		price:    dto.Price,
		quantity: dto.Quantity,
		user:     dto.User,
	}, nil
}

func (e Item) GetName() string {
	return e.name
}

func (e Item) GetCategory() string {
	return e.category
}

func (e Item) GetPrice() uint64 {
	return e.price
}

func (e Item) GetQty() uint32 {
	return e.quantity
}
