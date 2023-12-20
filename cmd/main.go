package main

import (
	"instagram/api"
	"instagram/api/handler"
	"instagram/config"
	"instagram/pkg/db"
	"instagram/storage"
)

func main() {
	cfg := config.NewConfig()
	db := db.ConnectToDb(cfg)

	storage1 := storage.NewStorage(db)
	handler := handler.NewHandler(storage1, cfg)

	api.Api(handler)
}
