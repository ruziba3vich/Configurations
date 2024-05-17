package services

import (
	"github.com/ruziba3vich/configurations/internal/models"
	"github.com/ruziba3vich/configurations/internal/repositories"
)

type Service struct {
	repository repositories.IAlbumRepo
}

func New(repository repositories.IAlbumRepo) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateAlbum(req models.CreateAlbumReq) (*models.Album, error) {
	return s.repository.CreateAlbum(req)
}

func (s *Service) GetAlbumByArtist(req models.GetAlbumByArtistRequest) (*models.Album, error) {
	return s.repository.GetAlbumByArtist(req)
}

func (s *Service) GetAlbumByTitle(req models.GetAlbumByTitleRequest) (*models.Album, error) {
	return s.repository.GetAlbumByTitle(req)
}

func (s *Service) GetAlbumByPrice(req models.GetAlbumByPriceRequest) (*models.Album, error) {
	return s.repository.GetAlbumByPrice(req)
}

func (s *Service) GetAlbumByGivenPriceInterval(req models.GetAlbumsByGivenInterval) ([]models.Album, error) {
	return s.repository.GetAlbumByGivenPriceInterval(req)
}
