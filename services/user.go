package services

import (
	"Shop-API/input"
	"Shop-API/model"
	"Shop-API/repository"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



type Service interface {
	RegisterUser(input input.RegisterUserInput) (model.User, string, error)
	Login(input input.LoginUserInput) (model.User, string, error)
	GetProfile(tokenString string) (model.User, error)
	UpdateProfile(id int, input input.UpdateProfileInput) (model.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input input.RegisterUserInput) (model.User, string, error) {
	user := model.User{}
	user.Nama = input.Nama
	user.Email = input.Email
	user.JenisKelamin = input.JenisKelamin
	user.NoTelp = input.NoTelp
	user.TanggalLahir = input.TanggalLahir
	user.Tentang = input.Tentang
	user.IsAdmin = input.IsAdmin
	user.IDProvinsi = input.IDProvinsi
	user.IDKota = input.IDKota
	user.Pekerjaan = input.Pekerjaan

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.KataSandi), bcrypt.MinCost)

	if err != nil {
		return user,"", err
	}

	user.KataSandi = string(passwordHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer" : strconv.Itoa(int(newUser.Id)),
		"ExpiresAt" : time.Now().Add(time.Hour * 48).Unix(), // 1 day
		
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return user, "", err
	}

	return newUser, tokenString, nil
}

func (s *service) Login(input input.LoginUserInput) (model.User, string, error) {
	email := input.Email
	password := input.KataSandi

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, "",err
	}


	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(password))

	if err != nil {
		return user, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer" : strconv.Itoa(int(user.Id)),
		"ExpiresAt" : time.Now().Add(time.Hour * 24).Unix(), // 1 day
		"Name" : user.Nama,
		
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return user, "", err
	}


	return user, tokenString, nil

}

func (s *service) GetProfile(tokenString string)(model.User, error){

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	var user model.User


	claims, ok := token.Claims.(jwt.MapClaims)

	if (!(ok && token.Valid)){
		return user, err
	}

	id := claims["Issuer"].(string)
		
	user, err = s.repository.FindById(id)

	return user, err
}

func (s *service) UpdateProfile(id int, input input.UpdateProfileInput)(model.User, error){
	userUpdate := new(model.User)

	var user model.User

	userUpdate.Nama = input.Nama
	userUpdate.NoTelp = input.NoTelp
	userUpdate.TanggalLahir = input.TanggalLahir
	userUpdate.Tentang = input.Tentang
	userUpdate.IDProvinsi = input.IDProvinsi
	userUpdate.IDKota = input.IDKota
	userUpdate.Pekerjaan = input.Pekerjaan


	updatedUser, err := s.repository.UpdateProfile(id,user,userUpdate)

	if err != nil{
		return updatedUser, err
	}

	return updatedUser, nil
}

