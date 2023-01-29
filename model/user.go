package model

import "time"

type User struct {
	Id int
	Nama         string 
	KataSandi    string 
	Email        string 
	NoTelp       string
	TanggalLahir string
	JenisKelamin string
	Tentang      string
	Pekerjaan    string 
	IDProvinsi   int    
	IDKota       int    
	IsAdmin      bool  
	CreatedAt      time.Time
	UpdatedAt      time.Time
}