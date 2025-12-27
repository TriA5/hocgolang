package handler

import (
	"net/http"
	"strconv"

	"be-ep/model"
	repository "be-ep/repository/impl"

	"github.com/gin-gonic/gin"
)

var userRepo = repository.NewUserRepository()

// tạo user
func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Load lại user với role để trả về response đầy đủ
	if err := userRepo.FindByID(user.ID, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// lấy tất cả user
func GetUsers(c *gin.Context) {
	var users []model.User

	if err := userRepo.FindAll(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// lấy user theo ID
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User

	if err := userRepo.FindByID(uint(id), &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// cập nhật user
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updateData model.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID từ param
	updateData.ID = uint(id)

	// Update user
	if err := userRepo.Update(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Load lại user với role để trả về
	var user model.User
	if err := userRepo.FindByID(uint(id), &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// xóa user
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := userRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
