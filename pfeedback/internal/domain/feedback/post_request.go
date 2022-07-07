package domain

type PostPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	CategoryID  int64  `json:"categoryId" validate:"required"`
}

type PostUpdatePayload struct {
	PostPayload
	Status string `json:"status"`
}
