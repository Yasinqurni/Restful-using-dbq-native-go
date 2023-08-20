package entity

import (
	"errors"
	"restful-api/domain/value_object"
	"restful-api/pkg"
)

type User struct {
	name     string
	email    *value_object.Email
	phone    *value_object.Phone
	password *value_object.Password
	role     *value_object.Role
	gender   *value_object.Gender

	//relation with item
	item []*Item
}

type UserDTO struct {
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string
	Gender   string
}

func NewUser(dto UserDTO) (*User, error) {

	if dto.Name == "" {
		return nil, errors.New(pkg.EMPTY_NAME)
	}
	email, err := value_object.Email{}.NewEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	phone, err := value_object.Phone{}.NewPhone(dto.Phone)
	if err != nil {
		return nil, err
	}
	password, err := value_object.Password{}.NewPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	role, err := value_object.Role{}.NewRole(dto.Role)
	if err != nil {
		return nil, err
	}
	gender, err := value_object.Gender{}.NewGender(dto.Gender)
	if err != nil {
		return nil, err
	}

	return &User{
		name:     dto.Name,
		email:    email,
		phone:    phone,
		password: password,
		role:     role,
		gender:   gender,
	}, nil
}

func (e *User) SetItems(items []*Item) *User {
	e.item = items
	return e
}

func (e User) GetName() string {
	return e.name
}

func (e User) GetEmail() string {
	return e.email.GetEmail()
}

func (e User) GetPhone() string {
	return e.phone.GetPhone()
}

func (e User) GetPassword() string {
	return e.password.GetPassword()
}

func (e User) GetRole() string {
	return e.role.GetRole()
}

func (e User) GetGender() string {
	return e.gender.GetGender()
}
