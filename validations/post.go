package validations

type (
	PostValidation struct {
		Title    string `json:"title" validate:"required"`
		Content  string `json:"content" validate:"required"`
		Slug     string `json:"slug" validate:"required"`
		Status   string `json:"status"`
		ImageURL string `json:"image_url"`
	}
)
