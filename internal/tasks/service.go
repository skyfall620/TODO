package tasks

import "context"

type TasksService struct {
	TaskRepository ITasksRepository
}

func NewTasksService(taskRepository ITasksRepository) *TasksService {
	return &TasksService{TaskRepository: taskRepository}
}

func (t *TasksService) Create(
	ctx context.Context,
	userId, listId int64,
	title, description string) (*Task, error) {
	return t.TaskRepository.Create(ctx, &Task{
		UserID:      userId,
		ListID:      listId,
		Title:       title,
		Description: description,
	})
}

func (t *TasksService) GetAllByListID(ctx context.Context, userId, listId int64) ([]Task, error) {
	return t.TaskRepository.GetAllByListID(ctx, userId, listId)
}
