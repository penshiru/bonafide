package main

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

//SaveUploadedFile saves a tmp copy in order for the file to be processed.
func SaveUploadedFile(file *multipart.FileHeader, dst string, dir string) error {
	src, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "Could not Open file")
	}
	defer src.Close()

	out, err := os.Create(filepath.Join(dir, dst))
	if err != nil {
		return errors.Wrap(err, "Could not create file in expected dir")
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return errors.Wrap(err, "Could not copy file")
	}

	return nil
}
