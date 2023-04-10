package repository

import (
	"golang-gin-jwt/entity"

	"github.com/stretchr/testify/mock"
)

type ProducRepositoryMock struct {
	Mock mock.Mock
}

func (p *ProducRepositoryMock) FindByID(id int) *entity.Product{
	argument := p.Mock.Called(id)

	if argument.Get(0) == nil {
		return nil
	}

	product := argument.Get(0).(entity.Product)
	return &product
}