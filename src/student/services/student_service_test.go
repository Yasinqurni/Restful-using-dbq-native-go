package service_test

import (
	repository "github-dbq/mock/student/repository"
	request "github-dbq/src/student/http/requests"
	service "github-dbq/src/student/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo = repository.MockStudentRepository{}

func TestGet(t *testing.T) {

	service := service.StudentServiceImpl(&repo)

	students, err := service.Get()

	assert.NotNil(t, students)
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	service := service.StudentServiceImpl(&repo)

	studentId := 1
	student, err := service.GetByID(uint(studentId))

	assert.NotNil(t, student)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	service := service.StudentServiceImpl(&repo)

	studentId := 1
	name := "name"
	err := service.Update(name, uint(studentId))

	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	service := service.StudentServiceImpl(&repo)

	studentID := uint(1)
	err := service.Delete(studentID)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {

	service := service.StudentServiceImpl(&repo)

	studentRequest := request.CreateRequest{
		Name:        "yasin",
		Age:         12,
		DateOfBirth: "12-12-1995",
	}
	err := service.Create(studentRequest)
	assert.NoError(t, err)
}
