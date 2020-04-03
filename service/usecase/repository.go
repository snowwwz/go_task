package usecase

import "github.com/yukinooz/go-task/service/domain"

// TaskRepository interface
type TaskRepository interface {
	Add() (domain.Task, error)
	Delete() (domain.Task, error)
	List() (domain.Task, error)
}
