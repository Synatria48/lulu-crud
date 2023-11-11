package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Menangani file upload
	uploadFile, err := handleFileUpload(c, "uploadFoto")
	if err != nil {
		return err
	}
	u.UploadFoto = uploadFile

	// Menangani multiple file upload
	uploadFiles, err := handleMultipleFileUpload(c, "uploadDokumen")
	if err != nil {
		return err
	}
	u.UploadDokumen = uploadFiles

	db.Create(&u)
	return c.JSON(http.StatusCreated, u)
}

func ListUsers(c echo.Context) error {
	var u []User

	db.Find(&u)
	return c.JSON(http.StatusOK, u)
}
