package posts

import "medium-be/internal/entity"

type PostInterface interface {
	CreatePost(newPost entity.Posts, tags []int) error
}
