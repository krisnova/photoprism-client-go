// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package test

import (
	"testing"

	"github.com/drummonds/photoprism-client-go/api/v1"
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

// TestHappyCreateUpdateDeleteAlbum
func TestHappyCreateUpdateDeleteAlbum(t *testing.T) {
	album := api.Album{
		AlbumTitle: WellKnownAlbumTitle,
	}

	newAlbum, err := Client.V1().CreateAlbum(album)
	if err != nil {
		t.Errorf("expected success creating album: %v", err)
		t.FailNow()
	}

	newAlbum.AlbumDescription = "An updated album description"
	newAlbum, err = Client.V1().UpdateAlbum(newAlbum)
	if err != nil {
		t.Errorf("unable to update test album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	err = Client.V1().DeleteAlbums([]string{newAlbum.AlbumUID})
	if err != nil {
		t.Errorf("expected delete album %s, album not deleted: %v", newAlbum.AlbumUID, err)
		t.FailNow()
	}

}

// TestHappyLikeDislikeAlbum
func TestHappyLikeDislikeAlbum(t *testing.T) {
	album := api.Album{
		AlbumTitle: WellKnownAlbumTitle,
	}

	newAlbum, err := Client.V1().CreateAlbum(album)
	if err != nil {
		t.Errorf("expected success creating album: %v", err)
		t.FailNow()
	}

	err = Client.V1().LikeAlbum(newAlbum.AlbumUID)
	if err != nil {
		t.Errorf("expected to like album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	err = Client.V1().DislikeAlbum(newAlbum.AlbumUID)
	if err != nil {
		t.Errorf("expected to unlike album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	err = Client.V1().DeleteAlbums([]string{newAlbum.AlbumUID})
	if err != nil {
		t.Errorf("expected delete album %s, album not deleted: %v", newAlbum.AlbumUID, err)
		t.FailNow()
	}
}

// TestHappyLikeDislikeAlbum
func TestSadLikeDislikeAlbum(t *testing.T) {
	err := Client.V1().LikeAlbum(UnknownAlbumID)
	if err == nil {
		t.Errorf("expected to error during unknown like album: %v", err)
		t.FailNow()
	}

	err = Client.V1().DislikeAlbum(UnknownAlbumID)
	if err == nil {
		t.Errorf("expected to error during unknown dislike album: %v", err)
		t.FailNow()
	}
}

// CloneAlbums
// TestHappyLikeDislikeAlbum
func TestHappyCloneAlbum(t *testing.T) {
	album := api.Album{
		AlbumTitle: WellKnownAlbumTitle,
	}

	newAlbum, err := Client.V1().CreateAlbum(album)
	if err != nil {
		t.Errorf("expected success creating album: %v", err)
		t.FailNow()
	}

	clonedAlbum, err := Client.V1().CloneAlbum(newAlbum)
	if err != nil {
		t.Errorf("expected to like album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	err = Client.V1().DeleteAlbums([]string{newAlbum.AlbumUID, clonedAlbum.AlbumUID})
	if err != nil {
		t.Errorf("expected delete album %s, album not deleted: %v", newAlbum.AlbumUID, err)
		t.FailNow()
	}
}

// TestSadCloneAlbum
func TestSadCloneAlbum(t *testing.T) {
	album := api.Album{
		AlbumUID: UnknownAlbumID,
	}
	_, err := Client.V1().CloneAlbum(album)
	if err == nil {
		t.Errorf("expected to error during unknown clone album: %v", err)
		t.FailNow()
	}
}

// TestAlbumAddDeletePhoto is a giant integration test
// that will exercise many methods in the SDK
//
// This is the most complete integration test in the suite
// and is the test that will also exercise adding and deleting
// photos from an album
func TestAlbumAddDeletePhoto(t *testing.T) {
	album := api.Album{
		AlbumTitle: WellKnownAlbumTitle,
	}

	newAlbum, err := Client.V1().CreateAlbum(album)
	if err != nil {
		t.Errorf("expected success creating album: %v", err)
		t.FailNow()
	}

	// Add Photos
	photos := []string{
		WellKnownPhotoID,
	}
	err = Client.V1().AddPhotosToAlbum(newAlbum.AlbumUID, photos)
	if err != nil {
		t.Errorf("expected to add photos to album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	// Get the photos by album
	updatedPhotos, err := Client.V1().GetPhotos(&api.PhotoOptions{
		Count:    100,
		AlbumUID: newAlbum.AlbumUID,
	})
	if err != nil {
		t.Errorf("expecting to list photos by album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	var updatedPhotoIDs []string
	for _, photo := range updatedPhotos {
		updatedPhotoIDs = append(updatedPhotoIDs, photo.PhotoUID)
	}
	if len(updatedPhotos) != 1 {
		t.Errorf("expecting 1 well known photo in album, found: %d", len(updatedPhotos))
	}

	err = Client.V1().DeletePhotosFromAlbum(newAlbum.AlbumUID, updatedPhotoIDs)
	if err != nil {
		t.Errorf("expected to delete newly created photos from album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	// Get the photos by album
	updatedPhotos, err = Client.V1().GetPhotos(&api.PhotoOptions{
		Count:    100,
		AlbumUID: newAlbum.AlbumUID,
	})
	if err != nil {
		t.Errorf("expecting to list photos by album: %v", err)
		// Note: We do NOT FailNow() here because we want to clean up
	}

	if len(updatedPhotos) != 0 {
		t.Errorf("expected empty album, found %d photos", len(updatedPhotos))
		// Note: We do NOT FailNow() here because we want to clean up
	}

	err = Client.V1().DeleteAlbums([]string{newAlbum.AlbumUID})
	if err != nil {
		t.Errorf("expected delete album %s, album not deleted: %v", newAlbum.AlbumUID, err)
		t.FailNow()
	}
}

func TestHappyGetAlbumDownload(t *testing.T) {
	// GetAlbumDownload should return a .zip file
	bytes, err := Client.V1().GetAlbumDownload(WellKnownAlbumID)
	if err != nil {
		t.Errorf("expecting album download: %v", err)
		t.FailNow()
	}
	t.Logf("bytes of .zip file downloaded: %db", len(bytes))
}

func TestSadGetAlbumDownload(t *testing.T) {
	_, err := Client.V1().GetPhotoDownload(UnknownAlbumID)
	if err == nil {
		t.Errorf("expected failure getting well known album: %v", err)
		t.FailNow()
	}
}
