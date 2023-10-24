package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadFile ..
func UploadFile(file *multipart.FileHeader) (string, error) {

	// generate filename
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

	// generate path
	path := filepath.Join("uploads", filename)

	// save file to path
	err := SaveFile(path, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// SaveFile ..
func SaveFile(path string, file *multipart.FileHeader) error {

	// open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// create destination file
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// copy file
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
