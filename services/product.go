package services

import (
	"errors"
	"golang-gin-jwt/entity"
	"golang-gin-jwt/repository"
)

type ProductService struct {
	Repository repository.ProducRepository
}

func (service ProductService) GetOneProduct(id int)(*entity.Product, error){
	product := service.Repository.FindByID(id)

	if product == nil {
		return nil, errors.New("Product not Found")
	}

	return product, nil
}