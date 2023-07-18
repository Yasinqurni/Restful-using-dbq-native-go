package service

import (
	"context"
	request "github-dbq/src/student/http/requests"
	model "github-dbq/src/student/models"
	repository "github-dbq/src/student/repositories"
	"time"
)

type studentService struct {
	repository repository.StudentRepository
	ctx        context.Context
}

func StudentServiceImpl(repository repository.StudentRepository) StudentService {
	return &studentService{
		repository: repository,
		ctx:        context.Background(),
	}
}

func (s *studentService) Get() ([]*model.Student, error) {

	students, err := s.repository.Get(s.ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentService) GetByID(id uint) (*model.Student, error) {

	student, err := s.repository.GetByID(id, s.ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentService) Update(name string, id uint) error {

	return s.repository.Update(name, id, s.ctx)
}

func (s *studentService) Delete(id uint) error {

	return s.repository.Delete(id, s.ctx)
}

func (s *studentService) Create(student request.CreateRequest) error {

	data := model.Student{
		Name:        student.Name,
		Age:         student.Age,
		DateOfBirth: student.DateOfBirth,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return s.repository.Create(data, s.ctx)
}
