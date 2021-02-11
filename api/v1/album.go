package api

// GET /api/v1/albums
func (v1 *V1Client) GetAlbum(uuid string) (Album, error) {
	album := Album{}

	// NOTE: Even though this method is singular GetAlbum
	// if will call the "albums" plural endpoint.
	err := v1.GET("/api/v1/albums/%s", uuid).JSON(&album)
	return album, err
}