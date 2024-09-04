package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportHeif(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportHeif(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportHeifWithQuality(img *vips.ImageRef, quality int, lossless bool) ([]byte, error) {

	params := vips.NewHeifExportParams()
	params.Quality = quality
	params.Lossless = lossless

	buf, _, err := img.ExportHeif(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// link to: https://www.libvips.org/API/current/VipsForeignSave.html#vips-heifsave
func ExportHeifWithParams(img *vips.ImageRef, quality int, bitdepth int, effort int, lossless bool) ([]byte, error) {

	params := vips.NewHeifExportParams()
	params.Quality = quality
	params.Bitdepth = bitdepth
	params.Effort = effort
	params.Lossless = lossless

	buf, _, err := img.ExportHeif(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
