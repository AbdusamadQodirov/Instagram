package storage

import (
	"database/sql"
	"instagram/storage/postgres"
	"instagram/storage/repo"
)

type StorageI interface {
	GetUserRepo() repo.UserRepoI
	GetPostRepo() repo.PostRepoI
	GetCommentRepo() repo.CommentRepoI
}

type Storage struct {
	UserRepo    repo.UserRepoI
	CommentRepo repo.CommentRepoI
	PostRepo    repo.PostRepoI
}

func NewStorage(db *sql.DB) StorageI {
	return &Storage{
		UserRepo: postgres.NewUserRepo(db),
		CommentRepo: postgres.NewCommentRepo(db),
		PostRepo: postgres.NewPostRepo(db),
	}
}

func (s *Storage) GetUserRepo() repo.UserRepoI {
	return s.UserRepo
}
func (s *Storage) GetPostRepo() repo.PostRepoI {
	return s.PostRepo
}
func (s *Storage) GetCommentRepo() repo.CommentRepoI {
	return s.CommentRepo
}
