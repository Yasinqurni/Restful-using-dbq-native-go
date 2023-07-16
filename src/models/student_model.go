package model

import (
	"time"
)

type Student struct {
	Id          uint
	Name        string
	DateOfBirth time.Time
	CreatedAt   time.Time
	DeletedAt   time.Time
}

func (s *Student) GetTableName() string {
	return "students"
}
