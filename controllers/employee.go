package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func CreateEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	var dept models.Department
	if err := config.DB.First(&dept, emp.DepartmentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID departemen tidak valid"})
		return
	}

	var existing models.Employee
	if err := config.DB.Where("employee_id = ?", emp.EmployeeID).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID karyawan sudah digunakan"})
		return
	}

	if err := config.DB.Create(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data karyawan"})
		return
	}

	config.DB.Preload("Department").First(&emp, emp.ID)

	c.JSON(http.StatusCreated, emp)
}

func GetAllEmployees(c *gin.Context) {
	var employees []models.Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data karyawan"})
		return
	}
	config.DB.Preload("Department").Find(&employees)
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	config.DB.Preload("Department").First(&emp, id)

	c.JSON(http.StatusOK, emp)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	if input.EmployeeID != emp.EmployeeID {
		var existing models.Employee
		if err := config.DB.Where("employee_id = ?", input.EmployeeID).First(&existing).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID karyawan sudah digunakan"})
			return
		}
	}

	var dept models.Department
	if err := config.DB.First(&dept, input.DepartmentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID departemen tidak valid"})
		return
	}

	emp.EmployeeID = input.EmployeeID
	emp.DepartmentID = input.DepartmentID
	emp.Name = input.Name
	emp.Address = input.Address

	if err := config.DB.Save(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data karyawan"})
		return
	}

	config.DB.Preload("Department").First(&emp, emp.ID)

	c.JSON(http.StatusOK, emp)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus karyawan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Karyawan berhasil dihapus"})
}
