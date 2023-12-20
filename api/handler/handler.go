package handler

import (
	"instagram/config"
	"instagram/storage"
)

type Handler struct {
	Storage storage.StorageI
	Cfg *config.Config
}

func NewHandler(storage storage.StorageI, cfg *config.Config) Handler {
	return Handler{
		Storage: storage,
		Cfg: cfg,
	}
}
