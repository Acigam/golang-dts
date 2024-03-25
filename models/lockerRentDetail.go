package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type LockerRentDetail struct {
	ID           uint        `gorm:"primaryKey" json:"id" form:"id"`
	LockerRentID uint        `gorm:"not null" json:"locker_rent_id" form:"locker_rent_id" valid:"required"`
	LockerID     uint        `gorm:"not null" json:"locker_id" form:"locker_id" valid:"required"`
	Price        int         `gorm:"not null" json:"price" form:"price"`
	ReturnStatus string      `gorm:"not null" json:"return_status" form:"return_status" valid:"in(complete|key lost|not available)"`
	CreatedAt    *time.Time  `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    *time.Time  `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	LockerRent   *LockerRent `gorm:"foreignKey:LockerRentID" json:"lockerRent" form:"lockerRent"`
	Locker       *Locker     `gorm:"foreignKey:LockerID" json:"locker" form:"locker"`
}

func (u *LockerRentDetail) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *LockerRentDetail) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
