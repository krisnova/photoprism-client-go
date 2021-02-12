package api

import "fmt"

// AlbumOptions are the parameters passed to get
// albums by various fields. Populate these as needed
// to pass to the SDK
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

// GetAlbums is used to list albums by query fields.
//
// GET /api/v1/albums
//
// Example Params: http://localhost:8080/api/v1/albums?count=24&offset=0&q=&category=&type=album
func (v1 *V1Client) GetAlbums(options *AlbumOptions) ([]Album, error) {
	albums := []Album{{}}

	if options == nil {

		// Default to sane options for query
		options = &AlbumOptions{
			ParamType: DefaultAlbumOptionsParamType,
			Q:         DefaultAlbumOptionsQ,
			Count:     DefaultAlbumOptionsCount,
			Offset:    DefaultAlbumOptionsOffset,
			Category:  DefaultAlbumOptionsCategory,
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

// GetAlbum is used to get an album by an UUID.
//
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
// POST /api/v1/albums
func (v1 *V1Client) CreateAlbum(album Album) (Album, error) {
	err := v1.POST(&album, "/api/v1/albums").JSON(&album)
	return album, err
}

// UpdateAlbum will update meta information about an album.
//
// PUT /api/v1/albums/:uid
func (v1 *V1Client) UpdateAlbum(album Album) (Album, error) {
	if album.AlbumUID == "" {
		return album, fmt.Errorf("missing album.AlbumUID in album")
	}
	err := v1.PUT(&album, "/api/v1/albums/%s", album.AlbumUID).JSON(&album)
	return album, err
}

// DeleteAlbums will batch delete a set of albums by ID.
//
// POST /api/v1/batch/albums/delete
func (v1 *V1Client) DeleteAlbums(albumUUIDs []string) error {
	payload := struct {
		Albums []string `json:"albums"`
	}{
		Albums: albumUUIDs,
	}
	resp := v1.POST(payload, "/api/v1/batch/albums/delete")
	return resp.Error
}

// LikeAlbum can be used to like an album.
//
// POST /api/v1/albums/:uid/like
//
// Parameters:
//   uid: string Album UID
func (v1 *V1Client) LikeAlbum(uuid string) error {
	resp := v1.POST(nil, "/api/v1/albums/%s/like", uuid)
	return resp.Error
}

// DislikeAlbum can be used to dislike an album.
//
// DELETE /api/v1/albums/:uid/like
//
// Parameters:
//   uid: string Album UID
func (v1 *V1Client) DislikeAlbum(uuid string) error {
	resp := v1.DELETE(nil, "/api/v1/albums/%s/like", uuid)
	return resp.Error
}

// CloneAlbum can be used to clone an album and will
// return the newly cloned album on success.
//
// POST /api/v1/albums/:uid/clone
func (v1 *V1Client) CloneAlbum(album Album) (Album, error) {
	if album.AlbumUID == "" {
		return album, fmt.Errorf("missing album.AlbumUID in album")
	}
	newAlbum := Album{}
	err := v1.POST(&album, "/api/v1/albums/%s/clone", album.AlbumUID).JSON(&newAlbum)
	return newAlbum, err
}

// AddPhotosToAlbum will associate a set of photos by UUID with an album by UUID
//
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

// DeletePhotosFromAlbum will disassociate a set of photos by UUID from an album by UUID
//
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

// GetAlbumDownload will return a .zip file of the album's content
// and can be used to download an album from the API.
//
// GET /api/v1/albums/:uid/dl
func (v1 *V1Client) GetAlbumDownload(uuid string) ([]byte, error) {
	// NOTE: Even though this method is singular GetAlbum
	// if will call the "albums" plural endpoint.
	resp := v1.GET("/api/v1/albums/%s?t=%s", uuid, v1.downloadToken)
	return resp.Body, resp.Error
}
