package service

import (
	"context"
	model "github-dbq/src/models"
	repository "github-dbq/src/repositories"
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
