package tiffstructure

import (
	"bytes"
	"io"
	"os"

	"io/ioutil"

	"github.com/dsoprea/go-exif/v2"
	"github.com/dsoprea/go-logging"
	"github.com/dsoprea/go-utility/image"
)

// TiffMediaParser knows how to parse a TIFF stream.
type TiffMediaParser struct {
}

// NewTiffMediaParser returns a new `TiffMediaParser` struct.
func NewTiffMediaParser() *TiffMediaParser {
	return new(TiffMediaParser)
}

// TiffMediaContext encapsulates parsed EXIF data.
type TiffMediaContext struct {
	rootIfd *exif.Ifd
	data    []byte
}

// Exif is an accessor that returns the root IFD.
func (tmc *TiffMediaContext) Exif() (rootIfd *exif.Ifd, data []byte, err error) {

	// TODO(dustin): Add test

	return tmc.rootIfd, tmc.data, nil
}

// Parse parses a TIFF stream given a `io.ReadSeeker`.
func (tmp *TiffMediaParser) Parse(rs io.ReadSeeker, size int) (mc riimage.MediaContext, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	exifData, err := ioutil.ReadAll(rs)
	log.PanicIf(err)

	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, exifData)
	log.PanicIf(err)

	tmc := &TiffMediaContext{
		rootIfd: index.RootIfd,
		data:    exifData,
	}

	return tmc, nil
}

// ParseFile parses a TIFF stream given a file-path.
func (tmp *TiffMediaParser) ParseFile(filepath string) (mc riimage.MediaContext, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	stat, err := f.Stat()
	log.PanicIf(err)

	size := stat.Size()

	mc, err = tmp.Parse(f, int(size))
	log.PanicIf(err)

	return mc, nil
}

// ParseBytes parses a TIFF stream given a byte-slice.
func (tmp *TiffMediaParser) ParseBytes(data []byte) (mc riimage.MediaContext, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	// TODO(dustin): Add test

	br := bytes.NewReader(data)

	mc, err = tmp.Parse(br, len(data))
	log.PanicIf(err)

	return mc, nil
}

// LooksLikeFormat returns a boolean indicating whether this looks like a TIFF
// stream.
func (tmp *TiffMediaParser) LooksLikeFormat(data []byte) bool {
	if len(data) < exif.ExifSignatureLength {
		return false
	}

	// TODO(dustin): Add test

	_, err := exif.ParseExifHeader(data)
	return err == nil
}

var (
	// Enforce interface conformance.
	_ riimage.MediaParser  = new(TiffMediaParser)
	_ riimage.MediaContext = new(TiffMediaContext)
)
