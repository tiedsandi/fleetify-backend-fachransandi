package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/services"
)

func CreateDepartment(c *gin.Context) {
	var input services.DepartmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	dept, err := services.CreateDepartmentService(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dept)
}

func GetDepartments(c *gin.Context) {
	depts, err := services.GetDepartmentsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, depts)
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")
	dept, err := services.GetDepartmentByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dept)
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var input services.DepartmentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	dept, err := services.UpdateDepartmentService(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dept)
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteDepartmentService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Departemen berhasil dihapus"})
}
