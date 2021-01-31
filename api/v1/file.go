package api

import "time"

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
	//Share           []FileShare   `json:"-" yaml:"-"`
	//Sync            []FileSync    `json:"-" yaml:"-"`
}

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
