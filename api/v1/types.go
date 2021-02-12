package api

import (
	"database/sql"
	"time"
)

type Photos []Photo

// Photo represents a photo, all its properties, and link to all its images and sidecar files.
type Photo struct {
	UUID             string       `gorm:"type:VARBINARY(42);index;" json:"DocumentID,omitempty" yaml:"DocumentID,omitempty"`
	TakenAt          time.Time    `gorm:"type:datetime;index:idx_photos_taken_uid;" json:"TakenAt" yaml:"TakenAt"`
	TakenAtLocal     time.Time    `gorm:"type:datetime;" yaml:"-"`
	TakenSrc         string       `gorm:"type:VARBINARY(8);" json:"TakenSrc" yaml:"TakenSrc,omitempty"`
	PhotoUID         string       `gorm:"type:VARBINARY(42);unique_index;index:idx_photos_taken_uid;" json:"UID" yaml:"UID"`
	PhotoType        string       `gorm:"type:VARBINARY(8);default:'image';" json:"Type" yaml:"Type"`
	TypeSrc          string       `gorm:"type:VARBINARY(8);" json:"TypeSrc" yaml:"TypeSrc,omitempty"`
	PhotoTitle       string       `gorm:"type:VARCHAR(255);" json:"Title" yaml:"Title"`
	TitleSrc         string       `gorm:"type:VARBINARY(8);" json:"TitleSrc" yaml:"TitleSrc,omitempty"`
	PhotoDescription string       `gorm:"type:TEXT;" json:"Description" yaml:"Description,omitempty"`
	DescriptionSrc   string       `gorm:"type:VARBINARY(8);" json:"DescriptionSrc" yaml:"DescriptionSrc,omitempty"`
	PhotoPath        string       `gorm:"type:VARBINARY(500);index:idx_photos_path_name;" json:"Path" yaml:"-"`
	PhotoName        string       `gorm:"type:VARBINARY(255);index:idx_photos_path_name;" json:"Name" yaml:"-"`
	OriginalName     string       `gorm:"type:VARBINARY(755);" json:"OriginalName" yaml:"OriginalName,omitempty"`
	PhotoStack       int8         `json:"Stack" yaml:"Stack,omitempty"`
	PhotoFavorite    bool         `json:"Favorite" yaml:"Favorite,omitempty"`
	PhotoPrivate     bool         `json:"Private" yaml:"Private,omitempty"`
	PhotoScan        bool         `json:"Scan" yaml:"Scan,omitempty"`
	PhotoPanorama    bool         `json:"Panorama" yaml:"Panorama,omitempty"`
	TimeZone         string       `gorm:"type:VARBINARY(64);" json:"TimeZone" yaml:"-"`
	PlaceID          string       `gorm:"type:VARBINARY(42);index;default:'zz'" json:"PlaceID" yaml:"-"`
	PlaceSrc         string       `gorm:"type:VARBINARY(8);" json:"PlaceSrc" yaml:"PlaceSrc,omitempty"`
	CellID           string       `gorm:"type:VARBINARY(42);index;default:'zz'" json:"CellID" yaml:"-"`
	CellAccuracy     int          `json:"CellAccuracy" yaml:"CellAccuracy,omitempty"`
	PhotoAltitude    int          `json:"Altitude" yaml:"Altitude,omitempty"`
	PhotoLat         float32      `gorm:"type:FLOAT;index;" json:"Lat" yaml:"Lat,omitempty"`
	PhotoLng         float32      `gorm:"type:FLOAT;index;" json:"Lng" yaml:"Lng,omitempty"`
	PhotoCountry     string       `gorm:"type:VARBINARY(2);index:idx_photos_country_year_month;default:'zz'" json:"Country" yaml:"-"`
	PhotoYear        int          `gorm:"index:idx_photos_country_year_month;" json:"Year" yaml:"Year"`
	PhotoMonth       int          `gorm:"index:idx_photos_country_year_month;" json:"Month" yaml:"Month"`
	PhotoDay         int          `json:"Day" yaml:"Day"`
	PhotoIso         int          `json:"Iso" yaml:"ISO,omitempty"`
	PhotoExposure    string       `gorm:"type:VARBINARY(64);" json:"Exposure" yaml:"Exposure,omitempty"`
	PhotoFNumber     float32      `gorm:"type:FLOAT;" json:"FNumber" yaml:"FNumber,omitempty"`
	PhotoFocalLength int          `json:"FocalLength" yaml:"FocalLength,omitempty"`
	PhotoQuality     int          `gorm:"type:SMALLINT" json:"Quality" yaml:"-"`
	PhotoResolution  int          `gorm:"type:SMALLINT" json:"Resolution" yaml:"-"`
	PhotoColor       uint8        `json:"Color" yaml:"-"`
	CameraID         uint         `gorm:"index:idx_photos_camera_lens;default:1" json:"CameraID" yaml:"-"`
	CameraSerial     string       `gorm:"type:VARBINARY(255);" json:"CameraSerial" yaml:"CameraSerial,omitempty"`
	CameraSrc        string       `gorm:"type:VARBINARY(8);" json:"CameraSrc" yaml:"-"`
	LensID           uint         `gorm:"index:idx_photos_camera_lens;default:1" json:"LensID" yaml:"-"`
	Details          *Details     `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Details" yaml:"Details"`
	Camera           *Camera      `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Camera" yaml:"-"`
	Lens             *Lens        `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Lens" yaml:"-"`
	Cell             *Cell        `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Cell" yaml:"-"`
	Place            *Place       `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false" json:"Place" yaml:"-"`
	Keywords         []Keyword    `json:"-" yaml:"-"`
	Albums           []Album      `json:"-" yaml:"-"`
	Files            []File       `yaml:"-"`
	Labels           []PhotoLabel `yaml:"-"`
	CreatedAt        time.Time    `yaml:"CreatedAt,omitempty"`
	UpdatedAt        time.Time    `yaml:"UpdatedAt,omitempty"`
	EditedAt         *time.Time   `yaml:"EditedAt,omitempty"`
	CheckedAt        *time.Time   `sql:"index" yaml:"-"`
	DeletedAt        *time.Time   `sql:"index" yaml:"DeletedAt,omitempty"`
}

