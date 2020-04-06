package usecase

import (
	"github.com/yukinooz/go_task/service/domain"
	"time"
)

// TaskRepository interface
type TaskRepository interface {
	Add(string, int, time.Time) error
	Delete(int) error
	List(bool) ([]domain.Task, error)
	Change(int, string, interface{}) error
}
