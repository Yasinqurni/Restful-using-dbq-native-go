package repository

import (
	"context"
	"restful-api/domain/entity"
)

type ItemRepository interface {
	Get(ctx context.Context) ([]*entity.Item, error)
	GetByID(id uint, ctx context.Context) (*entity.Item, error)
	Update(item entity.ItemDTO, id uint, ctx context.Context) error
	Delete(id uint, ctx context.Context) error
	Create(item entity.ItemDTO, ctx context.Context) error
}
