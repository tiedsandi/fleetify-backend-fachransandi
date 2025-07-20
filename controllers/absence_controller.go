package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/helpers"
	"github.com/tiedsandi/fleetify-backend-fachransandi/services"
)

type AttendanceRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

func CreateAttendanceIn(c *gin.Context) {
	var req AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid: " + err.Error()})
		return
	}

	result, err := services.CreateClockInService(req.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Absensi masuk berhasil",
		"attendance_id": result.AttendanceID,
		"clock_in_time": result.Timestamp,
		"employee_id":   result.EmployeeID,
	})
}

func UpdateAttendanceOut(c *gin.Context) {
	var req AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid: " + err.Error()})
		return
	}

	result, err := services.UpdateClockOutService(req.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Absensi keluar berhasil",
		"attendance_id":  result.AttendanceID,
		"clock_out_time": result.Timestamp,
		"employee_id":    result.EmployeeID,
	})
}

func GetAttendanceLogs(c *gin.Context) {
	dateFilter, err := helpers.ParseDateQueryParam(c, "tanggal")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format tanggal tidak valid. Gunakan format YYYY-MM-DD",
		})
		return
	}

	departmentID := c.Query("department_id")

	logs, err := services.GetAttendanceLogsService(dateFilter, departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data absensi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": logs})
}
