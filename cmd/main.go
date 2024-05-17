package main

import (
	"log"

	"github.com/ruziba3vich/configurations/api"
	"github.com/ruziba3vich/configurations/internal/config"
	"github.com/ruziba3vich/configurations/internal/services"
	"github.com/ruziba3vich/configurations/internal/worker"
)

func main() {

	albumImple := worker.NewAlbum(
		config.New("localhost", "7777"),
		"github.com/ruziba3vich/configurations/internal/db/storage.yaml",
	)
	service := services.New(albumImple)

	server := api.New(api.Option{
		Storage: service,
	})

	if err := server.Run(); err != nil {
		log.Fatal("Failed to run HTTP server:", err)
		log.Fatal(err)
	}

	log.Println("server has been stopped")
}
