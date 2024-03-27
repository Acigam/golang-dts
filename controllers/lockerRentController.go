package controllers

import (
	"final-project-acgm/database"
	"final-project-acgm/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LockerRentRequest struct {
	models.LockerRent
	LockerRentDetails []models.LockerRentDetail `json:"lockerRentDetails"`
}

// {
// 	"tenant_name": "John Doe",
// 	"paid_amount": 10000,
// 	"lockerRentDetails": [
// 	  {
// 		"locker_id": 1,
// 	  }
// 	]
// }

func LockerRentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := c.GetHeader("Content-Type")

	LockerRentRequest := LockerRentRequest{}
	LockerRentRequest.UserID = uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&LockerRentRequest)
	} else {
		c.ShouldBind(&LockerRentRequest)
	}

	tx := db.Begin()

	err := tx.Debug().Create(&LockerRentRequest.LockerRent).Error

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, detail := range LockerRentRequest.LockerRentDetails {
		detail.LockerRentID = LockerRentRequest.LockerRent.ID
		if err := tx.Debug().Create(&detail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"id":                LockerRentRequest.LockerRent.ID,
		"tenant_name":       LockerRentRequest.LockerRent.TenantName,
		"user_id":           LockerRentRequest.LockerRent.UserID,
		"sub_total":         LockerRentRequest.LockerRent.SubTotal,
		"paid_amount":       LockerRentRequest.LockerRent.PaidAmount,
		"change_amount":     LockerRentRequest.LockerRent.ChangeAmount,
		"return_time":       LockerRentRequest.LockerRent.ReturnTime,
		"lockerRentDetails": LockerRentRequest.LockerRentDetails,
	})
}

func LockerRentList(c *gin.Context) {
	db := database.GetDB()
	lockerRents := []models.LockerRent{}

	err := db.Debug().Find(&lockerRents).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locker_rents": lockerRents,
	})
}

func LockerRentFinish(c *gin.Context) {
	db := database.GetDB()

	locker_rent_id := c.Param("lockerRentId")

	var lockerRent models.LockerRent
	if err := db.Where("id = ?", locker_rent_id).First(&lockerRent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "LockerRent not found",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	if lockerRent.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Forbidden",
		})
		return
	}

	var lockerRentDetails []models.LockerRentDetail
	if err := db.Where("locker_rent_id = ?", locker_rent_id).Find(&lockerRentDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch LockerRentDetails",
		})
		return
	}

	for _, detail := range lockerRentDetails {
		if detail.ReturnStatus == "complete" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "LockerRent already finished",
			})
			return
		}
	}

	tx := db.Begin()

	for _, detail := range lockerRentDetails {
		detail.ReturnStatus = "complete"
		if err := tx.Debug().Save(&detail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not update LockerRentDetails",
			})
			return
		}
	}

	now := time.Now()
	lockerRent.ReturnTime = &now
	if err := tx.Debug().Save(&lockerRent).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update LockerRent",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"locker_rent": lockerRent,
	})

}

func LockerRentCancel(c *gin.Context) {
	db := database.GetDB()

	locker_rent_id := c.Param("lockerRentId")

	var lockerRent models.LockerRent
	if err := db.Where("id = ?", locker_rent_id).First(&lockerRent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "LockerRent not found",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	if lockerRent.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Forbidden",
		})
		return
	}

	var lockerRentDetails []models.LockerRentDetail
	if err := db.Where("locker_rent_id = ?", locker_rent_id).Find(&lockerRentDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch LockerRentDetails",
		})
		return
	}

	for _, detail := range lockerRentDetails {
		if detail.ReturnStatus == "complete" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "LockerRent already finished",
			})
			return
		}
	}

	tx := db.Begin()

	for _, detail := range lockerRentDetails {
		detail.ReturnStatus = "not available"
		if err := tx.Debug().Save(&detail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not update LockerRentDetails",
			})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"locker_rent": lockerRent,
	})
}
