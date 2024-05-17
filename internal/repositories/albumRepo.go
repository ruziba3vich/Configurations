package repositories

import "github.com/ruziba3vich/configurations/internal/models"

type IAlbumRepo interface {
	CreateAlbum(models.CreateAlbumReq) (*models.Album, error)
	GetAlbumByArtist(models.GetAlbumByArtistRequest) (*models.Album, error)
	GetAlbumByTitle(models.GetAlbumByTitleRequest) (*models.Album, error)
	GetAlbumByPrice(models.GetAlbumByPriceRequest) (*models.Album, error)
	GetAlbumByGivenPriceInterval(models.GetAlbumsByGivenInterval) ([]models.Album, error)
}
