package crud

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        `json:"-"`
	ID                uint64 `json:"-"`
	Name              string `json:"name"`
	Nik               string `json:"nik"`
	TanggalLahir      string `json:"tanggalLahir"`
	Alamat            string `json:"alamat"`
	Provinsi          string `json:"provinsi"`
	Kabupaten         string `json:"kabupaten"`
	Kecamatan         string `json:"kecamatan"`
	NomorHp           string `json:"nomorHp"`
	Email             string `json:"email"`
	JenisKelamin      string `json:"jenisKelamin"`
	TingkatPendidikan string `json:"tingkatPendidikan"`
	UploadFoto        string `json:"uploadFoto"`
	UploadDokumen     string `json:"uploadDokumen"`
}

type DTO struct {
	Name              string   `form:"nama"`
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
	UploadDokumen     []string `form:"uploadDokumen[]"`
}
