package main

import (
	"fmt"
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

	//check if folder exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0700)
		if err != nil {
			return errors.Wrap(err, "Could not create Folder")
		}
	}

	diro, _ := os.Getwd()
	fmt.Println(diro)
	fmt.Println(filepath.Join(dir, dst))
	//create tmp file
	out, err := os.Create("./tmp/test.txt")
	// out, err := os.Create(filepath.Join(dir, dst))
	if err != nil {
		return errors.Wrap(err, "Could not create file in expected dir")
	}
	defer out.Close()

	//copy content to new file
	_, err = io.Copy(out, src)
	if err != nil {
		return errors.Wrap(err, "Could not copy file")
	}

	return nil
}
