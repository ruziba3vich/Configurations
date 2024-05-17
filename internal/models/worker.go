package models

import (
	"database/sql"
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

type AlbumImple struct {
	db       *sql.DB
	config   *Config
	fileName string
}

func NewAlbum(db *sql.DB, c *Config, fileName string) *AlbumImple {
	return &AlbumImple{
		db:       db,
		config:   c,
		fileName: fileName,
	}
}

// / function to get all albums
func (a *AlbumImple) GetAlbums() ([]Album, error) {
	config, err := a.loadAllAlbums()
	if err != nil {
		return nil, err
	}
	return config.Albums, nil
}

// / function to create a new album
func (a *AlbumImple) CreateAlbum(req CreateAlbumReq) (*Album, error) {
	id := a.getLastAlbumId() + 1
	album := Album{
		id:     id,
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}
	a.config.Albums = append(a.config.Albums, album)
	if err := a.writeIntoYaml(); err != nil {
		return nil, err
	}
	return &album, nil
}

/// ------------------------------------- background functions ---------------------------------------///

func (a *AlbumImple) getSortedAlbumsByID() []Album {
	albums := a.config.Albums
	sort.Slice(albums, func(i, j int) bool {
		return albums[i].id < albums[j].id
	})
	return albums
}

func (a *AlbumImple) getLastAlbumId() int {
	albums := a.getSortedAlbumsByID()
	if len(albums) > 0 {
		return albums[len(albums)-1].id
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

func (a *AlbumImple) loadAllAlbums() (*Config, error) {
	data, err := os.ReadFile(a.fileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &a.config)
	return a.config, err
}
