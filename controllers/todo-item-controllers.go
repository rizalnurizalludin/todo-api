package controllers

import (
	"net/http"

	"github.com/rizalnurizalludin/todo-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTODOInput struct {
	ActivityGroupID uint   `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type UpdateTODOInput struct {
	ActivityGroupID uint   `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

func FindTODOS(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos []models.TODO
	db.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func CreateTODO(c *gin.Context) {
	// Validate input
	var input CreateTODOInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Priority == "" {
		input.Priority = "very_high"
	}

	// Create todo item
	todo := models.TODO{ActivityGroupID: input.ActivityGroupID, Title: input.Title, IsActive: input.IsActive, Priority: input.Priority}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func FindTODO(c *gin.Context) { // Get model if exist
	var todo models.TODO

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTODO(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var todo models.TODO
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	// Validate input
	var input UpdateTODOInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.TODO
	updatedInput.ActivityGroupID = input.ActivityGroupID
	updatedInput.Title = input.Title
	updatedInput.IsActive = input.IsActive
	updatedInput.Priority = input.Priority

	db.Model(&todo).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTODO(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.TODO
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Not Found", "message": "Activity with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "message": "Deleted Success", "data": true})
}
