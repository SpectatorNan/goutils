package vipsx

import (
	"github.com/SpectatorNan/goutils/errors"
	"github.com/davidbyttow/govips/v2/vips"
	"io"
	"os"
	"path"
)

func ConvertLocalImage(rawPath, optimizedPath string, imageType string, quality int, stripeMeta bool) error {
	return convertLocalImage(rawPath, optimizedPath, imageType, quality, stripeMeta)
}

func ConvertImageStream(reader io.Reader, imageType string, quality int, stripeMeta bool) ([]byte, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return convertImageFromBuffer(buf, imageType, quality, stripeMeta)
}

func convertLocalImage(rawPath, optimizedPath string, imageType string, quality int, stripeMeta bool) error {

	err := os.MkdirAll(path.Dir(optimizedPath), 0755)
	if err != nil {
		return err
	}

	buf, err := convertLoadByLocal(rawPath, imageType, quality, stripeMeta)
	if err != nil {
		return err
	}

	err = writeFile(optimizedPath, buf)
	if err != nil {
		return err
	}
	return nil
}

func convertImageFromBuffer(imageBuf []byte, imageType string, quality int, stripeMeta bool) ([]byte, error) {
	img, err := vips.LoadImageFromBuffer(imageBuf, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load image from buffer")
	}
	defer img.Close()

	buf, err := convertImageFromImageRef(img, imageType, quality, stripeMeta)

	return buf, err
}

func writeFile(targetPath string, data []byte) error {
	err := os.WriteFile(targetPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func convertLoadByLocal(rawPath, imageType string, quality int, stripeMeta bool) ([]byte, error) {
	img, err := vips.LoadImageFromFile(rawPath, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})
	if err != nil {
		return nil, err
	}
	defer img.Close()

	buf, err := convertImageFromImageRef(img, imageType, quality, stripeMeta)
	return buf, err
}
func convertImageFromImageRef(img *vips.ImageRef, imageType string, quality int, stripeMeta bool) ([]byte, error) {
	// Pre-process image(auto rotate, resize, etc.)
	err := preProcessImage(img, imageType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to pre-process image")
	}
	//if resize != nil {
	//	err = resize(img)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//}
	switch imageType {
	case "webp":
		return webpEncoder(img, quality)
	case "avif":
		return avifEncoder(img, quality, stripeMeta)
	case "jxl":
		return jxlEncoder(img, quality)
	default:
		return nil, errors.Errorf("unsupported image type %s", imageType)
	}
}
