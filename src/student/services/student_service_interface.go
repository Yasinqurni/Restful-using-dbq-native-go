package service

import (
	model "github-dbq/src/student/models"
)

type StudentService interface {
	Get() (*[]model.Student, error)
	GetByID(id uint) (*model.Student, error)
}
