package services

import (
	"Shop-API/input"
	"Shop-API/model"
	"Shop-API/repository"
	"os"

	"github.com/golang-jwt/jwt"
)

type ServiceToko interface {
	GetMyToko(tokenString string) (model.Toko, model.User, error)
	GetAllToko() ([]model.Toko, error)
	UpdateToko(id int, input input.UpdateTokoInput) (model.Toko, error)
}

type serviceToko struct{
	repositoryToko repository.RepositoryToko
}

func NewServiceToko(repositoryToko repository.RepositoryToko) *serviceToko{
	return &serviceToko{repositoryToko}
}

func (s *serviceToko) GetMyToko(tokenString string)(model.Toko, model.User, error){
	// JWT Claim
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	var user model.User


	claims, ok := token.Claims.(jwt.MapClaims)

	if (!(ok && token.Valid)){
		return model.Toko{},user, err
	}

	id := claims["Issuer"].(string)
		
	// get user
	user, err = s.repositoryToko.FindUserById(id)

	if err != nil{
		return model.Toko{},user, err
	}

	// get toko
	toko, err := s.repositoryToko.FindTokoByIdUser(id)
	if err != nil{
		return toko,user, err
	}

	return toko, user, nil
}

func (s *serviceToko) GetAllToko() ([]model.Toko, error){
	// get toko
	toko, err := s.repositoryToko.GetAllToko()

	if err != nil{
		return toko, err
	}

	return toko, err
}

func (s *serviceToko) UpdateToko(id int, input input.UpdateTokoInput) (model.Toko, error){
	tokoUpdate := new(model.Toko)

	var toko model.Toko

	tokoUpdate.NamaToko = input.NamaToko
	tokoUpdate.UrlToko = input.UrlToko

	updatedToko, err := s.repositoryToko.UpdateToko(id, toko, tokoUpdate)

	if err != nil{
		return updatedToko, err
	}

	return updatedToko, nil
}

