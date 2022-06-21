package posts

import "medium-be/internal/entity"

type PostInterface interface {
	CreatePost(newPost entity.Posts, tags []int) error
	EditPost(idPost int, editPost entity.Posts, tags []int) error
	ReadPost(idPost int) (entity.Posts, error)
}
