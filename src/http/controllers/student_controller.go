package controller

import (
	"encoding/json"
	"fmt"
	service "github-dbq/src/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
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
