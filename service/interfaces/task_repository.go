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

// Add create a task
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

// Delete a task
func (repo TaskRepository) Delete(id int) error {
	return repo.sql.Delete(&domain.Task{}, "id = ?", id)
}

// List tasks
func (repo TaskRepository) List() ([]domain.Task, error) {
	var tasks []domain.Task
	if err := repo.sql.SelectAll(&tasks, "delete_flg = ? AND status = ?", 0, 0); err != nil {
		return nil, err
	}
	return tasks, nil
}
