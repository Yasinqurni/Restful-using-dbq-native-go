package model

import "time"

type Classroom struct {
	ID        uint
	Name      string
	RoomSize  uint
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}
