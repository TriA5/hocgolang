package repository

import "be-ep/model"

type UserRepository interface {
	Create(user *model.User) error
	FindAll(users *[]model.User) error
	FindByID(id uint, user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}
