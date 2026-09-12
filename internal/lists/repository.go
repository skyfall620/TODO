package lists

import (
	"context"
	"database/sql"
)

type ListRepository struct {
	database *sql.DB
}

func NewListRepository(database *sql.DB) *ListRepository {
	return &ListRepository{database: database}
}

func (repo *ListRepository) Create(ctx context.Context, list *List) (*List, error) {
	query := `
		INSERT INTO lists (user_id, title)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	created := &List{
		UserID: list.UserID,
		Title:  list.Title,
	}

	err := repo.database.QueryRowContext(
		ctx,
		query,
		created.UserID,
		created.Title,
	).Scan(&created.ID, &created.CreatedAt)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (repo *ListRepository) GetAllByUserID(ctx context.Context, userID int64) ([]List, error) {
	query := `
		SELECT id, title
		FROM lists
		WHERE user_id = $1
	`

	rows, err := repo.database.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := make([]List, 0)

	for rows.Next() {
		var list List

		err := rows.Scan(&list.ID, &list.Title)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}
