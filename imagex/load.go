package imagex

import (
	"io"
	"mime/multipart"
)

func fileLoadFromMultipartFile(file multipart.File) ([]byte, error) {
	// Ensure the file pointer is at the start
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	// Read all the data from the file into a byte slice
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
