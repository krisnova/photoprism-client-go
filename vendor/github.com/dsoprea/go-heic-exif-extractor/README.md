[![Build Status](https://travis-ci.org/dsoprea/go-heic-exif-extractor.svg?branch=master)](https://travis-ci.org/dsoprea/go-heic-exif-extractor)
[![codecov](https://codecov.io/gh/dsoprea/go-heic-exif-extractor/branch/master/graph/badge.svg)](https://codecov.io/gh/dsoprea/go-heic-exif-extractor)
[![Report Card](https://goreportcard.com/badge/github.com/dsoprea/go-heic-exif-extractor/v2)](https://goreportcard.com/report/github.com/dsoprea/go-heic-exif-extractor/v2)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor/v2?status.svg)](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor/v2)

# Overview

This project invokes a third-party project to parse HEIC/HEIF ([ISO 23008-12](https://www.iso.org/standard/66067.html)) content and extract an EXIF blob and then invokes [go-exif](https://github.com/dsoprea/go-exif) to parse that EXIF blob. It satisfies the [riimage.MediaParser](https://github.com/dsoprea/go-utility/blob/master/v2/image/media_parser_type.go) interface and exists to provide HEIC/HEIF support to the [go-exif-knife](https://github.com/dsoprea/go-exif-knife) tool.

# Examples

See the [GoDoc page](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor/v2) for [usage examples](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor/v2#pkg-examples).
