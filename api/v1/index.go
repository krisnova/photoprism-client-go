package api

// Index is used to sync the backend storage with
// the database meta information
//
// POST /api/v1/index
func (v1 *V1Client) Index() error {
	resp := v1.POST(nil, "/api/v1/index")
	return resp.Error
}

// CancelIndex can be used to attempt to cancel a running index
// operation
//
// DELETE /api/v1/index
func (v1 *V1Client) CancelIndex() error {
	resp := v1.DELETE(nil, "/api/v1/index")
	return resp.Error
}
