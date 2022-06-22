package posts

import "medium-be/internal/entity"

type PostInterface interface {
	CreatePost(newPost entity.Posts, tags []int) error
	EditPost(idPost int, editPost entity.Posts, tags []int) error
	ReadPost(idPost int) (entity.Posts, error)
	ReadAllPost(idPost int, statusPost string) ([]entity.Posts, error)
	PublishPost(idPost, idUser int) error
	DeletePost(idPost, idUser int) error

	AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error)
	ReadPostByUserId(idUser int) ([]entity.Posts, error)
}
