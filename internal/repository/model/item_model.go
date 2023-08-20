package model

import (
	"errors"
	"restful-api/domain/entity"
	"restful-api/pkg"
	"time"
)

const (
	TABLE_NAME = "items"
)

type Item struct {
	ID        string     `dbq:"id"`
	UserID    string     `dbq:"user_id"`
	Name      string     `dbq:"name"`
	Category  string     `dbq:"category"`
	Price     uint64     `dbq:"price"`
	Quantity  uint32     `dbq:"quantity"`
	CreatedAt time.Time  `dbq:"created_at"`
	UpdatedAt *time.Time `dbq:"updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at"`

	//relation for user
	User *User `dbq:"-"`
}

func (m Item) GetTableName() string {
	return TABLE_NAME
}

func MappingManyItem(items []*Item) ([]*entity.Item, error) {

	if len(items) == 0 || items == nil {
		return nil, errors.New(pkg.NOT_FOUND_ITEM)
	}

	var itemEntities []*entity.Item

	for _, item := range items {
		itemDTO := entity.ItemDTO{
			Name:     item.Name,
			Category: item.Category,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
		itemEntity, err := entity.NewItem(itemDTO)
		if err != nil {
			return nil, err
		}
		itemEntities = append(itemEntities, itemEntity)
	}

	return itemEntities, nil
}

func MappingSingleItem(item *Item) (*entity.Item, error) {

	if item == nil {
		return nil, errors.New(pkg.NOT_FOUND_ITEM)
	}

	itemDTO := entity.ItemDTO{
		Name:     item.Name,
		Category: item.Category,
		Price:    item.Price,
		Quantity: item.Quantity,
	}
	itemEntity, err := entity.NewItem(itemDTO)
	if err != nil {
		return nil, err
	}

	return itemEntity, nil
}
