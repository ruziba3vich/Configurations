package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/configurations/api/handlers"
	"github.com/ruziba3vich/configurations/internal/repositories"
)

type Option struct {
	Storage repositories.IAlbumRepo
}

func New(option Option) *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handlers.New(&handlers.HandlerConfig{
		Service: option.Storage,
	})

	router.POST("/create-album", handler.CreateAlbumHandler)
	router.GET("/get-album-by-title/:title", handler.GetAlbumByTitleHandler)
	router.GET("/get-album-by-artist/:artist", handler.GetAlbumByArtistHandler)
	router.GET("/get-album-by-price/:price", handler.GetAlbumByPriceHandler)
	router.POST("/get-album-by-given-interval", handler.GetAlbumByGivenPriceIntervalHandler)

	return router
}
