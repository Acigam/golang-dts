package models

import (
	"final-project-acgm/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id" form:"id"` // form:"id" is used to bind the form data to the struct
	Username  string     `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Age       int        `gorm:"not null" json:"age" form:"age" valid:"required,range(1|130)"`
	IsAdmin   bool       `gorm:"not null;default:false" json:"is_admin" form:"is_admin"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autocreateTime" json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