// Details stores additional metadata fields for each photo to improve search performance.
type Details struct {
	PhotoID      uint      `gorm:"primary_key;auto_increment:false" yaml:"-"`
	Keywords     string    `gorm:"type:TEXT;" json:"Keywords" yaml:"Keywords"`
	KeywordsSrc  string    `gorm:"type:VARBINARY(8);" json:"KeywordsSrc" yaml:"KeywordsSrc,omitempty"`
	Notes        string    `gorm:"type:TEXT;" json:"Notes" yaml:"Notes,omitempty"`
	NotesSrc     string    `gorm:"type:VARBINARY(8);" json:"NotesSrc" yaml:"NotesSrc,omitempty"`
	Subject      string    `gorm:"type:VARCHAR(255);" json:"Subject" yaml:"Subject,omitempty"`
	SubjectSrc   string    `gorm:"type:VARBINARY(8);" json:"SubjectSrc" yaml:"SubjectSrc,omitempty"`
	Artist       string    `gorm:"type:VARCHAR(255);" json:"Artist" yaml:"Artist,omitempty"`
	ArtistSrc    string    `gorm:"type:VARBINARY(8);" json:"ArtistSrc" yaml:"ArtistSrc,omitempty"`
	Copyright    string    `gorm:"type:VARCHAR(255);" json:"Copyright" yaml:"Copyright,omitempty"`
	CopyrightSrc string    `gorm:"type:VARBINARY(8);" json:"CopyrightSrc" yaml:"CopyrightSrc,omitempty"`
	License      string    `gorm:"type:VARCHAR(255);" json:"License" yaml:"License,omitempty"`
	LicenseSrc   string    `gorm:"type:VARBINARY(8);" json:"LicenseSrc" yaml:"LicenseSrc,omitempty"`
	CreatedAt    time.Time `yaml:"-"`
	UpdatedAt    time.Time `yaml:"-"`
}

