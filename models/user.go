package models

import (
	"golang-gin-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
GormModel 			
FullName	string		`gorm:"not null" json:"full_name" form:"full_name" valid:"required~full name required"`
Email		string		`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email required,email~Invalid Email"`
Password	string		`gorm:"not null" json:"password" form:"password" valid:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
Role		string		`gorm:"not null" json:"role" form:"role" valid:"required~Role required"`
Product []Product		`gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"products"`
}


func (u *User) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil{
		err = errCreate
		return	
	}

	hashedPass, err := helpers.HashPass(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass
	
	return
}