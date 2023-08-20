package model

import "time"

type User struct {
	ID        string     `dbq:"id"`
	Name      string     `dbq:"name"`
	Email     string     `dbq:"email"`
	Phone     string     `dbq:"phone"`
	Password  string     `dbq:"password"`
	Role      string     `dbq:"role"`
	Gender    string     `dbq:"gender"`
	CreatedAt time.Time  `dbq:"created_at"`
	UpdatedAt *time.Time `dbq:"updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at"`

	//relation with item
	Item []*Item `dbq:"-"`
}
