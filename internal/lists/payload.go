package lists

type BodyRequest struct {
	Title string `json:"title" validate:"required"`
}
