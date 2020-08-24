package entity

import (
	"time"
)

type Video struct {
	ID          uint64    `json:id gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=10" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json : "description" binding:"max=20" gorm:"type:varchar(250)"`
	URL         string    `json : "url" binding:"required,url" gorm:"type:varchar(200)"`
	Author      Person    `json:"author" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json: "-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:updated_at gorm:"default:CURRENT_TIMESTAMP" `
}

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding : "required" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" binding : "required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int16  `json : "age" binding : "gte=1,lte=130" `
}
