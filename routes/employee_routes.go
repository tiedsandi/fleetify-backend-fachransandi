package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/controllers"
)

func EmployeeRoutes(router *gin.Engine) {
	router.POST("/employees", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetAllEmployees)
	router.GET("/employees/:id", controllers.GetEmployeeByID)
	router.PUT("/employees/:id", controllers.UpdateEmployee)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)
}
