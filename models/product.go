package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
GormModel 			
Title		string		`json:"title" form:"title" valid:"required~title required"`
Description	string		`json:"description" form:"description" valid:"required~description required"`
UserID		uint		
User *User
}


func (p *Product) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil{
		err = errCreate
		return	
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(projectDB *gorm.DB) (err error){
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil{
		err = errUpdate
		return	
	}

	err = nil
	return
}