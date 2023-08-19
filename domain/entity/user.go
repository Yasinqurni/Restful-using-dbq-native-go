package entity

import "restful-api/domain/value_object"

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
