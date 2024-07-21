package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportJp2k(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportJp2k(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportJp2kWithQuality(img *vips.ImageRef, quality int, lossless bool) ([]byte, error) {
	params := vips.NewJp2kExportParams()
	params.Quality = quality
	params.Lossless = lossless
	buf, _, err := img.ExportJp2k(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// link: https://www.libvips.org/API/current/VipsForeignSave.html#vips-jp2ksave
// ExportJp2kWithParams exports the image to JP2K format.
// quality: the compression quality factor. The default value produces file with approximately the same size as regular JPEG Q 75.
// lossless: enable lossless compression.
// tileWidth: tile width, set the tile size. The default is 512.
// tileHeight: tile height, set the tile size. The default is 512.
func ExportJp2kWithParams(img *vips.ImageRef, quality int, lossless bool, tileWidth int, tileHeight int, subsampleMode vips.SubsampleMode) ([]byte, error) {
	params := vips.NewJp2kExportParams()
	params.Quality = quality
	params.Lossless = lossless
	params.TileWidth = tileWidth
	params.TileHeight = tileHeight
	params.SubsampleMode = subsampleMode

	buf, _, err := img.ExportJp2k(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
