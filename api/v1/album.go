package api

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
