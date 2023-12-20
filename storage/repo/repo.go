package repo

import "instagram/models"

type UserRepoI interface {
	CreateUser(models.User) (models.User, error)
	GetUsers(limit, page int) ([]models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUser(models.User) (models.User)
	DeleteUserById(id string) error
}

type PostRepoI interface {
	CreatePost(models.Post) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPostsById(id string) ([]models.Post, error)
	UpdatePost(models.Post) (models.Post)
	DeletePost(id string) error
}

type CommentRepoI interface {
	CreateComment(models.Comment) (models.Comment, error)
	GetCommentByPost(post_id string) ([]models.Comment, error)
	UpdateComment(models.Comment) (models.Comment)
	DeleteComment(id string) error
}