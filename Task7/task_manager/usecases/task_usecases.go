package usecases

import (
	"task_manager/domain"
	"task_manager/repositories"
)

type TaskUseCase interface {
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
}

type taskUseCase struct {
	repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) TaskUseCase {
	return &taskUseCase{repo}
}

func (u *taskUseCase) CreateTask(task domain.Task) (domain.Task, error) {
	return u.repo.CreateTask(task)
}

func (u *taskUseCase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	return u.repo.UpdateTask(id, task)
}

func (u *taskUseCase) DeleteTask(id string) error {
	return u.repo.DeleteTask(id)
}

func (u *taskUseCase) GetTasks() ([]domain.Task, error) {
	return u.repo.GetTasks()
}

func (u *taskUseCase) GetTaskByID(id string) (domain.Task, error) {
	return u.repo.GetTaskByID(id)
}
