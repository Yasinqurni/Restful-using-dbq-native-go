package request

type CreateRequest struct {
	Name        string `json:"name"`
	Age         uint   `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
}
