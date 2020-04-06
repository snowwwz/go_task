package interfaces

import (
	"github.com/yukinooz/go_task/service/domain"
	"time"
)

// TaskRepository struct
type TaskRepository struct {
	sql SQLHandler
}

// NewTaskRepository create a new TaskRepository
func NewTaskRepository(handler SQLHandler) *TaskRepository {
	return &TaskRepository{
		sql: handler,
	}
}

// Add repository
func (repo TaskRepository) Add(name string, pri int, deadline time.Time) error {
	task := domain.Task{
		Name:      name,
		Priority:  pri,
		Deadline:  deadline,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return repo.sql.Create(&task)
}

// Delete repository
func (repo TaskRepository) Delete(id int) error {
	return repo.sql.Update(&domain.Task{}, "delete_flg", 1, "id = ?", id)
}

// List repository
func (repo TaskRepository) List(isAll bool) ([]domain.Task, error) {
	var tasks []domain.Task

	if isAll {
		if err := repo.sql.Select(&tasks, "delete_flg = ?", 0); err != nil {
			return nil, err
		}
	} else {
		if err := repo.sql.Select(&tasks, "delete_flg = ? AND status != ?", 0, 2); err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

// Change repository
func (repo TaskRepository) Change(id int, column string, data interface{}) error {
	return repo.sql.Update(&domain.Task{}, column, data, "id = ?", id)
}
