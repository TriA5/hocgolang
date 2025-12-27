package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"size:100;not null" json:"name"`

	Users []User `gorm:"foreignKey:RoleID" json:"-"`
}