// Camera model and make (as extracted from UpdateExif metadata)
type Camera struct {
	ID                uint       `gorm:"primary_key" json:"ID" yaml:"ID"`
	CameraSlug        string     `gorm:"type:VARBINARY(255);unique_index;" json:"Slug" yaml:"-"`
	CameraName        string     `gorm:"type:VARCHAR(255);" json:"Name" yaml:"Name"`
	CameraMake        string     `gorm:"type:VARCHAR(255);" json:"Make" yaml:"Make,omitempty"`
	CameraModel       string     `gorm:"type:VARCHAR(255);" json:"Model" yaml:"Model,omitempty"`
	CameraType        string     `gorm:"type:VARCHAR(255);" json:"Type,omitempty" yaml:"Type,omitempty"`
	CameraDescription string     `gorm:"type:TEXT;" json:"Description,omitempty" yaml:"Description,omitempty"`
	CameraNotes       string     `gorm:"type:TEXT;" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	CreatedAt         time.Time  `json:"-" yaml:"-"`
	UpdatedAt         time.Time  `json:"-" yaml:"-"`
	DeletedAt         *time.Time `sql:"index" json:"-" yaml:"-"`
}

// Lens represents camera lens (as extracted from UpdateExif metadata)
type Lens struct {
	ID              uint       `gorm:"primary_key" json:"ID" yaml:"ID"`
	LensSlug        string     `gorm:"type:VARBINARY(255);unique_index;" json:"Slug" yaml:"Slug,omitempty"`
	LensName        string     `gorm:"type:VARCHAR(255);" json:"Name" yaml:"Name"`
	LensMake        string     `gorm:"type:VARCHAR(255);" json:"Make" yaml:"Make,omitempty"`
	LensModel       string     `gorm:"type:VARCHAR(255);" json:"Model" yaml:"Model,omitempty"`
	LensType        string     `gorm:"type:VARCHAR(255);" json:"Type" yaml:"Type,omitempty"`
	LensDescription string     `gorm:"type:TEXT;" json:"Description,omitempty" yaml:"Description,omitempty"`
	LensNotes       string     `gorm:"type:TEXT;" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	CreatedAt       time.Time  `json:"-" yaml:"-"`
	UpdatedAt       time.Time  `json:"-" yaml:"-"`
	DeletedAt       *time.Time `sql:"index" json:"-" yaml:"-"`
}

// Cell represents a S2 cell with location data.
type Cell struct {
	ID           string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false;" json:"ID" yaml:"ID"`
	CellName     string    `gorm:"type:VARCHAR(255);" json:"Name" yaml:"Name,omitempty"`
	CellCategory string    `gorm:"type:VARCHAR(64);" json:"Category" yaml:"Category,omitempty"`
	PlaceID      string    `gorm:"type:VARBINARY(42);default:'zz'" json:"-" yaml:"PlaceID"`
	Place        *Place    `gorm:"PRELOAD:true" json:"Place" yaml:"-"`
	CreatedAt    time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt    time.Time `json:"UpdatedAt" yaml:"-"`
}

// Place used to associate photos to places
type Place struct {
	ID            string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false;" json:"PlaceID" yaml:"PlaceID"`
	PlaceLabel    string    `gorm:"type:VARBINARY(755);unique_index;" json:"Label" yaml:"Label"`
	PlaceCity     string    `gorm:"type:VARCHAR(255);" json:"City" yaml:"City,omitempty"`
	PlaceState    string    `gorm:"type:VARCHAR(255);" json:"State" yaml:"State,omitempty"`
	PlaceCountry  string    `gorm:"type:VARBINARY(2);" json:"Country" yaml:"Country,omitempty"`
	PlaceKeywords string    `gorm:"type:VARCHAR(255);" json:"Keywords" yaml:"Keywords,omitempty"`
	PlaceFavorite bool      `json:"Favorite" yaml:"Favorite,omitempty"`
	PhotoCount    int       `gorm:"default:1" json:"PhotoCount" yaml:"-"`
	CreatedAt     time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt     time.Time `json:"UpdatedAt" yaml:"-"`
}

// Keyword used for full text search
type Keyword struct {
	ID      uint   `gorm:"primary_key"`
	Keyword string `gorm:"type:VARCHAR(64);index;"`
	Skip    bool
}

