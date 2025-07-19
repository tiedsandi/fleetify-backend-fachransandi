package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func CreateDepartment(c *gin.Context) {
	var dept models.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Input tidak valid: " + err.Error()})
		return
	}

	if err := config.DB.Create(&dept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data departemen"})
		return
	}

	c.JSON(http.StatusCreated, dept)
}

func GetDepartments(c *gin.Context) {
	var departments []models.Department
	if err := config.DB.Find(&departments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data departemen"})
		return
	}
	c.JSON(http.StatusOK, departments)
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")
	var dept models.Department

	if err := config.DB.First(&dept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departemen tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, dept)
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var dept models.Department

	if err := config.DB.First(&dept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departemen tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	if err := config.DB.Save(&dept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data departemen"})
		return
	}

	c.JSON(http.StatusOK, dept)
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var dept models.Department

	if err := config.DB.First(&dept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departemen tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&dept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus departemen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Departemen berhasil dihapus"})
}
