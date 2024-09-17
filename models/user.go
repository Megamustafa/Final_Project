package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"-"`
	FullName    string         `json:"fullname"`
	Address     string         `json:"address"`
	PhoneNumber string         `json:"phone_number"`
}
