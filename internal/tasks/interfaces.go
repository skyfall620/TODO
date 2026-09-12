package tasks

import (
	"context"
)

type ITasksRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	GetAllByListID(ctx context.Context, userID int64, listID int64) ([]Task, error)
}
