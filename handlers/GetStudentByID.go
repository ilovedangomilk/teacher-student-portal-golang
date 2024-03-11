package handlers

import (
	"net/http"
	"strconv" // Import strconv for converting string to uint

	"golang-project/models"
	"golang-project/util"

	"github.com/gin-gonic/gin"
)

// GetStudentByID retrieves a student by their StudentID and displays their attendance.
func GetStudentByID(c *gin.Context) {
	// Connect to the database
	db := util.ConnectDatabase()

	// Get the StudentID from the URL parameter and convert it to uint
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID format"})
		return
	}

	var student models.Student
	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// If the student is found, return their data
	c.JSON(http.StatusOK, gin.H{
		"studentId":   student.StudentID,
		"email":       student.Email,
		"isSuspended": student.IsSuspended,
		"attendance":  student.Attendance,
	})
}
