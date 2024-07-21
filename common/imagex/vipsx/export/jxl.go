package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportJxl(img *vips.ImageRef) ([]byte, error) {

	buf, _, err := img.ExportJxl(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportJxlWithQuality(img *vips.ImageRef, quality int, lossless bool) ([]byte, error) {

	params := vips.NewJxlExportParams()
	params.Quality = quality
	params.Lossless = lossless
	buf, _, err := img.ExportJxl(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// ExportJxlWithParams exports the image to JXL format.
// quality: 0-100, 100 means lossless
// lossless: enables lossless compression
// tier: decode speed tier
// distance: maximum encoding error
// effort:  encoding effort
func ExportJxlWithParams(img *vips.ImageRef, quality int, lossless bool, tier int, distance float64, effort int) ([]byte, error) {
	params := vips.NewJxlExportParams()
	params.Quality = quality
	params.Lossless = lossless
	params.Tier = tier
	params.Distance = distance
	params.Effort = effort
	buf, _, err := img.ExportJxl(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
