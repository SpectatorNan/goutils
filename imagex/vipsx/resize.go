package vipsx

import (
	"github.com/SpectatorNan/goutils/imagex"
	"github.com/SpectatorNan/goutils/imagex/vipsx/export"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
)

func ResizeLocalImageFile(rawPath string, optimizedPath string, extraParams imagex.ExtraParams, extraParamsCropInteresting string) error {
	return resizeLocalImageFile(rawPath, optimizedPath, extraParams, extraParamsCropInteresting)
}

func resizeLocalImageFile(rawPath string, optimizedPath string, extraParams imagex.ExtraParams, extraParamsCropInteresting string) error {
	err := os.MkdirAll(path.Dir(optimizedPath), 0755)
	if err != nil {
		return errors.Wrapf(err, "failed to create directory for %s", optimizedPath)
	}
	buf, err := resizeLocalImage(rawPath, extraParams, extraParamsCropInteresting)
	if err != nil {
		return err
	}
	err = writeFile(optimizedPath, buf)
	if err != nil {
		return errors.Wrapf(err, "failed to write image to %s", optimizedPath)
	}
	return nil

}

func ResizeImageStream(reader io.Reader, extraParams imagex.ExtraParams, extraParamsCropInteresting string) ([]byte, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return resizeImageFromBuffer(buf, extraParams, extraParamsCropInteresting)
}
func resizeImageFromBuffer(imageBuf []byte, extraParams imagex.ExtraParams, extraParamsCropInteresting string) ([]byte, error) {
	img, err := vips.LoadImageFromBuffer(imageBuf, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load image from buffer")
	}
	defer img.Close()

	err = resizeImage(img, extraParams, extraParamsCropInteresting)

	buf, err := export.ExportJpeg(img)

	return buf, err
}

func resizeLocalImage(rawPath string, extraParams imagex.ExtraParams, extraParamsCropInteresting string) ([]byte, error) {
	img, err := vips.LoadImageFromFile(rawPath, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})
	if err != nil {
		return nil, err
	}
	defer img.Close()

	err = resizeImage(img, extraParams, extraParamsCropInteresting)

	format := img.Format()
	switch format {
	case vips.ImageTypeJPEG:
		break
	}
	buf, err := export.ExportJpeg(img)

	return buf, err
}

func resizeImage(img *vips.ImageRef, extraParams imagex.ExtraParams, extraParamsCropInteresting string) error {
	imageHeight := img.Height()
	imageWidth := img.Width()

	imgHeightWidthRatio := float32(imageHeight) / float32(imageWidth)

	// Here we have width, height and max_width, max_height
	// Both pairs cannot be used at the same time

	// max_height and max_width are used to make sure bigger images are resized to max_height and max_width
	// e.g, 500x500px image with max_width=200,max_height=100 will be resized to 100x100
	// while smaller images are untouched

	// If both are used, we will use width and height

	if extraParams.MaxHeight > 0 && extraParams.MaxWidth > 0 {
		// If any of it exceeds
		if imageHeight > extraParams.MaxHeight || imageWidth > extraParams.MaxWidth {
			// Check which dimension exceeds most
			heightExceedRatio := float32(imageHeight) / float32(extraParams.MaxHeight)
			widthExceedRatio := float32(imageWidth) / float32(extraParams.MaxWidth)
			// If height exceeds more, like 500x500 -> 200x100 (2.5 < 5)
			// Take max_height as new height ,resize and retain ratio
			if heightExceedRatio > widthExceedRatio {
				err := img.Thumbnail(int(float32(extraParams.MaxHeight)/imgHeightWidthRatio), extraParams.MaxHeight, 0)
				if err != nil {
					return err
				}
			} else {
				err := img.Thumbnail(extraParams.MaxWidth, int(float32(extraParams.MaxWidth)*imgHeightWidthRatio), 0)
				if err != nil {
					return err
				}
			}
		}
	}

	if extraParams.MaxHeight > 0 && imageHeight > extraParams.MaxHeight && extraParams.MaxWidth == 0 {
		err := img.Thumbnail(int(float32(extraParams.MaxHeight)/imgHeightWidthRatio), extraParams.MaxHeight, 0)
		if err != nil {
			return err
		}
	}

	if extraParams.MaxWidth > 0 && imageWidth > extraParams.MaxWidth && extraParams.MaxHeight == 0 {
		err := img.Thumbnail(extraParams.MaxWidth, int(float32(extraParams.MaxWidth)*imgHeightWidthRatio), 0)
		if err != nil {
			return err
		}
	}

	if extraParams.Width > 0 && extraParams.Height > 0 {
		var cropInteresting vips.Interesting
		switch extraParamsCropInteresting {
		case "InterestingNone":
			cropInteresting = vips.InterestingNone
		case "InterestingCentre":
			cropInteresting = vips.InterestingCentre
		case "InterestingEntropy":
			cropInteresting = vips.InterestingEntropy
		case "InterestingAttention":
			cropInteresting = vips.InterestingAttention
		case "InterestingLow":
			cropInteresting = vips.InterestingLow
		case "InterestingHigh":
			cropInteresting = vips.InterestingHigh
		case "InterestingAll":
			cropInteresting = vips.InterestingAll
		default:
			cropInteresting = vips.InterestingAttention
		}

		err := img.Thumbnail(extraParams.Width, extraParams.Height, cropInteresting)
		if err != nil {
			return err
		}
	}
	if extraParams.Width > 0 && extraParams.Height == 0 {
		err := img.Thumbnail(extraParams.Width, int(float32(extraParams.Width)*imgHeightWidthRatio), 0)
		if err != nil {
			return err
		}
	}
	if extraParams.Height > 0 && extraParams.Width == 0 {
		err := img.Thumbnail(int(float32(extraParams.Height)/imgHeightWidthRatio), extraParams.Height, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
