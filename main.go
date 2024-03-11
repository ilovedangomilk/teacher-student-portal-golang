package main

import (
	"golang-project/handlers"
	"golang-project/util"

	"github.com/gin-gonic/gin"
	"github.com/ilovedangomilk/teacher-student-portal-golang/models"
)

func main() {
	db := util.ConnectDatabase()
	// Automigrate your models
	db.AutoMigrate(&models.Teacher{}, &models.Student{})

	r := gin.Default()

	// Define your routes here
	r.POST("/register", handlers.RegisterStudents)

	// Add a GET route to retrieve a student by ID
	r.GET("/students/:id", handlers.GetStudentByID)

	// More routes...

	r.Run() // listen and serve on 0.0.0.0:8080
}
