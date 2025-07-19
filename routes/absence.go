package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/controllers"
)

func AbsenceRoutes(router *gin.Engine) {
	router.POST("/absence", controllers.CreateAttendanceIn)
	router.PUT("/absence", controllers.UpdateAttendanceOut)
	router.GET("/absence", controllers.GetAttendanceLogs)
}