// Album represents a photo album
type Album struct {
	ID               uint        `gorm:"primary_key" json:"ID" yaml:"-"`
	AlbumUID         string      `gorm:"type:VARBINARY(42);unique_index;" json:"UID" yaml:"UID"`
	CoverUID         string      `gorm:"type:VARBINARY(42);" json:"CoverUID" yaml:"CoverUID,omitempty"`
	FolderUID        string      `gorm:"type:VARBINARY(42);index;" json:"FolderUID" yaml:"FolderUID,omitempty"`
	AlbumSlug        string      `gorm:"type:VARBINARY(255);index;" json:"Slug" yaml:"Slug"`
	AlbumPath        string      `gorm:"type:VARBINARY(500);index;" json:"Path" yaml:"-"`
	AlbumType        string      `gorm:"type:VARBINARY(8);default:'album';" json:"Type" yaml:"Type,omitempty"`
	AlbumTitle       string      `gorm:"type:VARCHAR(255);" json:"Title" yaml:"Title"`
	AlbumLocation    string      `gorm:"type:VARCHAR(255);" json:"Location" yaml:"Location,omitempty"`
	AlbumCategory    string      `gorm:"type:VARCHAR(255);index;" json:"Category" yaml:"Category,omitempty"`
	AlbumCaption     string      `gorm:"type:TEXT;" json:"Caption" yaml:"Caption,omitempty"`
	AlbumDescription string      `gorm:"type:TEXT;" json:"Description" yaml:"Description,omitempty"`
	AlbumNotes       string      `gorm:"type:TEXT;" json:"Notes" yaml:"Notes,omitempty"`
	AlbumFilter      string      `gorm:"type:VARBINARY(1024);" json:"Filter" yaml:"Filter,omitempty"`
	AlbumOrder       string      `gorm:"type:VARBINARY(32);" json:"Order" yaml:"Order,omitempty"`
	AlbumTemplate    string      `gorm:"type:VARBINARY(255);" json:"Template" yaml:"Template,omitempty"`
	AlbumCountry     string      `gorm:"type:VARBINARY(2);index:idx_albums_country_year_month;default:'zz'" json:"Country" yaml:"Country,omitempty"`
	AlbumYear        int         `gorm:"index:idx_albums_country_year_month;" json:"Year" yaml:"Year,omitempty"`
	AlbumMonth       int         `gorm:"index:idx_albums_country_year_month;" json:"Month" yaml:"Month,omitempty"`
	AlbumDay         int         `json:"Day" yaml:"Day,omitempty"`
	AlbumFavorite    bool        `json:"Favorite" yaml:"Favorite,omitempty"`
	AlbumPrivate     bool        `json:"Private" yaml:"Private,omitempty"`
	CreatedAt        time.Time   `json:"CreatedAt" yaml:"CreatedAt,omitempty"`
	UpdatedAt        time.Time   `json:"UpdatedAt" yaml:"UpdatedAt,omitempty"`
	DeletedAt        *time.Time  `sql:"index" json:"DeletedAt" yaml:"DeletedAt,omitempty"`
	Photos           PhotoAlbums `gorm:"foreignkey:AlbumUID;association_foreignkey:AlbumUID" json:"-" yaml:"Photos,omitempty"`
}

type PhotoAlbums []PhotoAlbum

// PhotoAlbum represents the many_to_many relation between Photo and Album
type PhotoAlbum struct {
	PhotoUID  string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false" json:"PhotoUID" yaml:"UID"`
	AlbumUID  string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false;index" json:"AlbumUID" yaml:"-"`
	Order     int       `json:"Order" yaml:"Order,omitempty"`
	Hidden    bool      `json:"Hidden" yaml:"Hidden,omitempty"`
	Missing   bool      `json:"Missing" yaml:"Missing,omitempty"`
	CreatedAt time.Time `json:"CreatedAt" yaml:"CreatedAt,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt" yaml:"-"`
	Photo     *Photo    `gorm:"PRELOAD:false" yaml:"-"`
	Album     *Album    `gorm:"PRELOAD:true" yaml:"-"`
}

type Files []File

