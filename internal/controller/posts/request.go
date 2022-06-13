package posts

type PostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
	Tags  []int  `json:"tags" validate:"required"`
}
