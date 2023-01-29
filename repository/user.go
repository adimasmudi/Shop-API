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
	// UpdateProfile(user model.User) (model.User, error)
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

// func (r *repository) UpdateProfile(user model.User) (model.User, error){
// 	err := r.db.Save(&user).Error

// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