// File represents an image or sidecar file that belongs to a photo.
type File struct {
	ID              uint          `gorm:"primary_key" json:"-" yaml:"-"`
	Photo           *Photo        `json:"-" yaml:"-"`
	PhotoID         uint          `gorm:"index;" json:"-" yaml:"-"`
	PhotoUID        string        `gorm:"type:VARBINARY(42);index;" json:"PhotoUID" yaml:"PhotoUID"`
	InstanceID      string        `gorm:"type:VARBINARY(42);index;" json:"InstanceID,omitempty" yaml:"InstanceID,omitempty"`
	FileUID         string        `gorm:"type:VARBINARY(42);unique_index;" json:"UID" yaml:"UID"`
	FileName        string        `gorm:"type:VARBINARY(755);unique_index:idx_files_name_root;" json:"Name" yaml:"Name"`
	FileRoot        string        `gorm:"type:VARBINARY(16);default:'/';unique_index:idx_files_name_root;" json:"Root" yaml:"Root,omitempty"`
	OriginalName    string        `gorm:"type:VARBINARY(755);" json:"OriginalName" yaml:"OriginalName,omitempty"`
	FileHash        string        `gorm:"type:VARBINARY(128);index" json:"Hash" yaml:"Hash,omitempty"`
	FileSize        int64         `json:"Size" yaml:"Size,omitempty"`
	FileCodec       string        `gorm:"type:VARBINARY(32)" json:"Codec" yaml:"Codec,omitempty"`
	FileType        string        `gorm:"type:VARBINARY(32)" json:"Type" yaml:"Type,omitempty"`
	FileMime        string        `gorm:"type:VARBINARY(64)" json:"Mime" yaml:"Mime,omitempty"`
	FilePrimary     bool          `json:"Primary" yaml:"Primary,omitempty"`
	FileSidecar     bool          `json:"Sidecar" yaml:"Sidecar,omitempty"`
	FileMissing     bool          `json:"Missing" yaml:"Missing,omitempty"`
	FilePortrait    bool          `json:"Portrait" yaml:"Portrait,omitempty"`
	FileVideo       bool          `json:"Video" yaml:"Video,omitempty"`
	FileDuration    time.Duration `json:"Duration" yaml:"Duration,omitempty"`
	FileWidth       int           `json:"Width" yaml:"Width,omitempty"`
	FileHeight      int           `json:"Height" yaml:"Height,omitempty"`
	FileOrientation int           `json:"Orientation" yaml:"Orientation,omitempty"`
	FileProjection  string        `gorm:"type:VARBINARY(16);" json:"Projection,omitempty" yaml:"Projection,omitempty"`
	FileAspectRatio float32       `gorm:"type:FLOAT;" json:"AspectRatio" yaml:"AspectRatio,omitempty"`
	FileMainColor   string        `gorm:"type:VARBINARY(16);index;" json:"MainColor" yaml:"MainColor,omitempty"`
	FileColors      string        `gorm:"type:VARBINARY(9);" json:"Colors" yaml:"Colors,omitempty"`
	FileLuminance   string        `gorm:"type:VARBINARY(9);" json:"Luminance" yaml:"Luminance,omitempty"`
	FileDiff        uint32        `json:"Diff" yaml:"Diff,omitempty"`
	FileChroma      uint8         `json:"Chroma" yaml:"Chroma,omitempty"`
	FileError       string        `gorm:"type:VARBINARY(512)" json:"Error" yaml:"Error,omitempty"`
	ModTime         int64         `json:"ModTime" yaml:"-"`
	CreatedAt       time.Time     `json:"CreatedAt" yaml:"-"`
	CreatedIn       int64         `json:"CreatedIn" yaml:"-"`
	UpdatedAt       time.Time     `json:"UpdatedAt" yaml:"-"`
	UpdatedIn       int64         `json:"UpdatedIn" yaml:"-"`
	DeletedAt       *time.Time    `sql:"index" json:"DeletedAt,omitempty" yaml:"-"`
	Share           []FileShare   `json:"-" yaml:"-"`
	Sync            []FileSync    `json:"-" yaml:"-"`
}

