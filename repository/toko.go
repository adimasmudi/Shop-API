package repository

import (
	"Shop-API/model"

	"gorm.io/gorm"
)

type RepositoryToko interface {
	SaveToko(toko model.Toko) (model.Toko, error)
}

type repositoryToko struct{
	db *gorm.DB
}

func NewRepositoryToko(db *gorm.DB) *repositoryToko{
	return &repositoryToko{db}
}

