package api

// GetPhoto can be used to get a photo by UUID
//
//GET /api/v1/photos/:uuid
//
// Parameters:
//   uuid: string PhotoUID as returned by the API
func (v1 *V1Client) GetPhoto(uuid string) (Photo, error) {
	photo := Photo{
		UUID:     uuid,
		PhotoUID: uuid,
	}
	err := v1.GET("/api/v1/photos/%s", uuid).JSON(&photo)
	return photo, err
}

// PhotoOptions is used while listing photos. These
// fields can be optionally set to query for specific
// photos.
type PhotoOptions struct {
	Count    int
	Offset   int
	AlbumUID string
	Filter   string
	Merged   bool
	Country  string
	Camera   int
	Order    string
	Q        string
}

const (
	DefaultPhotoOptionsCount  = 60
	DefaultPhotoOptionsOffset = 0
	DefaultPhotoOptionsMerged = true
	DefaultPhotoOptionsCamera = 0
	DefaultPhotoOptionsOrder  = "oldest"
)

// GET /api/v1/photos/
//
// http://localhost:8080/api/v1/photos?
// count=60&offset=0&album=aqoe4m9204aigugh&filter=&merged=true&country=&camera=0&order=oldest&q=
func (v1 *V1Client) GetPhotos(options *PhotoOptions) ([]Photo, error) {
	var photos []Photo
	if options == nil {
		options = &PhotoOptions{
			Count:  DefaultPhotoOptionsCount,
			Offset: DefaultPhotoOptionsOffset,
			Merged: DefaultPhotoOptionsMerged,
			Order:  DefaultPhotoOptionsOrder,
			Camera: DefaultPhotoOptionsCamera,
		}
	}
	if options.Count == 0 {
		return photos, nil
	}
	if options.Order == "" {
		options.Order = DefaultPhotoOptionsOrder
	}
	err := v1.GET("/api/v1/photos?count=%d&offset=%d&album=%s&filter=%s&merged=%t&country=%s&camera=%d&order=%s&q=%s",
		options.Count, options.Offset, options.AlbumUID, options.Filter, options.Merged, options.Country, options.Camera, options.Order, options.Q).JSON(&photos)
	return photos, err
}

// PUT /api/v1/photos/:uid
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) UpdatePhoto(photo Photo) (Photo, error) {
	err := v1.PUT(&photo, "/api/v1/photos/%s", photo.PhotoUID).JSON(&photo)
	return photo, err
}

// GET /api/v1/photos/:uuid/dl
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) GetPhotoDownload(uuid string) ([]byte, error) {
	resp := v1.GET("/api/v1/photos/%s/dl?t=%s", uuid, v1.downloadToken)
	return resp.Body, resp.Error
}

// GET /api/v1/photos/:uuid/yaml
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) GetPhotoYaml(uuid string) ([]byte, error) {
	resp := v1.GET("/api/v1/photos/%s/yaml", uuid)
	return resp.Body, resp.Error
}

// POST /api/v1/photos/:uuid/approve
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) ApprovePhoto(uuid string) error {
	resp := v1.POST(nil, "/api/v1/photos/%s/approve", uuid)
	return resp.Error
}

// POST /api/v1/photos/:uid/like
//
// Parameters:
//   uid: string PhotoUID as returned by the API
func (v1 *V1Client) LikePhoto(uuid string) error {
	resp := v1.POST(nil, "/api/v1/photos/%s/like", uuid)
	return resp.Error
}

// DELETE /api/v1/photos/:uuid/like
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) DislikePhoto(uuid string) error {
	resp := v1.DELETE(nil, "/api/v1/photos/%s/approve", uuid)
	return resp.Error
}

// POST /api/v1/photos/:uid/files/:file_uid/primary
//
// Parameters:
//   uid: string PhotoUID as returned by the API
//   file_uid: string File UID as returned by the API
func (v1 *V1Client) PhotoPrimary(uuid, fileuuid string) error {
	resp := v1.POST(nil, "/api/v1/photos/%s/files/%s/primary", uuid, fileuuid)
	return resp.Error
}
