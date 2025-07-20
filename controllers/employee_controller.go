package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/services"
)

func CreateEmployee(c *gin.Context) {
	var input services.EmployeeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	employee, err := services.CreateEmployeeService(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func GetAllEmployees(c *gin.Context) {
	employees, err := services.GetAllEmployeesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := services.GetEmployeeByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var input services.EmployeeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid: " + err.Error()})
		return
	}

	employee, err := services.UpdateEmployeeService(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteEmployeeService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Karyawan berhasil dihapus"})
}
