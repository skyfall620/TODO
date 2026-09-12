package tasks

import (
	"context"
	"database/sql"
)

type TaskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *sql.DB) *TaskRepository {
	return &TaskRepository{database: database}
}

func (repo *TaskRepository) Create(ctx context.Context, task *Task) (*Task, error) {
	query := `
		INSERT INTO tasks (user_id, list_id, title, description, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	created := &Task{
		UserID:      task.UserID,
		ListID:      task.ListID,
		Title:       task.Title,
		Description: task.Description,
		Status:      "open",
	}

	err := repo.database.QueryRowContext(
		ctx,
		query,
		created.UserID,
		created.ListID,
		created.Title,
		created.Description,
		created.Status,
	).Scan(&created.ID, &created.CreatedAt, &created.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (repo *TaskRepository) GetAllByListID(ctx context.Context, userID int64, listID int64) ([]Task, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE user_id = $1 AND list_id = $2
	`

	rows, err := repo.database.QueryContext(ctx, query, userID, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}
