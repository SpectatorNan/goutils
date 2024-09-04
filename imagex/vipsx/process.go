package vipsx

import (
	"errors"
	"github.com/SpectatorNan/goutils/imagex"
	"github.com/davidbyttow/govips/v2/vips"
	"slices"
)

var (
	// Source image encoder ignore list for WebP and AVIF
	// We shouldn't convert Unknown and AVIF to WebP
	webpIgnore = []vips.ImageType{vips.ImageTypeUnknown, vips.ImageTypeAVIF}
	// We shouldn't convert Unknown,AVIF and GIF to AVIF
	avifIgnore = append(webpIgnore, vips.ImageTypeGIF)
)

func preProcessImage(img *vips.ImageRef, imageType string) error {
	// Check Width/Height and ignore image formats
	switch imageType {
	case "webp":
		if img.Metadata().Width > imagex.WebpMax || img.Metadata().Height > imagex.WebpMax {
			return errors.New("WebP: image too large")
		}
		imageFormat := img.Format()
		if slices.Contains(webpIgnore, imageFormat) {
			// Return err to render original image
			return errors.New("WebP encoder: ignore image type")
		}
	case "avif":
		if img.Metadata().Width > imagex.AvifMax || img.Metadata().Height > imagex.AvifMax {
			return errors.New("AVIF: image too large")
		}
		imageFormat := img.Format()
		if slices.Contains(avifIgnore, imageFormat) {
			// Return err to render original image
			return errors.New("AVIF encoder: ignore image type")
		}
	}
	// Auto rotate
	err := img.AutoRotate()
	if err != nil {
		return err
	}
	//if enableExtraParams {
	//	err = resizeImage(img, extraParams, extraParamsCropInteresting)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}
