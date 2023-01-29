package model

import "time"

type Toko struct {
	IdUser    int
	NamaToko  string
	UrlToko   string
	CreatedAt time.Time
	UpdatedAt time.Time
}