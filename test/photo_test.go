package test

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/kris-nova/photoprism-client-go/api/v1"
)

//TODO Test GetPhotos()

func TestHappyGetPhotos(t *testing.T) {
	_, err := Client.V1().GetPhotos(&api.PhotoOptions{
		Count:    1,
		AlbumUID: WellKnownAlbumID,
	})
	if err != nil {
		t.Errorf("expected success getting well known photo: %v", err)
		t.FailNow()
	}
}

func TestHappyGetPhoto(t *testing.T) {
	_, err := Client.V1().GetPhoto(WellKnownPhotoID)
	if err != nil {
		t.Errorf("expected success getting well known photo: %v", err)
		t.FailNow()
	}
}

func TestSadGetPhoto(t *testing.T) {
	photo, err := Client.V1().GetPhoto(UnknownPhotoID)
	if err != nil {
		t.Logf("success returning error for unknown photo: %v", err)
		return
	}
	t.Errorf("expected error for unknown photo: %s", photo.UUID)
	t.FailNow()
}

func TestHappyUpdatePhoto(t *testing.T) {
	photo, err := Client.V1().GetPhoto(WellKnownPhotoID)
	if err != nil {
		t.Errorf("error getting well known photo: %v", err)
		t.FailNow()
		return
	}
	photo.PhotoDescription = fmt.Sprintf("Sample App Description: %s", time.Now().String())
	_, err = Client.V1().UpdatePhoto(photo)
	if err != nil {
		t.Errorf("error updating photo: %v", err)
		t.FailNow()
		return
	}
}

func TestSadUpdatePhoto(t *testing.T) {
	photo, err := Client.V1().GetPhoto(WellKnownPhotoID)
	if err != nil {
		t.Errorf("error getting well known photo: %v", err)
		t.FailNow()
		return
	}
	photo.UUID = UnknownPhotoID
	photo.PhotoDescription = fmt.Sprintf("Sample App Description: %s", time.Now().String())
	photo, err = Client.V1().UpdatePhoto(photo)
	if err != nil {
		t.Logf("expecting failure at updating photo unknown photo: %v", err)
		return
	}
	t.Errorf("expecting failure updaitng bad photo id: %s", photo.UUID)
	t.FailNow()
}

func TestHappyGetPhotoDownload(t *testing.T) {
	photo, err := Client.V1().GetPhoto(WellKnownPhotoID)
	if err != nil {
		t.Errorf("error getting well known photo: %v", err)
		t.FailNow()
		return
	}
	file, err := Client.V1().GetPhotoDownload(WellKnownPhotoID)
	if err != nil {
		t.Errorf("expected success getting well known photo: %v", err)
		t.FailNow()
	}
	for _, f := range photo.Files {
		fileName := path.Base(f.FileName)
		t.Logf("Downloaded [%s]", fileName)
		t.Logf("Photo Bytes: %d", len(file))
	}
}

func TestSadGetPhotoDownload(t *testing.T) {
	_, err := Client.V1().GetPhotoDownload(UnknownPhotoID)
	if err == nil {
		t.Errorf("expected failure getting well known photo: %v", err)
		t.FailNow()
	}
}
