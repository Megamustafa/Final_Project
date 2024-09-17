package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
}

type Product struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ProductTypeID uint           `json:"product_type_id"`
	Description   string         `json:"description"`
	Price         int            `json:"price"`
	ProductType   ProductType    `json:"product_type"`
}
