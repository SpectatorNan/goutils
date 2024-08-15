package imagex

import "C"

const (
	WebpMax = 16383
	AvifMax = 65536
)

type ExtraParams struct {
	Width     int // in px
	Height    int // in px
	MaxWidth  int // in px
	MaxHeight int // in px
}

type ImageInfo struct {
	ImageType ImageType
	ImageSize ImageSize
}

// ImageType represents an image type
type ImageType int

// ImageType enum
const (
	ImageTypeUnknown ImageType = iota
	ImageTypeGIF
	ImageTypeJPEG
	ImageTypeMagick
	ImageTypePDF
	ImageTypePNG
	ImageTypeSVG
	ImageTypeTIFF
	ImageTypeWEBP
	ImageTypeHEIF
	ImageTypeBMP
	ImageTypeAVIF
	ImageTypeJP2K
	ImageTypeJXL
)

type ImageSize struct {
	Width  int
	Height int
}
