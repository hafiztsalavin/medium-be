package post

import (
	"medium-be/internal/entity"
	_postRepo "medium-be/internal/repository/postgres/posts"

	_redisRepo "medium-be/internal/repository/redis"
)

type PostService interface {
	CreatePost(newPost entity.Posts, tags []int) error
	EditPost(idPost int, editPost entity.Posts, tags []int) error
	ReadPost(idPost int) (entity.Posts, error)
	ReadAllPost(idPost int, statusPost string) ([]entity.Posts, error)
	PublishPost(idPost, idUser int) error
	DeletePost(idPost, idUser int) error

	AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error)
	ReadPostByPostId(idPost int) (entity.Posts, error)
}

type postService struct {
	postRepo  _postRepo.PostRepository
	redisRepo _redisRepo.RedisRepository
}

func NewPostService(postRepo _postRepo.PostRepository, redisRepo _redisRepo.RedisRepository) PostService {
	return &postService{
		postRepo:  postRepo,
		redisRepo: redisRepo,
	}
}

func (ps *postService) CreatePost(newPost entity.Posts, tags []int) error {
	err := ps.postRepo.CreatePost(newPost, tags)

	return err
}

func (ps *postService) EditPost(idPost int, editPost entity.Posts, tags []int) error {
	err := ps.postRepo.EditPost(idPost, editPost, tags)
	return err
}

func (ps *postService) ReadPost(idPost int) (entity.Posts, error) {

	post, err := ps.postRepo.ReadPost(idPost)
	return post, err
}

func (ps *postService) ReadAllPost(idPost int, statusPost string) ([]entity.Posts, error) {
	posts, err := ps.postRepo.ReadAllPost(idPost, statusPost)
	return posts, err
}

func (ps *postService) PublishPost(idPost, idUser int) error {
	err := ps.postRepo.PublishPost(idPost, idUser)
	return err
}

func (ps *postService) DeletePost(idPost, idUser int) error {
	err := ps.postRepo.DeletePost(idPost, idUser)
	return err
}

func (ps *postService) AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error) {
	posts, err := ps.postRepo.AllPostPublish(filter)
	return posts, err
}

func (ps *postService) ReadPostByPostId(idPost int) (entity.Posts, error) {
	post, err := ps.postRepo.ReadPostByPostId(idPost)
	return post, err
}
