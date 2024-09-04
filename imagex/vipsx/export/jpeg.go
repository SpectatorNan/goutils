package export

import "github.com/davidbyttow/govips/v2/vips"

func ExportJpeg(img *vips.ImageRef) ([]byte, error) {
	buf, _, err := img.ExportJpeg(nil)
	if err != nil {
		return nil, err
	}
	return buf, nil

}

func ExportJpegWithQuality(img *vips.ImageRef, quality int, stripMetadata bool) ([]byte, error) {
	params := vips.NewJpegExportParams()
	params.StripMetadata = stripMetadata
	params.Quality = quality

	buf, _, err := img.ExportJpeg(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// link to: https://www.libvips.org/API/current/VipsForeignSave.html#vips-jpegsave
// ExportJpegWithParams exports the image to JPEG format.
// stripMetadata: strip all metadata from the image
// quality: quality factor
// optimizeCoding: compute optimal Huffman coding tables
// interlace: write an interlaced (progressive) jpeg
// subsampleMode:  VipsForeignSubsample, chroma subsampling mode
// trellisQuant: apply trellis quantisation to each 8x8 block
// overshootDeringing: overshoot samples with extreme values
// optimizeScans: split DCT coefficients into separate scans
// quantTable: quantization table index
func ExportJpegWithParams(img *vips.ImageRef, stripMetadata bool, quality int, optimizeCoding bool, interlace bool,
	subsampleMode vips.SubsampleMode, trellisQuant bool, overshootDeringing bool, optimizeScans bool, quantTable int) ([]byte, error) {
	params := vips.NewJpegExportParams()
	params.StripMetadata = stripMetadata
	params.Quality = quality
	params.Interlace = interlace
	params.OptimizeCoding = optimizeCoding
	params.SubsampleMode = subsampleMode
	params.TrellisQuant = trellisQuant
	params.OvershootDeringing = overshootDeringing
	params.OptimizeScans = optimizeScans
	params.QuantTable = quantTable

	buf, _, err := img.ExportJpeg(params)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
