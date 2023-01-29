package formatter

import "Shop-API/model"

type UserFormatter struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	Email string `json:"email"`
	JenisKelamin string `json:"jenisKelamin"`
	NoTelp string `json:"nomerTelepon"`
	TanggalLahir string `json:"tanggalLahir"`
	Tentang string `json:"tentang"`
	IsAdmin  bool `json:"isAdmin"`
	IDProvinsi  int `json:"idProvinsi"`
	IDKota   int `json:"idKota"`
	Pekerjaan  string `json:"pekerjan"`
	Token  string `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter {
	formatter := UserFormatter{
		Id : user.Id,
		Nama:       user.Nama,
		Pekerjaan : user.Pekerjaan,
		Email:      user.Email,
		JenisKelamin : user.JenisKelamin ,
		NoTelp : user.NoTelp,
		TanggalLahir : user.TanggalLahir, 
		Tentang : user.Tentang, 
		IsAdmin : user.IsAdmin, 
		IDProvinsi : user.IDProvinsi, 
		IDKota : user.IDKota,
		Token:      token,
	}


	return formatter
}
