package handler

import (
	"net/http"
	"strconv"

	"be-ep/model"
	repository "be-ep/repository/impl"

	"github.com/gin-gonic/gin"
)

var roleRepo = repository.NewRoleRepository()

// tạo role
func CreateRole(c *gin.Context) {
	var role model.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := roleRepo.Create(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// lấy tất cả role
func GetRoles(c *gin.Context) {
	var roles []model.Role

	if err := roleRepo.FindAll(&roles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// lấy role theo ID
func GetRoleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role model.Role

	if err := roleRepo.FindByID(uint(id), &role); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// cập nhật role
func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role model.Role

	if err := roleRepo.FindByID(uint(id), &role); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := roleRepo.Update(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// xóa role
func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := roleRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
