package crud

import (
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

func handleFileUpload(c echo.Context, formFieldName string) (string, error) {
	file, err := c.FormFile(formFieldName)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Tentukan nama file untuk penyimpanan
	fileName := filepath.Join("dokumen/foto", file.Filename)

	dst, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Salin isi file dari src ke dst
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return file.Filename, nil
}

// handleMultipleFileUpload menangani multiple file upload dan mengembalikan daftar nama file yang disimpan
func handleMultipleFileUpload(c echo.Context, formFieldName string) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[formFieldName]
	var fileNames []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Tentukan nama file untuk penyimpanan
		fileName := filepath.Join("dokumen/uploads", file.Filename)
		fileNames = append(fileNames, file.Filename)

		dst, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Salin isi file dari src ke dst
		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}
	}

	return fileNames, nil
}
