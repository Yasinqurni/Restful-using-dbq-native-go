package repository

import (
	"context"
	model "github-dbq/src/student/models"
)

type StudentRepository interface {
	Get(ctx context.Context) ([]*model.Student, error)
	GetByID(id uint, ctx context.Context) (*model.Student, error)
	Update(name string, id uint, ctx context.Context) error
	Delete(id uint, ctx context.Context) error
	Create(student model.Student, ctx context.Context) error
}
