package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Post struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of your post is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your Post is required"`
	UserID      uint
	User        *User
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
