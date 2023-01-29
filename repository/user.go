package repository

import (
	"Shop-API/model"
	"strconv"

	"gorm.io/gorm"
)

type Repository interface{
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id string) (model.User, error)
	UpdateProfile(id int, user model.User, userUpdate interface{}) (model.User, error)
	SaveToko(toko model.Toko) (model.Toko, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Save(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (model.User,  error){
	var user model.User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil{
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(id string) (model.User,  error){
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

func (r *repository) UpdateProfile(id int, user model.User, userUpdate interface{}) (model.User, error){
	r.db.Find(&user, "id=?",id)
	err := r.db.Where("id=?",id).Updates(userUpdate).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) SaveToko(toko model.Toko) (model.Toko, error) {
	err := r.db.Create(&toko).Error

	if err != nil {
		return toko, err
	}

	return toko, nil
}
