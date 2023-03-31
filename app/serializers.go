package app

type (
	ExampleSerializer struct {
		ID      uint   `json:"id"`
		Content string `json:"content" binding:"required"`
	}
)
