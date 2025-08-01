package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	DepartmentRoutes(router)
	EmployeeRoutes(router)
	AbsenceRoutes(router)
}
