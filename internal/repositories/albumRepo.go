package repositories

import "context"

type IAlbumRepo interface {
	CreateAlbum(context.Context) error
}
