package media

import (
	"io"
	"mime/multipart"
	"os"
	"trainer-helper/utils"
)

func SaveFile(file *multipart.FileHeader) (name string, err error) {
	name = utils.RandomUUID()

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(name)
	if err != nil {
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	return
}
