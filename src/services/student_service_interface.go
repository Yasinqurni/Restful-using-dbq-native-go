package service

import (
	model "github-dbq/src/models"
)

type StudentService interface {
	Get() ([]*model.Student, error)
	GetByID(id uint) (*model.Student, error)
	Update(name string, id uint) error
}
