package api

import "fmt"

type AlbumOptions struct {
	ParamType string
	Q         string
	Count     int
	Offset    int
	Category  string
}

const (
	DefaultAlbumOptionsParamType = "album"
	DefaultAlbumOptionsQ         = ""
	DefaultAlbumOptionsCount     = 24
	DefaultAlbumOptionsOffset    = 0
	DefaultAlbumOptionsCategory  = ""
)

// GET /api/v1/albums
//
// Example Params: http://localhost:8080/api/v1/albums?count=24&offset=0&q=&category=&type=album
func (v1 *V1Client) GetAlbums(options *AlbumOptions) ([]Album, error) {
	albums := []Album{{}}

	if options == nil {

		// Default to sane options for query
		options = &AlbumOptions{
			ParamType: "album",
			Q:         "",
			Count:     24,
			Offset:    0,
			Category:  "",
		}
	}

	// Checks for missing fields
	if options.Count == 0 {
		return albums, nil
	}
	if options.ParamType == "" {
		options.ParamType = DefaultAlbumOptionsParamType
	}

	// NOTE: Even though this method is singular GetAlbum
	// if will call the "albums" plural endpoint.
	err := v1.GET("/api/v1/albums?count=%d&offset=%d&q=%s&category=%s&type=%s", options.Count, options.Offset, options.Q, options.Category, options.ParamType).JSON(&albums)
	return albums, err
}

// GET /api/v1/albums/:uuid
func (v1 *V1Client) GetAlbum(uuid string) (Album, error) {
	album := Album{}

	// NOTE: Even though this method is singular GetAlbum
	// if will call the "albums" plural endpoint.
	err := v1.GET("/api/v1/albums/%s", uuid).JSON(&album)
	return album, err
}

// CreateAlbum is used to create a new Album.
//
// CreateAlbum will default to sane values
// such that an empty Album{} object will still
// create a new album.
//
//POST /api/v1/albums
func (v1 *V1Client) CreateAlbum(object Album) (Album, error) {
	err := v1.POST(&object, "/api/v1/albums").JSON(&object)
	return object, err
}

// PUT /api/v1/albums/:uid
func (v1 *V1Client) UpdateAlbum(object Album) (Album, error) {
	if object.AlbumUID == "" {
		return object, fmt.Errorf("missing album.AlbumUID in album")
	}
	err := v1.PUT(&object, "/api/v1/albums/%s", object.AlbumUID).JSON(&object)
	return object, err
}

// POST /api/v1/batch/albums/delete
func (v1 *V1Client) DeleteAlbums(uuids []string) error {
	payload := struct {
		Albums []string `json:"albums"`
	}{
		Albums: uuids,
	}
	resp := v1.POST(payload, "/api/v1/batch/albums/delete")
	return resp.Error
}

// POST /api/v1/albums/:uid/like
//
// Parameters:
//   uid: string Album UID
func (v1 *V1Client) LikeAlbum(uuid string) error {
	resp := v1.POST(nil, "/api/v1/albums/%s/like", uuid)
	return resp.Error
}

// DELETE /api/v1/albums/:uid/like
//
// Parameters:
//   uid: string Album UID
func (v1 *V1Client) DislikeAlbum(uuid string) error {
	resp := v1.DELETE(nil, "/api/v1/albums/%s/like", uuid)
	return resp.Error
}

// POST /api/v1/albums/:uid/clone
func (v1 *V1Client) CloneAlbum(object Album) (Album, error) {
	if object.AlbumUID == "" {
		return object, fmt.Errorf("missing album.AlbumUID in album")
	}
	newAlbum := Album{}
	err := v1.POST(&object, "/api/v1/albums/%s/clone", object.AlbumUID).JSON(&newAlbum)
	return newAlbum, err
}

// POST /api/v1/albums/:uid/photos
func (v1 *V1Client) AddPhotosToAlbum(albumUUID string, photoIDs []string) error {
	payload := struct {
		Photos []string `json:"photos"`
	}{
		Photos: photoIDs,
	}
	resp := v1.POST(&payload, "/api/v1/albums/%s/photos", albumUUID)
	return resp.Error
}

// DELETE /api/v1/albums/:uid/photos
func (v1 *V1Client) DeletePhotosFromAlbum(albumUUID string, photoIDs []string) error {
	payload := struct {
		Photos []string `json:"photos"`
	}{
		Photos: photoIDs,
	}
	resp := v1.DELETE(&payload, "/api/v1/albums/%s/photos", albumUUID)
	return resp.Error
}

// GET /api/v1/albums/:uid/dl
func (v1 *V1Client) GetAlbumDownload(uuid string) ([]byte, error) {
	// NOTE: Even though this method is singular GetAlbum
	// if will call the "albums" plural endpoint.
	resp := v1.GET("/api/v1/albums/%s?t=%s", uuid, v1.downloadToken)
	return resp.Body, resp.Error
}
