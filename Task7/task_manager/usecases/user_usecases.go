package usecases

import (
	"task_manager/domain"
	"task_manager/repositories"
)

type UserUseCase interface {
	RegisterUser(user domain.User) (domain.User, error)
	LoginUser(username, password string) (domain.User, error)
	PromoteUser(id string) (domain.User, error)
}

type userUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return &userUseCase{repo}
}

func (u *userUseCase) RegisterUser(user domain.User) (domain.User, error) {
	return u.repo.CreateUser(user)
}

func (u *userUseCase) LoginUser(username, password string) (domain.User, error) {
	return u.repo.AuthenticateUser(username, password)
}

func (u *userUseCase) PromoteUser(id string) (domain.User, error) {
	return u.repo.PromoteUser(id)
}
