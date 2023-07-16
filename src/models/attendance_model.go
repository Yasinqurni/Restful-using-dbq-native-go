package model

import (
	"database/sql"
	"time"
)

type Attendance struct {
	ID          uint
	Absen       bool
	Date        sql.NullTime
	StudentID   uint
	ClassroomID uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
	DeletedAt   time.Time

	//relations
	Student   *Student
	Classroom *Classroom
}
