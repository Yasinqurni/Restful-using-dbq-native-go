package service

import (
	"context"
	request "github-dbq/src/http/requests"
	model "github-dbq/src/models"
	repository "github-dbq/src/repositories"
	"time"
)

type studentService struct {
	repository repository.StudentRepository
}

func StudentServiceImpl(repository repository.StudentRepository) StudentService {
	return &studentService{repository: repository}
}

func (s *studentService) Get() ([]*model.Student, error) {
	ctx := context.Background()
	students, err := s.repository.Get(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentService) GetByID(id uint) (*model.Student, error) {
	ctx := context.Background()
	student, err := s.repository.GetByID(id, ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentService) Update(name string, id uint) error {
	ctx := context.Background()
	return s.repository.Update(name, id, ctx)
}

func (s *studentService) Delete(id uint) error {
	ctx := context.Background()
	return s.repository.Delete(id, ctx)
}

func (s *studentService) Create(student request.CreateRequest) error {
	ctx := context.Background()

	data := model.Student{
		Name:        student.Name,
		Age:         student.Age,
		DateOfBirth: student.DateOfBirth,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return s.repository.Create(data, ctx)
}
