package imagex

import (
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
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

func GetImageSize(r io.Reader) (image.Config, string, error) {
	config, is, err := image.DecodeConfig(r)
	if err != nil {
		return config, "", err
	}
	return config, is, nil
}

// 检查文件是否为图片
func isImageFile(file multipart.File) bool {
	// 读取文件头的前512个字节以检查 MIME 类型
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		log.Println("Error reading file:", err)
		return false
	}

	// 将文件指针移回开头，以便后续处理文件时不会受影响
	file.Seek(0, 0)

	// 检测文件的 MIME 类型
	mimeType := http.DetectContentType(buffer)
	return strings.HasPrefix(mimeType, "image/")
}