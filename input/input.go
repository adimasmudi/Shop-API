package input

type RegisterUserInput struct {
	Nama         string `json:"nama" binding:"required"`
	KataSandi    string `json:"kataSandi" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	NoTelp       string `json:"no_telp" binding:"required"`
	TanggalLahir string `json:"tanggalLahir" binding:"required"`
	JenisKelamin string `json:"jenisKelamin" binding:"required"`
	Tentang      string `json:"tentang" binding:"required"`
	Pekerjaan    string `json:"pekerjaan" binding:"required"`
	IDProvinsi   int    `json:"idProvinsi" binding:"required"`
	IDKota       int    `json:"idKota" binding:"required"`
	IsAdmin      bool   `json:"isAdmin" binding:"required"`
}

type LoginUserInput struct {
	Email     string `json:"email" binding:"required,email"`
	KataSandi string `json:"kataSandi" binding:"required"`
}

// type UpdateProfileInput struct {
// 	Nama         string `json:"nama" `
// 	NoTelp       string `json:"no_telp" `
// 	TanggalLahir string `json:"tanggalLahir" `
// 	Tentang      string `json:"tentang" `
// 	Pekerjaan    string `json:"pekerjaan" `
// 	IDProvinsi   int    `json:"idProvinsi" `
// 	IDKota       int    `json:"idKota" `
// }