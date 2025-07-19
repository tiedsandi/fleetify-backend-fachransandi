package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/controllers"
)

func DepartmentRoutes(router *gin.Engine) {
	router.POST("/departments", controllers.CreateDepartment)
	router.GET("/departments", controllers.GetDepartments)
	router.GET("/departments/:id", controllers.GetDepartmentByID)
	router.PUT("/departments/:id", controllers.UpdateDepartment)
	router.DELETE("/departments/:id", controllers.DeleteDepartment)
}