// FileSync represents a one-to-many relation between File and Account for syncing with remote services.
type FileSync struct {
	RemoteName string `gorm:"primary_key;auto_increment:false;type:VARBINARY(255)"`
	AccountID  uint   `gorm:"primary_key;auto_increment:false"`
	FileID     uint   `gorm:"index;"`
	RemoteDate time.Time
	RemoteSize int64
	Status     string `gorm:"type:VARBINARY(16);"`
	Error      string `gorm:"type:VARBINARY(512);"`
	Errors     int
	File       *File
	Account    *Account
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// FileShare represents a one-to-many relation between File and Account for pushing files to remote services.
type FileShare struct {
	FileID     uint   `gorm:"primary_key;auto_increment:false"`
	AccountID  uint   `gorm:"primary_key;auto_increment:false"`
	RemoteName string `gorm:"primary_key;auto_increment:false;type:VARBINARY(255)"`
	Status     string `gorm:"type:VARBINARY(16);"`
	Error      string `gorm:"type:VARBINARY(512);"`
	Errors     int
	File       *File
	Account    *Account
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// PhotoLabel represents the many-to-many relation between Photo and label.
// Labels are weighted by uncertainty (100 - confidence)
type PhotoLabel struct {
	PhotoID     uint   `gorm:"primary_key;auto_increment:false"`
	LabelID     uint   `gorm:"primary_key;auto_increment:false;index"`
	LabelSrc    string `gorm:"type:VARBINARY(8);"`
	Uncertainty int    `gorm:"type:SMALLINT"`
	Photo       *Photo `gorm:"PRELOAD:false"`
	Label       *Label `gorm:"PRELOAD:true"`
}

// Label is used for photo, album and location categorization
type Label struct {
	ID               uint       `gorm:"primary_key" json:"ID" yaml:"-"`
	LabelUID         string     `gorm:"type:VARBINARY(42);unique_index;" json:"UID" yaml:"UID"`
	LabelSlug        string     `gorm:"type:VARBINARY(255);unique_index;" json:"Slug" yaml:"-"`
	CustomSlug       string     `gorm:"type:VARBINARY(255);index;" json:"CustomSlug" yaml:"-"`
	LabelName        string     `gorm:"type:VARCHAR(255);" json:"Name" yaml:"Name"`
	LabelPriority    int        `json:"Priority" yaml:"Priority,omitempty"`
	LabelFavorite    bool       `json:"Favorite" yaml:"Favorite,omitempty"`
	LabelDescription string     `gorm:"type:TEXT;" json:"Description" yaml:"Description,omitempty"`
	LabelNotes       string     `gorm:"type:TEXT;" json:"Notes" yaml:"Notes,omitempty"`
	LabelCategories  []*Label   `gorm:"many2many:categories;association_jointable_foreignkey:category_id" json:"-" yaml:"-"`
	PhotoCount       int        `gorm:"default:1" json:"PhotoCount" yaml:"-"`
	CreatedAt        time.Time  `json:"CreatedAt" yaml:"-"`
	UpdatedAt        time.Time  `json:"UpdatedAt" yaml:"-"`
	DeletedAt        *time.Time `sql:"index" json:"DeletedAt,omitempty" yaml:"-"`
	New              bool       `gorm:"-" json:"-" yaml:"-"`
}

// FileInfos represents meta data about a file
type FileInfos struct {
	FileWidth       int
	FileHeight      int
	FileOrientation int
	FileAspectRatio float32
	FileMainColor   string
	FileColors      string
	FileLuminance   string
	FileDiff        uint32
	FileChroma      uint8
}

// Account represents a remote service account for uploading, downloading or syncing media files.
type Account struct {
	ID            uint   `gorm:"primary_key"`
	AccName       string `gorm:"type:VARCHAR(255);"`
	AccOwner      string `gorm:"type:VARCHAR(255);"`
	AccURL        string `gorm:"type:VARBINARY(512);"`
	AccType       string `gorm:"type:VARBINARY(255);"`
	AccKey        string `gorm:"type:VARBINARY(255);"`
	AccUser       string `gorm:"type:VARBINARY(255);"`
	AccPass       string `gorm:"type:VARBINARY(255);"`
	AccError      string `gorm:"type:VARBINARY(512);"`
	AccErrors     int
	AccShare      bool
	AccSync       bool
	RetryLimit    int
	SharePath     string `gorm:"type:VARBINARY(500);"`
	ShareSize     string `gorm:"type:VARBINARY(16);"`
	ShareExpires  int
	SyncPath      string `gorm:"type:VARBINARY(500);"`
	SyncStatus    string `gorm:"type:VARBINARY(16);"`
	SyncInterval  int
	SyncDate      sql.NullTime `deepcopier:"skip"`
	SyncUpload    bool
	SyncDownload  bool
	SyncFilenames bool
	SyncRaw       bool
	CreatedAt     time.Time  `deepcopier:"skip"`
	UpdatedAt     time.Time  `deepcopier:"skip"`
	DeletedAt     *time.Time `deepcopier:"skip" sql:"index"`
}
