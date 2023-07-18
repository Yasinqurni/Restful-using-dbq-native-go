package model

import (
	"database/sql"
	"time"
)

type Attendance struct {
	ID          uint         `dbq:"id"`
	Absen       bool         `dbq:"absen"`
	Date        sql.NullTime `dbq:"date"`
	StudentID   uint         `dbq:"student_id"`
	ClassroomID uint         `dbq:"classroom_id"`
	UpdatedAt   time.Time    `dbq:"updated_at"`
	CreatedAt   time.Time    `dbq:"created_at"`
	DeletedAt   time.Time    `dbq:"deleted_at"`

	//relations
	Student   *Student   `dbq:"-"`
	Classroom *Classroom `dbq:"-"`
}
