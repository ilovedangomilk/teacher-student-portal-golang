package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Email    string    `gorm:"uniqueIndex"`
	Students []Student `gorm:"many2many:teacher_students;"`
}
