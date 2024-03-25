package controllers

import (
	"final-project-acgm/database"
	"final-project-acgm/models"
	"net/http"
	"strconv"

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
	contentType := c.GetHeader("Content-Type")
	_, _ = db, contentType
	Locker := models.Locker{}

	lockerId, _ := strconv.Atoi(c.Param("lockerId"))
	Locker.ID = uint(lockerId)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Locker)
	} else {
		c.ShouldBind(&Locker)
	}

	err := db.Debug().Model(&Locker).Where("id = ?", lockerId).Updates(models.Locker{
		Name:   Locker.Name,
		Status: Locker.Status,
		Price:  Locker.Price,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Locker)
}

func LockerRentCancel(c *gin.Context) {

}
