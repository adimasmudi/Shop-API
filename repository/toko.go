package repository

import (
	"Shop-API/model"
	"strconv"

	"gorm.io/gorm"
)

type RepositoryToko interface {
	FindUserById(id string)(model.User, error)
	FindTokoByIdUser(id string)(model.Toko, error)
	UpdateToko(id int, toko model.Toko, tokoUpdate interface{}) (model.Toko, error)
}

type repositoryToko struct{
	db *gorm.DB
}

func NewRepositoryToko(db *gorm.DB) *repositoryToko{
	return &repositoryToko{db}
}

func (r *repositoryToko) FindUserById(id string) (model.User,  error){
	var user model.User
	intId, err :=  strconv.Atoi(id)

	if err != nil{
		return user, err
	}

	err2 := r.db.Where("id = ?",intId).Find(&user).Error

	if err2 != nil{
		return user, err2
	}

	return user, nil
}

func (r *repositoryToko) FindTokoByIdUser(id string) (model.Toko,  error){
	var toko model.Toko
	intId, err :=  strconv.Atoi(id)

	if err != nil{
		return toko, err
	}

	err2 := r.db.Where("id_user = ?",intId).Find(&toko).Error

	if err2 != nil{
		return toko, err2
	}

	return toko, nil
}

func (r *repositoryToko) UpdateToko(id int, toko model.Toko, tokoUpdate interface{}) (model.Toko, error){
	r.db.Find(&toko, "id=?",id)

	err := r.db.Where("id=?",id).Updates(tokoUpdate).Error

	if err != nil {
		return toko, err
	}

	return toko, err
}
