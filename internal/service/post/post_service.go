package post

import (
	"encoding/json"
	"fmt"
	"medium-be/internal/entity"
	_postRepo "medium-be/internal/repository/postgres/posts"

	_redisRepo "medium-be/internal/repository/redis"
)

type PostService interface {
	CreatePost(newPost entity.Posts, tags []int) error
	EditPost(idPost int, editPost entity.Posts, tags []int) error
	ReadPost(idPost, idUser int) (entity.Posts, error)
	ReadAllPost(idPost int, statusPost string) ([]entity.Posts, error)
	PublishPost(idPost, idUser int) error
	DeletePost(idPost, idUser int) error

	AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error)
	ReadPostByPostId(idPost int) (entity.Posts, error)
}

const (
	postEntity = "post"
)

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
	ps.redisRepo.DeleteCache(postEntity)

	return err
}

func (ps *postService) EditPost(idPost int, editPost entity.Posts, tags []int) error {
	err := ps.postRepo.EditPost(idPost, editPost, tags)
	ps.redisRepo.DeleteCache(postEntity)

	return err
}

func (ps *postService) ReadPost(idPost, idUser int) (entity.Posts, error) {
	var post entity.Posts

	key := ps.keyPrefix(int(post.ID), int(post.UserID))
	postCached, _ := ps.redisRepo.GetCache(key)
	if err := json.Unmarshal([]byte(postCached), &post); err == nil {
		return post, nil
	}

	post, err := ps.postRepo.ReadPost(idPost)
	if err == nil && idUser == int(post.UserID) {
		postString, _ := json.Marshal(&post)
		ps.redisRepo.CreateCache(key, postString, 0)
	}
	return post, err
}

func (ps *postService) ReadAllPost(idPost int, statusPost string) ([]entity.Posts, error) {
	posts, err := ps.postRepo.ReadAllPost(idPost, statusPost)
	ps.redisRepo.DeleteCache(postEntity)

	return posts, err
}

func (ps *postService) PublishPost(idPost, idUser int) error {
	err := ps.postRepo.PublishPost(idPost, idUser)
	ps.redisRepo.DeleteCache(postEntity)

	return err
}

func (ps *postService) DeletePost(idPost, idUser int) error {
	err := ps.postRepo.DeletePost(idPost, idUser)
	ps.redisRepo.DeleteCache(postEntity)

	return err
}

func (ps *postService) AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error) {
	var posts []entity.Posts
	key := ps.keyAllPost(filter)

	postCached, _ := ps.redisRepo.GetCache(key)
	if err := json.Unmarshal([]byte(postCached), &posts); err == nil {
		fmt.Println("dari sini")
		return posts, nil
	}
	posts, err := ps.postRepo.AllPostPublish(filter)

	if err == nil {
		postsString, _ := json.Marshal(&posts)
		ps.redisRepo.CreateCache(key, postsString, 0)
	}

	return posts, err
}

func (ps *postService) ReadPostByPostId(idPost int) (entity.Posts, error) {
	var post entity.Posts
	key := ps.keyPrefix(idPost, 0)

	postCached, _ := ps.redisRepo.GetCache(key)
	if err := json.Unmarshal([]byte(postCached), &post); err == nil {
		return post, nil
	}

	post, err := ps.postRepo.ReadPostByPostId(idPost)
	if err == nil {
		postString, _ := json.Marshal(&post)
		ps.redisRepo.CreateCache(key, postString, 0)
	}
	return post, err
}

func (ts *postService) keyPrefix(postID, userID int) string {
	return fmt.Sprintf("%s:%d:%d", postEntity, userID, postID)
}

func (ts *postService) keyAllPost(filter interface{}) string {
	key := postEntity + ":" + fmt.Sprint(filter)
	return key
}
