package postgres

import (
	"database/sql"
	"instagram/models"
	"instagram/storage/repo"
	"io"
	"log"
)

type CommentRepo struct {
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) repo.CommentRepoI {
	return &CommentRepo{
		DB: db,
	}
}

func (c *CommentRepo) CreateComment(comment models.Comment) (models.Comment, error) {
	query := `
	INSERT INTO comments(
		id, 
		title,
		msg,
		user_id,
		post_id
	)
	VALUES(
		$1,
		$2,
		$3,
		$4,
		$5
	)`
	_, err := c.DB.Exec(query, comment.ID, comment.Title, comment.Msg, comment.UserID, comment.PostID)
	if err != nil {
		log.Println("Error on creating comment:", err)
		return models.Comment{}, err
	}
	return comment, nil
}
func (c *CommentRepo) GetCommentByPost(post_id string) ([]models.Comment, error) {
	var comment models.Comment
	var comments []models.Comment

	query := `
	SELECT * FROM comments WHERE post_id = $1`

	rows, err := c.DB.Query(query, post_id)
	if err != nil {
		log.Println("Error on querying data:", err)
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&comment.ID, &comment.Title, &comment.Msg, &comment.UserID, &comment.PostID)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error on scan comment from database to get comment by postid:", err)
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
func (c *CommentRepo) UpdateComment(comment models.Comment) models.Comment {
	query := `
	UPDATE comments SET title = $1,
						msg = $2,
						user_id = $3,
						post_id = $4
	WHERE id = $5`

	_, err := c.DB.Exec(query, comment.Title, comment.Msg, comment.UserID, comment.PostID, comment.ID)
	if err != nil {
		log.Println("Error on update comment:", err)
		return models.Comment{}
	}

	return comment
}
func (c *CommentRepo) DeleteComment(id string) error {
	query := `
	DELETE FROM comments WHERE id = $1`

	_, err := c.DB.Exec(query, id); 
	if err != nil {
		log.Println("Error on deleting comment:", err)
		return err
	}
	
	return nil
}
