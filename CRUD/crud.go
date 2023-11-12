package crud

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	u := DTO{}

	if err := c.Bind(&u); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(u)

	dbU := User{}

	dbU.Name = u.Name
	dbU.Nik = u.Nik
	dbU.TanggalLahir = u.TanggalLahir
	dbU.Alamat = u.Alamat
	dbU.Provinsi = u.Provinsi
	dbU.Kabupaten = u.Kabupaten
	dbU.Kecamatan = u.Kecamatan
	dbU.NomorHp = u.NomorHp
	dbU.Email = u.Email
	dbU.JenisKelamin = u.JenisKelamin
	dbU.TingkatPendidikan = u.TingkatPendidikan

	// Menangani file upload
	uploadFile, err := handleFileUpload(c, "uploadFoto")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(uploadFile)
	dbU.UploadFoto = uploadFile

	// Menangani multiple file upload
	uploadFiles, err := handleMultipleFileUpload(c, "uploadDokumen[]")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(strings.Join(uploadFiles, ","))
	dbU.UploadDokumen = strings.Join(uploadFiles, ",")

	err = db.Create(&dbU).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func ListUsers(c echo.Context) error {
	var u []User

	db.Find(&u)
	return c.JSON(http.StatusOK, u)
}
