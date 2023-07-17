package model

import "time"

type Classroom struct {
	ID        uint      `dbq:"id"`
	Name      string    `dbq:"name"`
	Capacity  uint      `dbq:"capacity"`
	UpdatedAt time.Time `dbq:"updated_at"`
	CreatedAt time.Time `dbq:"created_at"`
	DeletedAt time.Time `dbq:"deleted_at"`
}
