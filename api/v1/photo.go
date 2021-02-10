package api

import (
	"time"
)

// Photo represents a photo, all its properties, and link to all its images and sidecar files.
type Photo struct {
	Meta
	ID               uint      `gorm:"primary_key" yaml:"-"`
	UUID             string    `gorm:"type:VARBINARY(42);index;" json:"DocumentID,omitempty" yaml:"DocumentID,omitempty"`
	TakenAt          time.Time `gorm:"type:datetime;index:idx_photos_taken_uid;" json:"TakenAt" yaml:"TakenAt"`
	TakenAtLocal     time.Time `gorm:"type:datetime;" yaml:"-"`
	TakenSrc         string    `gorm:"type:VARBINARY(8);" json:"TakenSrc" yaml:"TakenSrc,omitempty"`
	PhotoUID         string    `gorm:"type:VARBINARY(42);unique_index;index:idx_photos_taken_uid;" json:"UID" yaml:"UID"`
	PhotoType        string    `gorm:"type:VARBINARY(8);default:'image';" json:"Type" yaml:"Type"`
	TypeSrc          string    `gorm:"type:VARBINARY(8);" json:"TypeSrc" yaml:"TypeSrc,omitempty"`
	PhotoTitle       string    `gorm:"type:VARCHAR(255);" json:"Title" yaml:"Title"`
	TitleSrc         string    `gorm:"type:VARBINARY(8);" json:"TitleSrc" yaml:"TitleSrc,omitempty"`
	PhotoDescription string    `gorm:"type:TEXT;" json:"Description" yaml:"Description,omitempty"`
	DescriptionSrc   string    `gorm:"type:VARBINARY(8);" json:"DescriptionSrc" yaml:"DescriptionSrc,omitempty"`
	PhotoPath        string    `gorm:"type:VARBINARY(500);index:idx_photos_path_name;" json:"Path" yaml:"-"`
	PhotoName        string    `gorm:"type:VARBINARY(255);index:idx_photos_path_name;" json:"Name" yaml:"-"`
	OriginalName     string    `gorm:"type:VARBINARY(755);" json:"OriginalName" yaml:"OriginalName,omitempty"`
	PhotoStack       int8      `json:"Stack" yaml:"Stack,omitempty"`
	PhotoFavorite    bool      `json:"Favorite" yaml:"Favorite,omitempty"`
	PhotoPrivate     bool      `json:"Private" yaml:"Private,omitempty"`
	PhotoScan        bool      `json:"Scan" yaml:"Scan,omitempty"`
	PhotoPanorama    bool      `json:"Panorama" yaml:"Panorama,omitempty"`
	TimeZone         string    `gorm:"type:VARBINARY(64);" json:"TimeZone" yaml:"-"`
	PlaceID          string    `gorm:"type:VARBINARY(42);index;default:'zz'" json:"PlaceID" yaml:"-"`
	PlaceSrc         string    `gorm:"type:VARBINARY(8);" json:"PlaceSrc" yaml:"PlaceSrc,omitempty"`
	CellID           string    `gorm:"type:VARBINARY(42);index;default:'zz'" json:"CellID" yaml:"-"`
	CellAccuracy     int       `json:"CellAccuracy" yaml:"CellAccuracy,omitempty"`
	PhotoAltitude    int       `json:"Altitude" yaml:"Altitude,omitempty"`
	PhotoLat         float32   `gorm:"type:FLOAT;index;" json:"Lat" yaml:"Lat,omitempty"`
	PhotoLng         float32   `gorm:"type:FLOAT;index;" json:"Lng" yaml:"Lng,omitempty"`
	PhotoCountry     string    `gorm:"type:VARBINARY(2);index:idx_photos_country_year_month;default:'zz'" json:"Country" yaml:"-"`
	PhotoYear        int       `gorm:"index:idx_photos_country_year_month;" json:"Year" yaml:"Year"`
	PhotoMonth       int       `gorm:"index:idx_photos_country_year_month;" json:"Month" yaml:"Month"`
	PhotoDay         int       `json:"Day" yaml:"Day"`
	PhotoIso         int       `json:"Iso" yaml:"ISO,omitempty"`
	PhotoExposure    string    `gorm:"type:VARBINARY(64);" json:"Exposure" yaml:"Exposure,omitempty"`
	PhotoFNumber     float32   `gorm:"type:FLOAT;" json:"FNumber" yaml:"FNumber,omitempty"`
	PhotoFocalLength int       `json:"FocalLength" yaml:"FocalLength,omitempty"`
	PhotoQuality     int       `gorm:"type:SMALLINT" json:"Quality" yaml:"-"`
	PhotoResolution  int       `gorm:"type:SMALLINT" json:"Resolution" yaml:"-"`
	PhotoColor       uint8     `json:"Color" yaml:"-"`
	CameraID         uint      `gorm:"index:idx_photos_camera_lens;default:1" json:"CameraID" yaml:"-"`
	CameraSerial     string    `gorm:"type:VARBINARY(255);" json:"CameraSerial" yaml:"CameraSerial,omitempty"`
	CameraSrc        string    `gorm:"type:VARBINARY(8);" json:"CameraSrc" yaml:"-"`
	LensID           uint      `gorm:"index:idx_photos_camera_lens;default:1" json:"LensID" yaml:"-"`
	//Details          *Details  `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Details" yaml:"Details"`
	//Camera           *Camera      `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Camera" yaml:"-"`
	//Lens             *Lens        `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Lens" yaml:"-"`
	//Cell             *Cell        `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Cell" yaml:"-"`
	//Place            *Place       `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Place" yaml:"-"`
	//Keywords         []Keyword    `json:"-" yaml:"-"`
	//Albums           []Album      `json:"-" yaml:"-"`
	Files []File `yaml:"-"`
	//Labels           []PhotoLabel `yaml:"-"`
	CreatedAt time.Time  `yaml:"CreatedAt,omitempty"`
	UpdatedAt time.Time  `yaml:"UpdatedAt,omitempty"`
	EditedAt  *time.Time `yaml:"EditedAt,omitempty"`
	CheckedAt *time.Time `sql:"index" yaml:"-"`
	DeletedAt *time.Time `sql:"index" yaml:"DeletedAt,omitempty"`
}

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
