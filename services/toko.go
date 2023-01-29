package services

import (
	"Shop-API/input"
	"Shop-API/model"
	"Shop-API/repository"
)

type ServiceToko interface {
	AddToko(id int,input input.AddToko) (model.Toko, error)
}

type serviceToko struct{
	repositoryToko repository.RepositoryToko
}

func NewServiceToko(repositoryToko repository.RepositoryToko) *serviceToko{
	return &serviceToko{repositoryToko}
}

