package interfaces

import "github.com/yukinooz/go-task/service/domain"

// TaskRepository struct
type TaskRepository struct {
	sqlHandler SQLHandler
}

// NewTaskRepository create a new TaskRepository
func NewTaskRepository(handler SQLHandler) *TaskRepository {
	return &TaskRepository{
		sqlHandler: handler,
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
	return domain.Task{}, nil
}
