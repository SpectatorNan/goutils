package imagex

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func Decode(r io.Reader) (image.Config, string, error) {
	return image.DecodeConfig(r)
}

func CheckImage(r io.Reader) error {
	_, _, err := image.Decode(r)
	return err
}

func CheckSupportImage(r io.Reader, supportTypes []string) (bool, error) {
	_, imgType, err := image.Decode(r)
	if err != nil {
		return false, err
	}
	for _, t := range supportTypes {
		if t == imgType {
			return true, nil
		}
	}

	return false, nil
}
