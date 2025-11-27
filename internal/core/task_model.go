package core

import "time"

type Task struct {
	ID          uint64
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
