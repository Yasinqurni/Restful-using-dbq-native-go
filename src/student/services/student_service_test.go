package service_test

import (
	"context"
	model "github-dbq/src/student/models"
	service "github-dbq/src/student/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockStudentRepository struct{}

func (m *mockStudentRepository) Get(ctx context.Context) ([]*model.Student, error) {

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

func (m *mockStudentRepository) GetByID(id uint, ctx context.Context) (*model.Student, error) {

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

func (m *mockStudentRepository) Update(name string, id uint, ctx context.Context) error {
	return nil
}

func (m *mockStudentRepository) Delete(id uint, ctx context.Context) error {
	return nil
}

func (m *mockStudentRepository) Create(data model.Student, ctx context.Context) error {
	return nil
}

func TestGet(t *testing.T) {
	repo := &mockStudentRepository{}
	service := service.StudentServiceImpl(repo)

	students, err := service.Get()

	assert.NotNil(t, students)
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	repo := &mockStudentRepository{}
	service := service.StudentServiceImpl(repo)

	studentId := 1
	student, err := service.GetByID(uint(studentId))

	assert.NotNil(t, student)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	repo := &mockStudentRepository{}
	service := service.StudentServiceImpl(repo)

	studentId := 1
	name := "name"
	err := service.Update(name, uint(studentId))

	assert.Nil(t, err)
}

// func TestDelete(t *testing.T) {
// 	repo := &mockStudentRepository{}
// 	service := StudentServiceImpl(repo)

// 	studentID := uint(1)
// 	err := service.Delete(studentID)
// 	assert.NoError(t, err)
// }

// func TestCreate(t *testing.T) {
// 	repo := &mockStudentRepository{}
// 	service := StudentServiceImpl(repo)

// 	studentRequest := requests.CreateRequest{
// 		Name:        "Jane Doe",
// 		Age:         21,
// 		DateOfBirth: "2002-01-01",
// 	}
// 	err := service.Create(studentRequest)
// 	assert.NoError(t, err)
// }
