package repository

import (
	"be-ep/db"
	"be-ep/model"
	"be-ep/repository"
)

type roleRepository struct{}

func NewRoleRepository() repository.RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) Create(role *model.Role) error {
	return db.DB.Create(role).Error
}

func (r *roleRepository) FindAll(roles *[]model.Role) error {
	return db.DB.Find(roles).Error
}

func (r *roleRepository) FindByID(id uint, role *model.Role) error {
	return db.DB.First(role, id).Error
}

func (r *roleRepository) Update(role *model.Role) error {
	return db.DB.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return db.DB.Delete(&model.Role{}, id).Error
}
