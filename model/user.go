package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"size:100;not null" json:"name"`
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `gorm:"not null" json:"password,omitempty"`
	Age        int    `json:"age"`
	RoleID     uint   `json:"roleId"`
	Role       Role   `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
