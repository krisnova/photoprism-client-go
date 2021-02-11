package test

import (
	"testing"

	"github.com/kris-nova/client-go/api/v1"
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

func TestHappyGetAlbumsOptionsCount1(t *testing.T) {
	options := api.AlbumOptions{
		Count: 1,
	}
	albums, err := Client.V1().GetAlbums(&options)
	if err != nil {
		t.Errorf("expected success listing 1 album: %v", err)
		t.FailNow()
	}
	if len(albums) != 1 {
		t.Errorf("expected 1 album length, got: %d", len(albums))
		t.FailNow()
	}
}

func TestHappyGetAlbumsNil(t *testing.T) {
	albums, err := Client.V1().GetAlbums(nil)
	if err != nil {
		t.Errorf("expected success listing albums: %v", err)
		t.FailNow()
	}
	t.Logf("Listed %d albums", len(albums))
}

func TestSadGetAlbums(t *testing.T) {
	options := api.AlbumOptions{
		Category: UnknownCategory,
	}
	albums, err := Client.V1().GetAlbums(&options)
	if err != nil {
		t.Errorf("error listing albums: %v", err)
		t.FailNow()
		return
	}

	// Note: by defualt we return "{}" which counts as 1 album
	if len(albums) != 1 {
		t.Errorf("Non zero length of albums")
		t.FailNow()
	}
}
