package repository

import (
	"context"
	"study/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	err := r.pool.QueryRow(ctx,
		"INSERT INTO users (username, age) VALUES ($1, $2) RETURNING id",
		u.Username, u.Age,
	).Scan(&u.ID)
	return err
}
