package tasks

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
