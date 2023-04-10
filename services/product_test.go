package services

import (
	"golang-gin-jwt/entity"
	"golang-gin-jwt/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProducRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}




func TestProductServiceGetOneProductNotFound(t *testing.T){
	productRepository.Mock.On("FindByID", 1).Return(nil)
	product, err := productService.GetOneProduct(1)


	assert.Nil(t,product)
	assert.NotNil(t,err)
	assert.Equal(t, "Product not Found", err.Error(), "Error response has to be 'Product not Found'")
}

func TestProductServiceGetOneProductFound(t *testing.T){
	data := entity.Product{
	ID: 1,
	Title: "Title",
	Description: "Desc", 
	UserID : 1,
	}
	productRepository.Mock.On("FindByID", 2).Return(data)
	product, err := productService.GetOneProduct(2)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, data.ID, product.ID)
	assert.Equal(t, data.Title, product.Title)
	assert.Equal(t, data.Description, product.Description)
	assert.Equal(t, data.UserID, product.UserID)
}

func TestProductServiceGetAllProductFound(t *testing.T){
	data := []entity.Product{
		{	
		ID: 2,
		Title: "Title",
		Description: "Desc", 
		UserID : 2,
		},
		{	
		ID: 3,
		Title: "Title",
		Description: "Desc", 
		UserID : 3,
		},
		{	
		ID: 4,
		Title: "Title",
		Description: "Desc", 
		UserID : 4,
		},
	}
	for i, x := range data {
	productRepository.Mock.On("FindByID", i+3).Return(x)
	product, err := productService.GetOneProduct(i+3)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, x.ID, product.ID)
	assert.Equal(t, x.Title, product.Title)
	assert.Equal(t, x.Description, product.Description)
	assert.Equal(t, x.UserID, product.UserID)
	}
}


func TestProductServiceGetAllProductNotFound(t *testing.T){
	data := []entity.Product{
		{	
		ID: 2,
		Title: "Title",
		Description: "Desc", 
		UserID : 2,
		},
		{	
		ID: 3,
		Title: "Title",
		Description: "Desc", 
		UserID : 3,
		},
		{	
		ID: 4,
		Title: "Title",
		Description: "Desc", 
		UserID : 4,
		},
	}
	for i := range data {
	productRepository.Mock.On("FindByID", i+6).Return(nil)
	product, err := productService.GetOneProduct(i+6)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not Found", err.Error(), "Error response has to be 'Product not Found'")
	}
}