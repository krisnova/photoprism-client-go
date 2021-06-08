//  Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
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
	photo.PhotoUID = UnknownPhotoID
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
