package models

type GetAlbumByTitleRequest struct {
	Title string `json:"title"`
}

type GetAlbumByArtistRequest struct {
	Artist string `json:"artist"`
}

type GetAlbumByPriceRequest struct {
	Price float64 `json:"price"`
}

type CreateAlbumReq struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
