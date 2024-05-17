package worker

import (
	"errors"
	"os"
	"sort"

	"github.com/ruziba3vich/configurations/internal/config"
	"github.com/ruziba3vich/configurations/internal/models"
	"gopkg.in/yaml.v2"
)

type AlbumImple struct {
	config   *config.Config
	fileName string
}

func NewAlbum(c *config.Config, fileName string) *AlbumImple {
	return &AlbumImple{
		config:   c,
		fileName: fileName,
	}
}

// / function to get all albums
func (a *AlbumImple) GetAlbums() ([]models.Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}
	return config.Albums, nil
}

// / function to create a new album
func (a *AlbumImple) CreateAlbum(req models.CreateAlbumReq) (*models.Album, error) {
	id := a.getLastAlbumId() + 1
	album := models.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}
	album.SetId(id)
	a.config.Albums = append(a.config.Albums, album)
	if err := a.writeIntoYaml(); err != nil {
		return nil, err
	}
	return &album, nil
}

// / function to get album by artist
func (a *AlbumImple) GetAlbumByArtist(req models.GetAlbumByArtistRequest) (*models.Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}

	for _, album := range config.Albums {
		if album.Artist == req.Artist {
			return &album, nil
		}
	}

	return nil, errors.New("no album with the given artist found")
}

// / function to get an album by title
func (a *AlbumImple) GetAlbumByTitle(req models.GetAlbumByTitleRequest) (*models.Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}

	for _, album := range config.Albums {
		if album.Title == req.Title {
			return &album, nil
		}
	}

	return nil, errors.New("no album with the given title found")
}

// / function to get an album by price
func (a *AlbumImple) GetAlbumByPrice(req models.GetAlbumByPriceRequest) (*models.Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}

	for _, album := range config.Albums {
		if album.Price == req.Price {
			return &album, nil
		}
	}

	return nil, errors.New("no album with the given price found")
}

func (a *AlbumImple) GetAlbumByGivenPriceInterval(req models.GetAlbumsByGivenInterval) ([]models.Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}

	var albums []models.Album

	for _, album := range config.Albums {
		if album.Price >= req.Start && album.Price <= req.End {
			albums = append(albums, album)
		}
	}

	return albums, nil
}

/// ------------------------------------- background functions ---------------------------------------///

func (a *AlbumImple) getSortedAlbumsByID() []models.Album {
	albums := a.config.Albums
	sort.Slice(albums, func(i, j int) bool {
		return albums[i].GetId() < albums[j].GetId()
	})
	return albums
}

func (a *AlbumImple) getLastAlbumId() int {
	albums := a.getSortedAlbumsByID()
	if len(albums) > 0 {
		return albums[len(albums)-1].GetId()
	}
	return 0
}

func (a *AlbumImple) writeIntoYaml() error {
	data, err := yaml.Marshal(&a.config)
	if err != nil {
		return err
	}

	err = os.WriteFile(a.fileName, data, 0644)
	return err
}

func (a *AlbumImple) loadAllAlbums() (*config.Config, error) {
	data, err := os.ReadFile(a.fileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &a.config)
	return a.config, err
}
