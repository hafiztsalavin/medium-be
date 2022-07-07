package posts

type PostResponse struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}
