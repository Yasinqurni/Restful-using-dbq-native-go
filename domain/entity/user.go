package entity

import "restful-api/domain/value_object"

type User struct {
	name     string
	email    value_object.Email
	phone    value_object.Phone
	password value_object.Password
	role     value_object.Role
	gender   value_object.Gender

	//relation with item
	item []*Item
}

type UserDTO struct {
	Name     string
	Email    value_object.Email
	Phone    value_object.Phone
	Password value_object.Password
	Role     value_object.Role
	Gender   value_object.Gender

	//relation with item
	Item []*Item
}

func NewUser(dto UserDTO) *User {
	return &User{
		name:     dto.Name,
		email:    dto.Email,
		phone:    dto.Phone,
		password: dto.Password,
		role:     dto.Role,
		gender:   dto.Gender,
		item:     dto.Item,
	}
}

func (e *User) GetName() string {
	return e.name
}

func (e *User) GetEmail() value_object.Email {
	return e.email
}

func (e *User) GetPhone() value_object.Phone {
	return e.phone
}

func (e *User) GetPassword() value_object.Password {
	return e.password
}

func (e *User) GetRole() value_object.Role {
	return e.role
}

func (e *User) GetGender() value_object.Gender {
	return e.gender
}
