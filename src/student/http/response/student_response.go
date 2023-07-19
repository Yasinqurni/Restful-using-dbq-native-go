package response

import model "github-dbq/src/student/models"

type StudentResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Age         uint   `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
}

func (r *StudentResponse) List(students []*model.Student) []*StudentResponse {

	var responses []*StudentResponse
	for _, student := range students {
		response := StudentResponse{
			Id:          student.Id,
			Name:        student.Name,
			Age:         student.Age,
			DateOfBirth: student.DateOfBirth,
		}
		responses = append(responses, &response)
	}

	return responses
}

func (r *StudentResponse) Detail(student *model.Student) *StudentResponse {

	response := StudentResponse{
		Id:          student.Id,
		Name:        student.Name,
		Age:         student.Age,
		DateOfBirth: student.DateOfBirth,
	}

	return &response
}
