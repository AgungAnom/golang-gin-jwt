package repository

import "golang-gin-jwt/entity"

type ProducRepository interface {
	FindByID(id int) *entity.Product
}