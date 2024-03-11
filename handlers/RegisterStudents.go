package handlers

import (
	"errors"
	"net/http"

	"golang-project/models"
	"golang-project/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterStudents(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Utilize the models and the DB connection from util.ConnectDatabase()
	db := util.ConnectDatabase()

	var student models.Student
	if err := db.First(&student, input.StudentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Update the attendance for the found student
	student.Attendance++
	if err := db.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Attendance updated successfully"})
}
