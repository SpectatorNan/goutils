package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportPng(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportPng(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportPngWithQuality(img *vips.ImageRef, quality int, compression int, filter vips.PngFilter, interlace bool, palette bool) ([]byte, error) {
	params := vips.NewPngExportParams()
	params.Compression = compression
	params.Filter = filter
	params.Interlace = interlace
	params.Quality = quality
	params.Palette = palette
	buf, _, err := img.ExportPng(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// link to:https://www.libvips.org/API/current/VipsForeignSave.html#vips-pngsave
func ExportPNGWithParams(img *vips.ImageRef, stripMetadata bool, compression int, filter vips.PngFilter, interlace bool,
	quality int, palette bool, dither float64, bitdepth int, profile string) ([]byte, error) {
	params := vips.NewPngExportParams()
	params.StripMetadata = stripMetadata
	params.Compression = compression
	params.Filter = filter
	params.Interlace = interlace
	params.Quality = quality
	params.Palette = palette
	params.Dither = dither
	params.Bitdepth = bitdepth
	params.Profile = profile
	buf, _, err := img.ExportPng(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
