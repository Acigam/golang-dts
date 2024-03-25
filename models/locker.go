package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type status string

const (
	AVAILABLE     status = "available"
	NOT_AVAILABLE status = "not available"
	KEY_LOST      status = "key lost"
)

type Locker struct {
	ID        uint       `gorm:"primaryKey" json:"id" form:"id"`
	Name      string     `gorm:"not null" json:"name" form:"name" valid:"required"`
	Status    status     `gorm:"not null" json:"status" form:"status" valid:"required, in(available|not available|key lost)"`
	Price     int        `gorm:"not null;default:10000" json:"price" form:"price"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autocreateTime" json:"updated_at,omitempty"`
}

func (u *Locker) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *Locker) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
