package model

import (
	"time"
)

type Student struct {
	Id          uint      `dbq:"id"`
	Name        string    `dbq:"name"`
	DateOfBirth time.Time `dbq:"date_of_birth"`
	UpdatedAt   time.Time `dbq:"updated_at"`
	CreatedAt   time.Time `dbq:"created_at"`
	DeletedAt   time.Time `dbq:"deleted_at"`
}

func (s *Student) GetTableName() string {
	return "students"
}
