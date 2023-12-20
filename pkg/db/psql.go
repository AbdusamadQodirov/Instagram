package db

import (
	"database/sql"
	"fmt"
	"instagram/config"

	_ "github.com/lib/pq"
)

func ConnectToDb(cfg *config.Config) *sql.DB {
	datasourceName := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", cfg.PSQL.User, cfg.PSQL.DBname, cfg.PSQL.Password, cfg.PSQL.SSlmode)
	db, err := sql.Open("postgres", datasourceName)
	if err != nil {
		fmt.Println("Error on connection db:", err)
		return &sql.DB{}
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error on ping:", err)
		return &sql.DB{}
	}
	fmt.Println("Connected to Db!")
	return db
}
