package test

import (
	"testing"
)

func TestHappyGetAlbum(t *testing.T) {
	_, err := Client.V1().GetAlbum(WellKnownAlbumID)
	if err != nil {
		t.Errorf("expected success getting well known album: %v", err)
		t.FailNow()
	}
}

func TestSadGetAlbum(t *testing.T) {
	album, err := Client.V1().GetAlbum(UnknownAlbumID)
	if err != nil {
		t.Logf("success returning error for unknown album: %v", err)
		return
	}
	t.Errorf("expected error for unknown album: %d", album.ID)
	t.FailNow()
}
