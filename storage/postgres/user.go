package postgres

import (
	"database/sql"
	"instagram/models"
	"instagram/storage/repo"
	"log"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepoI {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) CreateUser(newUser models.User) (models.User, error) {
	query := `
	INSERT INTO users(
		id,
		email,
		fullname,
		username,
		passwordd		
	)
	VALUES (
		$1, $2, $3, $4, $5
	);`
	_, err := u.DB.Exec(query, newUser.ID, newUser.Email, newUser.Fullname, newUser.Username, newUser.Password)

	if err != nil {
		log.Println("Error on insert data to db:", err)
		return models.User{}, err
	}

	return newUser, nil
}
func (u *UserRepo) GetUsers(limit, page int) ([]models.User, error) {
	offset := (page - 1) * limit
	query := `
	SELECT * from users LIMIT $1 OFFSET $2`
	rows, err := u.DB.Query(query, limit, offset)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var user models.User
	var users []models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Fullname, &user.Username, &user.Password)
		if err != nil {
			log.Println("Error on scan data:", err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
func (u *UserRepo) GetUserById(id string) (models.User, error) {
	query := `
	SELECT * FROM users WHERE id = $1`
	row := u.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Fullname, &user.Username, &user.Password)

	if err != nil {
		log.Println("Error on scan user:", err)
		return models.User{}, err
	}

	return user, nil
}
func (u *UserRepo) UpdateUser(user models.User) models.User {
	query := `
	UPDATE users SET email = $1,
					 fullname = $2,
					 username = $3,
					 passwordd = $4 WHERE id = $5`
	_, err := u.DB.Exec(query, user.Email, user.Fullname, user.Username, user.Password, user.ID)

	if err != nil {
		log.Println("Error on Exec:", err)
		return models.User{}
	}
	return user
}
func (u *UserRepo) DeleteUserById(id string) error {
	query := `
	DELETE FROM users WHERE id = $1`

	if _, err := u.DB.Exec(query, id); err != nil {
		log.Println("Error on exec:", err)
		return err
	}

	return nil
}
