package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/configurations/internal/models"
	"github.com/ruziba3vich/configurations/internal/repositories"
)

type handler struct {
	service repositories.IAlbumRepo
}

type HandlerConfig struct {
	Service repositories.IAlbumRepo
}

func New(c *HandlerConfig) *handler {
	return &handler{
		service: c.Service,
	}
}

func (h *handler) CreateAlbumHandler(c *gin.Context) {
	var req models.CreateAlbumReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := h.service.CreateAlbum(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *handler) GetAlbumByArtistHandler(c *gin.Context) {
	artist := c.Query("artist")
	req := models.GetAlbumByArtistRequest{
		Artist: artist,
	}

	album, err := h.service.GetAlbumByArtist(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, album)
	}
}

func (h *handler) GetAlbumByPriceHandler(c *gin.Context) {
	price, err := strconv.ParseFloat((c.Query("price")), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	req := models.GetAlbumByPriceRequest{
		Price: float64(price),
	}

	album, err := h.service.GetAlbumByPrice(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, album)
	}
}

func (h *handler) GetAlbumByTitleHandler(c *gin.Context) {
	title := c.Query("title")
	req := models.GetAlbumByTitleRequest{
		Title: title,
	}

	album, err := h.service.GetAlbumByTitle(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, album)
	}
}

func (h *handler) GetAlbumByGivenPriceIntervalHandler(c *gin.Context) {
	var req models.GetAlbumsByGivenInterval
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	albums, err := h.service.GetAlbumByGivenPriceInterval(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, albums)
	}
}
