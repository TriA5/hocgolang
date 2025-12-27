package repository

import (
	"be-ep/db"
	"be-ep/model"
	"be-ep/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *model.User) error {
	return db.DB.Create(user).Error
}

func (r *userRepository) FindAll(users *[]model.User) error {
	return db.DB.Preload("Role").Find(users).Error
}

func (r *userRepository) FindByID(id uint, user *model.User) error {
	return db.DB.Preload("Role").First(user, id).Error
}

func (r *userRepository) Update(user *model.User) error {
	return db.DB.Model(user).Updates(map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"age":      user.Age,
		"role_id":  user.RoleID,
	}).Error
}

func (r *userRepository) Delete(id uint) error {
	return db.DB.Delete(&model.User{}, id).Error
}
