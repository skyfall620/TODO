package lists

import "context"

type ListsService struct {
	ListsRepository IListsRepository
}

func NewListsService(listsRepository IListsRepository) *ListsService {
	return &ListsService{ListsRepository: listsRepository}
}

func (l *ListsService) Create(ctx context.Context, userID int64, title string) (*List, error) {
	return l.ListsRepository.Create(ctx, &List{
		UserID: userID,
		Title:  title,
	})
}

func (l *ListsService) GetAllByUserID(ctx context.Context, userID int64) ([]List, error) {
	return l.ListsRepository.GetAllByUserID(ctx, userID)
}
