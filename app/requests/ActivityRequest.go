package requests

type (
	CreateActivity struct {
		Title string `json:"title" validate:"required"`
		Email string `json:"email" validate:"required"`
	}

	UpdateActivity struct {
		ID    int    `json:"id"`
		Title string `json:"title" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
)
