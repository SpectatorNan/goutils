package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportAvif(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportAvif(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportAvifWithQuality(img *vips.ImageRef, quality int, bitdepth int, effort int, lossless bool) ([]byte, error) {
	params := vips.NewAvifExportParams()
	params.Quality = quality
	params.Bitdepth = bitdepth
	params.Effort = effort
	params.Lossless = lossless
	buf, _, err := img.ExportAvif(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportAvifWithParams(img *vips.ImageRef, stripMetadata bool, quality int, bitdepth int, effort int, lossless bool) ([]byte, error) {
	params := vips.NewAvifExportParams()
	params.StripMetadata = stripMetadata
	params.Quality = quality
	params.Bitdepth = bitdepth
	params.Effort = effort
	params.Lossless = lossless
	buf, _, err := img.ExportAvif(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
