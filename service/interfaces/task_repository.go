package interfaces

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/yukinooz/go_task/service/domain"
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
func (repo TaskRepository) Add() (domain.Task, error) {
	return domain.Task{}, nil
}

// Delete a task
func (repo TaskRepository) Delete() (domain.Task, error) {
	return domain.Task{}, nil
}

// List tasks
func (repo TaskRepository) List() (domain.Task, error) {
	// rows, _ := repo.sql.Model(&tasks{}).Where("status = ?", 0).Select("id, name, priority, deadline").Rows()
	eventsEx := []domain.Task{}

	db := repo.sql.Where(&eventsEx, "delete_flg = ?", 0)
	spew.Dump(eventsEx)
	return domain.Task{}, nil
}
