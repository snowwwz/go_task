package domain

import "time"

// Task domain struct
type Task struct {
	ID        int
	Name      string
	Status    int
	Priority  string
	Deadline  time.Time
	DeleteFlg int
	CreatedAt time.Time
	UpdatedAt time.Time
}
