package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportGif(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportGIF(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil

}

func ExportGifWithQuality(img *vips.ImageRef, quality int, effort int, bitdepth int) ([]byte, error) {
	params := vips.NewGifExportParams()
	params.Quality = quality
	params.Effort = effort
	params.Bitdepth = bitdepth
	buf, _, err := img.ExportGIF(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportGifWithParams(img *vips.ImageRef, stripMetadata bool, quality int, dither float64, effort int, bitdepth int) ([]byte, error) {
	params := vips.NewGifExportParams()
	params.StripMetadata = stripMetadata
	params.Quality = quality
	params.Dither = dither
	params.Effort = effort
	params.Bitdepth = bitdepth
	buf, _, err := img.ExportGIF(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
