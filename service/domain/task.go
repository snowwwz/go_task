package domain

import "time"

// Task domain struct
type Task struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Status    int
	Priority  int
	Deadline  time.Time
	DeleteFlg int
	CreatedAt time.Time
	UpdatedAt time.Time
}
