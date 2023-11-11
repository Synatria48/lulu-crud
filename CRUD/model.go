package crud

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name              string   `form:"name"`
	Nik               string   `form:"nik"`
	TanggalLahir      string   `form:"tanggalLahir"`
	Alamat            string   `form:"alamat"`
	Provinsi          string   `form:"provinsi"`
	Kabupaten         string   `form:"kabupaten"`
	Kecamatan         string   `form:"kecamatan"`
	NomorHp           string   `form:"nomorHp"`
	Email             string   `form:"email"`
	JenisKelamin      string   `form:"jenisKelamin"`
	TingkatPendidikan string   `form:"tingkatPendidikan"`
	UploadFoto        string   `form:"uploadFoto"`
	UploadDokumen     []string `form:"uploadDokumen"`
}
