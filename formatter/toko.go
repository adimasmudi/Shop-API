package formatter

import "Shop-API/model"

type tokoFormatter struct {
	Id       int         `json:"id"`
	NamaToko string      `json:"nama_toko"`
	UrlToko  string      `json:"url_toko"`
	User     interface{} `json:"user"`
}

func FormatToko(toko model.Toko, user model.User) tokoFormatter {
	formatter := tokoFormatter{
		Id:       toko.Id,
		NamaToko: toko.NamaToko,
		UrlToko:  toko.UrlToko,
		User:     user,
	}

	return formatter
}