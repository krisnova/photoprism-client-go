package api

// POST /api/v1/index
func (v1 *V1Client) Index() error {
	resp := v1.POST(nil, "/api/v1/index")
	return resp.Error
}

// DELETE /api/v1/index
func (v1 *V1Client) CancelIndex() error {
	resp := v1.DELETE(nil, "/api/v1/index")
	return resp.Error
}
