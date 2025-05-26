package media

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"trainer-helper/utils"
)

func SaveFile(file *multipart.FileHeader, rootFilePath string) (name string, err error) {
	name = utils.RandomUUID()

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join(rootFilePath, name))
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
