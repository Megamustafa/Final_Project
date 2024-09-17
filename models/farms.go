package models

import (
	"time"

	"gorm.io/gorm"
)

type AquacultureFarms struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UserID    uint           `json:"user_id"`
	FarmID    uint           `json:"farm_id"`
	User      User           `json:"user"`
	Farm      Farm           `json:"farm"`
}

type FarmType struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
}

type Farm struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FarmTypeID  uint           `json:"farm_type_id"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	FarmType    FarmType       `json:"farm_type"`
}
