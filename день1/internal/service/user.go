package service

import (
	"context"
	"errors"
	"study/internal/model"
	"study/internal/repository"
)

var (
	ErrName = errors.New("Invalid name")
	ErrAge  = errors.New("Invalid Age")
	ErrId   = errors.New("Invalid Id")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, u *model.User) error {
	if u.Username == "" || u.Username == " " {
		return ErrName
	}
	if u.Age <= 0 || u.Age > 150 {
		return ErrAge
	}

	return s.repo.Create(ctx, u)
}
func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}
func (s *UserService) GetById(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {

		return nil, ErrId
	}
	return s.repo.GetById(ctx, id)
}
