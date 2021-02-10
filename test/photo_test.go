package test

import (
	"fmt"
	"testing"
	"time"
)

func TestHappyGetPhoto(t *testing.T) {
	_, err := Client.V1().GetPhoto(WellKnownPhotoID)
	if err != nil {
		t.Errorf("expected success getting well known photo: %v", err)
		t.FailNow()
	}
}

func TestSadGetPhoto(t *testing.T) {
	photo, err := Client.V1().GetPhoto("1234567890")
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
	photo.UUID = "1234567890"
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
	_, err := Client.V1().GetPhotoDownload(WellKnownPhotoID)
	if err != nil {
		t.Errorf("expected success getting well known photo: %v", err)
		t.FailNow()
	}
}

func TestSadGetPhotoDownload(t *testing.T) {
	file, err := Client.V1().GetPhotoDownload("1234567890")
	if err != nil {
		t.Logf("success returning error for unknown photo: %v", err)
		return
	}
	t.Errorf("expected error for unknown file: %s", file.FileName)
	t.FailNow()
}
