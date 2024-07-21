package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportTiff(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportTiff(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExportTiffWithQuality(img *vips.ImageRef, quality int, compression vips.TiffCompression, predictor vips.TiffPredictor) ([]byte, error) {
	params := vips.NewTiffExportParams()
	params.Compression = compression
	params.Quality = quality
	params.Predictor = predictor

	buf, _, err := img.ExportTiff(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// link to: https://www.libvips.org/API/current/VipsForeignSave.html#vips-tiffsave
// ExportTiffWithParams exports the image to TIFF format.
func ExportTiffWithParams(img *vips.ImageRef, stripMetadata bool, quality int, compression vips.TiffCompression, predictor vips.TiffPredictor) ([]byte, error) {
	params := vips.NewTiffExportParams()
	params.StripMetadata = stripMetadata
	params.Compression = compression
	params.Quality = quality
	params.Predictor = predictor

	buf, _, err := img.ExportTiff(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
