package controllers

import (
	"net/http"
	"time"

	"github.com/rizalnurizalludin/todo-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateActivityInput struct {
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UpdateActivityInput struct {
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	DeletedAt time.Time `json:"deleted_at"`
}

func FindActivities(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var activities []models.ActivityGroup
	db.Find(&activities)

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activities})
}

func CreateActivity(c *gin.Context) {
	// Validate input
	var input CreateActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Activity
	activity := models.ActivityGroup{Email: input.Email, Title: input.Title}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&activity)

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

func FindActivity(c *gin.Context) { // Get model if exist
	var activity models.ActivityGroup

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

func UpdateActivity(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var activity models.ActivityGroup
	if err := db.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	// Validate input
	var input UpdateActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ActivityGroup
	updatedInput.Email = input.Email
	updatedInput.Title = input.Title

	db.Model(&activity).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

func DeleteActivity(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var activity_group models.ActivityGroup
	if err := db.Where("id = ?", c.Param("id")).First(&activity_group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	db.Delete(&activity_group)

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "message": "Deleted Success", "data": true})
}
