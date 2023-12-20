package postgres

import (
	"database/sql"
	"instagram/models"
	"instagram/storage/repo"
	"io"
	"log"
)

type PostRepo struct {
	DB *sql.DB
}

func NewPostRepo(db *sql.DB) repo.PostRepoI {
	return &PostRepo{
		DB: db,
	}
}

func (p *PostRepo) CreatePost(post models.Post) (models.Post, error) {
	query := `
	INSERT INTO posts (
		id,
		title,
		msg,
		user_id
	)
	VALUES(
		$1,
		$2,
		$3,
		$4
	);`

	_, err := p.DB.Exec(query, post.ID, post.Title, post.Msg, post.UserID)
	if err != nil {
		log.Println("Error on create post:", err)
		return models.Post{}, err
	}

	return post, nil
}
func (p *PostRepo) GetPosts() ([]models.Post, error) {
	query := `
	SELECT * FROM posts`

	rows, err := p.DB.Query(query)
	if err != nil {
		log.Println("Error on getPost on database:", err)
		return nil, err
	}
	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Msg, &post.UserID)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error on scan data to get post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (p *PostRepo) GetPostsById(id string) ([]models.Post, error) {
	var post models.Post
	var posts []models.Post

	query := `
	SELECT * FROM posts WHERE user_id = $1
	`
	rows, err := p.DB.Query(query, id)
	if err != nil {
		log.Println("Error on get posts by id:", err)
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Title, &post.Msg, &post.UserID); err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error on scan data from db:", err)
			return nil, err
		}

		posts = append(posts, post)

	}
	return posts, nil
}
func (p *PostRepo) UpdatePost(post models.Post) models.Post {
	query := `
	UPDATE posts SET title = $1, 
					 msg = $2,
					 user_id = $3 
	WHERE id = $4 `

	_, err := p.DB.Exec(query, post.Title, post.Msg, post.UserID, post.ID)
	if err != nil {
		log.Println("Error on update post:", err)
		return models.Post{}
	}

	return post
}
func (p *PostRepo) DeletePost(id string) error {
	query := `
	DELETE FROM posts WHERE id = $1`
	_, err := p.DB.Exec(query, id)
	if err != nil {
		log.Println("Error on deleting post:", err)
		return err
	}
	return nil
}
