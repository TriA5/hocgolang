package repository

import "be-ep/model"

type RoleRepository interface {
	Create(role *model.Role) error
	FindAll(roles *[]model.Role) error
	FindByID(id uint, role *model.Role) error
	Update(role *model.Role) error
	Delete(id uint) error
}
