package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type LockerRent struct {
	ID           uint       `gorm:"primaryKey" json:"id" form:"id"`
	TenantName   string     `gorm:"not null" json:"tenant_name" form:"tenant_name" valid:"required"`
	UserID       uint       `gorm:"not null" json:"user_id"`
	SubTotal     int        `json:"sub_total" form:"sub_total"`
	PaidAmount   int        `gorm:"not null" json:"paid_amount" form:"paid_amount" valid:"required"`
	ChangeAmount int        `json:"change_amount" form:"change_amount"`
	ReturnTime   *time.Time `json:"return_time" form:"return_time"`
	CreatedAt    *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	User         *User      `gorm:"foreignKey:UserID" json:"user" form:"user"`
}

func (u *LockerRent) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *LockerRent) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
