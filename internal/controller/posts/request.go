package posts

type PostRequest struct {
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
	Status string `json:"status"`
	Tags   []int  `json:"tags" validate:"required"`
}
