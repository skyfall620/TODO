package lists

import "context"

type IListsRepository interface {
	Create(ctx context.Context, list *List) (*List, error)
	GetAllByUserID(ctx context.Context, userID int64) ([]List, error)
}
