package repository_mock

import (
	"context"
	model "github-dbq/src/student/models"
	"time"
)

type MockStudentRepository struct{}

func (m *MockStudentRepository) Get(ctx context.Context) ([]*model.Student, error) {

	student := &model.Student{
		Id:          1,
		Name:        "John Doe",
		Age:         20,
		DateOfBirth: "2003-01-01",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return []*model.Student{student}, nil
}

func (m *MockStudentRepository) GetByID(id uint, ctx context.Context) (*model.Student, error) {

	student := &model.Student{
		Id:          1,
		Name:        "John Doe",
		Age:         20,
		DateOfBirth: "2003-01-01",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return student, nil
}

func (m *MockStudentRepository) Update(name string, id uint, ctx context.Context) error {
	return nil
}

func (m *MockStudentRepository) Delete(id uint, ctx context.Context) error {
	return nil
}

func (m *MockStudentRepository) Create(data model.Student, ctx context.Context) error {
	return nil
}
