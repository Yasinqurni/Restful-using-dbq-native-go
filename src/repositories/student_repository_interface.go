package repository

import (
	"context"
	model "github-dbq/src/models"
)

type StudentRepository interface {
	Get(ctx context.Context) ([]*model.Student, error)
	GetByID(id uint, ctx context.Context) (*model.Student, error)
	Update(name string, id uint, ctx context.Context) error
}
