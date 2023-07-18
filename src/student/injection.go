package src

import (
	"database/sql"
	controller "github-dbq/src/student/http/controllers"
	repository "github-dbq/src/student/repositories"
	service "github-dbq/src/student/services"
)

func InitModule(db *sql.DB) (controller.StudentController, service.StudentService, repository.StudentRepository) {
	StudentRepository := repository.StudentRepositoryImpl(db)
	StudentService := service.StudentServiceImpl(StudentRepository)
	StudentController := controller.StudentControllerImpl(StudentService)

	return StudentController, StudentService, StudentRepository
}
