package vipsx

import (
	"github.com/SpectatorNan/goutils/imagex"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/pkg/errors"
	"io"
)

func LoadImageConfig(reader io.Reader) (*imagex.ImageInfo, error) {
	ref, err := loadImageFromReader(reader)
	if err != nil {
		return nil, err
	}
	meta := ref.Metadata()

	return &imagex.ImageInfo{
		ImageSize: imagex.ImageSize{
			Width:  meta.Width,
			Height: meta.Height,
		},
		ImageType: convertImageTypeForVips2ImageX(meta.Format),
	}, nil
}

func convertImageTypeForVips2ImageX(vipsType vips.ImageType) imagex.ImageType {
	switch vipsType {
	case vips.ImageTypeUnknown:
		return imagex.ImageTypeUnknown
	case vips.ImageTypeGIF:
		return imagex.ImageTypeGIF
	case vips.ImageTypeJPEG:
		return imagex.ImageTypeJPEG
	case vips.ImageTypeMagick:
		return imagex.ImageTypeMagick
	case vips.ImageTypePDF:
		return imagex.ImageTypePDF
	case vips.ImageTypePNG:
		return imagex.ImageTypePNG
	case vips.ImageTypeSVG:
		return imagex.ImageTypeSVG
	case vips.ImageTypeTIFF:
		return imagex.ImageTypeTIFF
	case vips.ImageTypeWEBP:
		return imagex.ImageTypeWEBP
	case vips.ImageTypeHEIF:
		return imagex.ImageTypeHEIF
	case vips.ImageTypeBMP:
		return imagex.ImageTypeBMP
	case vips.ImageTypeAVIF:
		return imagex.ImageTypeAVIF
	case vips.ImageTypeJP2K:
		return imagex.ImageTypeJP2K
	case vips.ImageTypeJXL:
		return imagex.ImageTypeJXL
	default:
		return imagex.ImageTypeUnknown
	}
}

func loadImageFromReader(reader io.Reader) (*vips.ImageRef, error) {
	imageBuf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	img, err := vips.LoadImageFromBuffer(imageBuf, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load image from buffer")
	}
	return img, nil
}
