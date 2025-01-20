package app

import (
	"golang-demo/internal/config"
	"golang-demo/internal/database"
)

func Run() error {
	cfg, err := config.LoadConfig()

	if err != nil {
		return err
	}

	db, err := database.Connect(cfg.Database)

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
