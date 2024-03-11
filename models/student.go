package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email       string `gorm:"uniqueIndex"`
	IsSuspended bool
	Attendance  int
	StudentID   uint `json:"studentId"`
}
