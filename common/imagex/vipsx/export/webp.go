package export

import "github.com/davidbyttow/govips/v2/vips"

func EportWebp(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportWebp(nil)

	if err != nil {
		return nil, err
	}

	return buf, nil

}

func ExportWebpWithQuality(img *vips.ImageRef, stripMetadata bool, quality int, lossless bool, nearLossless bool, reductionEffort int) ([]byte, error) {
	params := vips.NewWebpExportParams()
	params.StripMetadata = stripMetadata
	params.Quality = quality
	params.Lossless = lossless
	params.NearLossless = nearLossless
	params.ReductionEffort = reductionEffort
	buf, _, err := img.ExportWebp(params)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// link: https://www.libvips.org/API/current/VipsForeignSave.html#vips-webpsave
// ExportWebpWithParams exports the image to WebP format.
// quality: It has the range 0 - 100, with the default 75.
func ExportWebpWithParams(img *vips.ImageRef, stripMetadata bool, quality int, lossless bool,
	nearLossless bool, reductionEffort int, iccProfile string, minSize bool, minKeyFrames int, maxKeyFrames int) ([]byte, error) {
	params := vips.NewWebpExportParams()
	params.Quality = quality
	params.Lossless = lossless
	params.StripMetadata = stripMetadata
	params.NearLossless = nearLossless
	params.ReductionEffort = reductionEffort
	params.IccProfile = iccProfile
	params.MinSize = minSize
	params.MinKeyFrames = minKeyFrames
	params.MaxKeyFrames = maxKeyFrames

	buf, _, err := img.ExportWebp(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
