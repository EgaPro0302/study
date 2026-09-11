package repository

import (
	"context"
	"errors"
	"study/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

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

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	rows, err := r.pool.Query(ctx, "SELECT id,username,age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Age,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func (r *UserRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	var u model.User
	rows := r.pool.QueryRow(ctx,
		"SELECT id, username,age FROM users WHERE id =$1",
		id)
	if err := rows.Scan(&u.ID, &u.Username, &u.Age); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil

}
