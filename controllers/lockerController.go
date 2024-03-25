package controllers

import (
	"final-project-acgm/database"
	"final-project-acgm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LockerCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := c.GetHeader("Content-Type")
	_, _ = db, contentType
	Locker := models.Locker{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Locker)
	} else {
		c.ShouldBind(&Locker)
	}

	err := db.Create(&Locker).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":     Locker.ID,
		"name":   Locker.Name,
		"status": Locker.Status,
		"price":  Locker.Price,
	})
}

func LockerList(c *gin.Context) {
	db := database.GetDB()
	lockers := []models.Locker{}

	err := db.Find(&lockers).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lockers": lockers,
	})
}

func LockerUpdate(c *gin.Context) {
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

	err := db.Model(&Locker).Where("id = ?", lockerId).Updates(models.Locker{
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

func LockerDelete(c *gin.Context) {
	db := database.GetDB()
	lockerId, _ := strconv.Atoi(c.Param("lockerId"))

	err := db.Debug().Delete(&models.Locker{}, lockerId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Locker deleted successfully",
	})
}
