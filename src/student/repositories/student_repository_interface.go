package repository

import (
	"context"
	model "github-dbq/src/student/models"
)

type StudentRepository interface {
	Get(ctx context.Context) (*[]model.Student, error)
	GetByID(id uint, ctx context.Context) (*model.Student, error)
	// Create(student model.Student) error
}
