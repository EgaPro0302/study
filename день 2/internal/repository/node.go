package repository

import (
	"context"
	"errors"
	"study/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// глобальна ошибка что бы было более четабельнее
var ErrTaskNotFound = errors.New("task not found")

// делает кабель подлкючения к бд
type NoteRepository struct {
	pool *pgxpool.Pool
}

// позволяет подлкючиться к этому кабелю
func NewNoteRepository(pool *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{pool: pool}
}

// получается создоем что то в бд
func (r *NoteRepository) Create(ctx context.Context, n *model.Note) error {
	err := r.pool.QueryRow(ctx,
		"INSERT INTO Task (title, content) VALUES ($1, $2) RETURNING id, createdat",
		n.Title, n.Content,
	).Scan(&n.Id, &n.CreatedAt)
	return err
}

// показываем все что есть в таблице
func (r *NoteRepository) GetAll(ctx context.Context) ([]model.Note, error) {
	var notes []model.Note
	rows, err := r.pool.Query(ctx, "SELECT id, title, content, createdat FROM Task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var n model.Note
		if err := rows.Scan(
			&n.Id,
			&n.Title,
			&n.Content,
			&n.CreatedAt,
		); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}

// находит человека по айдишнику
func (r *NoteRepository) GetById(ctx context.Context, id int) (*model.Note, error) {
	var n model.Note
	rows := r.pool.QueryRow(ctx,
		"SELECT id,title,content, createdat FROM Task WHERE id = $1",
		id,
	)
	if err := rows.Scan(&n.Id, &n.Title, &n.Content, &n.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return &n, nil
}

// удаление по айди
func (r *NoteRepository) DeleteById(ctx context.Context, id int) error {
	tag, err := r.pool.Exec(ctx, "DELETE FROM Task WHERE id = $1", id)
	if err != nil {
		return err
	}
	// тут проверяем есть ли у нас вообще акой пользоватлеь, если что выводим 404
	if tag.RowsAffected() == 0 {
		return ErrTaskNotFound
	}
	return nil
}
