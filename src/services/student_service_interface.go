package service

import (
	request "github-dbq/src/http/requests"
	model "github-dbq/src/models"
)

type StudentService interface {
	Get() ([]*model.Student, error)
	GetByID(id uint) (*model.Student, error)
	Update(name string, id uint) error
	Delete(id uint) error
	Create(student request.CreateRequest) error
}
