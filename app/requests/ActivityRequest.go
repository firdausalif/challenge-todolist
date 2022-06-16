package requests

type (
	CreateActivity struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}

	UpdateActivity struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Email string `json:"email"`
	}
)
