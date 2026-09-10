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
