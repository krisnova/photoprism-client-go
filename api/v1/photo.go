package api



// GET /api/v1/photos/:uuid
//
// Parameters:
//   uuid: string PhotoUID as returned by the API
func (v1 *V1Client) GetPhoto(uuid string) (Photo, error) {
	object := Photo{
		UUID: uuid,
	}
	err := v1.GET("/api/v1/photos/%s", uuid).JSON(&object)
	return object, err
}

// PUT /api/v1/photos/:uid
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func (v1 *V1Client) UpdatePhoto(object Photo) (Photo, error) {
	err := v1.PUT(&object, "/api/v1/photos/%s", object.UUID).JSON(&object)
	return object, err
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
