package src

import (
	"database/sql"
	"github-dbq/config"
	controller "github-dbq/src/http/controllers"
	repository "github-dbq/src/repositories"
	service "github-dbq/src/services"
)

func InitModule(db *sql.DB, env *config.Config) (controller.StudentController, service.StudentService, repository.StudentRepository) {
	StudentRepository := repository.StudentRepositoryImpl(db, env)
	StudentService := service.StudentServiceImpl(StudentRepository)
	StudentController := controller.StudentControllerImpl(StudentService)

	return StudentController, StudentService, StudentRepository
}
