package controller

import (
	"encoding/json"
	service "github-dbq/src/student/services"
	"net/http"
	"strconv"
)

type studentController struct {
	service service.StudentService
}

func StudentControllerImpl(service service.StudentService) StudentController {
	return &studentController{service: service}
}

func (c *studentController) Get(w http.ResponseWriter, r *http.Request) {

	students, err := c.service.Get()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func (c *studentController) GetByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	num, _ := strconv.ParseUint(id, 10, 64)

	student, err := c.service.GetByID(uint(num))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
