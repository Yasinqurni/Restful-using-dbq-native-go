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
	var ctx context.Context
	student, err := s.repository.GetByID(id, ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}
