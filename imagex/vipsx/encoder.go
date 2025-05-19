package vipsx

import (
	"github.com/SpectatorNan/goutils/errors"
	export2 "github.com/SpectatorNan/goutils/imagex/vipsx/export"
	"github.com/davidbyttow/govips/v2/vips"
)

func webpEncoder(img *vips.ImageRef, quality int) ([]byte, error) {
	var (
		buf []byte
		err error
	)

	if quality >= 100 {
		//buf, _, err = img.ExportWebp(&vips.WebpExportParams{
		//	Lossless:      true,
		//	StripMetadata: true,
		//})
		buf, err = export2.ExportWebpWithQuality(img, true, 0, true, false, 0)
	} else {
		//ep := vips.WebpExportParams{
		//	Quality:       quality,
		//	Lossless:      false,
		//	StripMetadata: true,
		//}
		// If some special images cannot encode with default ReductionEffort(0), then retry from 0 to 6
		// Example: https://github.com/webp-sh/webp_server_go/issues/234
		for i := 0; i < 7; i++ {
			//ep.ReductionEffort = i
			//buf, _, err = img.ExportWebp(&ep)
			buf, err = export2.ExportWebpWithQuality(img, true, quality, false, false, i)
			if err == nil {
				break
			}
		}
		//buf, _, err = img.ExportWebp(&ep)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to encode image to WebP")
	}

	return buf, nil
}

func jxlEncoder(img *vips.ImageRef, quality int) ([]byte, error) {
	var (
		buf []byte
		err error
	)

	// If quality >= 100, we use lossless mode
	if quality >= 100 {
		//buf, _, err = img.ExportJxl(&vips.JxlExportParams{
		//	Effort:   1,
		//	Tier:     4,
		//	Lossless: true,
		//	Distance: 1.0,
		//})
		buf, err = export2.ExportJxlWithParams(img, 0, true, 4, 1.0, 1)
	} else {
		//buf, _, err = img.ExportJxl(&vips.JxlExportParams{
		//	Effort:   1,
		//	Tier:     4,
		//	Quality:  quality,
		//	Lossless: false,
		//	Distance: 1.0,
		//})
		buf, err = export2.ExportJxlWithParams(img, quality, false, 4, 1.0, 1)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to encode image to JXL")
	}

	return buf, nil
}

func avifEncoder(img *vips.ImageRef, quality int, stripeMeta bool) ([]byte, error) {
	var (
		buf []byte
		err error
	)

	// If quality >= 100, we use lossless mode
	if quality >= 100 {
		//buf, _, err = img.ExportAvif(&vips.AvifExportParams{
		//	Lossless:      true,
		//	StripMetadata: stripeMeta,
		//})
		buf, err = export2.ExportAvifWithParams(img, stripeMeta, 0, 0, 0, true)
	} else {
		//buf, _, err = img.ExportAvif(&vips.AvifExportParams{
		//	Quality:       quality,
		//	Lossless:      false,
		//	StripMetadata: stripeMeta,
		//})
		buf, err = export2.ExportAvifWithParams(img, stripeMeta, quality, 0, 0, false)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to encode image to AVIF")
	}

	return buf, nil
}
