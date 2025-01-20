package app

import (
	"golang-demo/internal/config"
	"golang-demo/internal/database"
	"golang-demo/internal/service"
	"golang-demo/internal/transport/http"
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

	bookService := service.NewBookService(db)

	router := http.SetupRouter(bookService)

	return router.Run("localhost:8080")
}
