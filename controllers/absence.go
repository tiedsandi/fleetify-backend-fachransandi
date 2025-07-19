package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/helpers"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

type AttendanceRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

func CreateAttendanceIn(c *gin.Context) {
	var req AttendanceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var emp models.Employee
	if err := config.DB.Where("employee_id = ?", req.EmployeeID).First(&emp).Error; err != nil {
		c.JSON(404, gin.H{"error": "Employee not found"})
		return
	}

	today := time.Now()
	start := today.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	var existing models.Attendance
	err := config.DB.
		Where("employee_id = ? AND clock_in >= ? AND clock_in < ?", req.EmployeeID, start, end).
		First(&existing).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have already clocked in today"})
		return
	}

	EmployeeId, _ := helpers.GenerateAttendanceID(req.EmployeeID)

	attendance := models.Attendance{
		EmployeeID:   req.EmployeeID,
		AttendanceID: EmployeeId,
		ClockIn:      time.Now(),
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create attendance"})
		return
	}

	history := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   EmployeeId,
		DateAttendance: time.Now(),
		AttendanceType: 1, // 1 = clock-in
		Description:    "Clock in",
	}

	if err := config.DB.Create(&history).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create attendance history"})
		return
	}

	c.JSON(201, gin.H{
		"message":       "Clock in successful",
		"attendance_id": EmployeeId,
		"clock_in_time": attendance.ClockIn,
		"employee_id":   req.EmployeeID,
	})

}

func UpdateAttendanceOut(c *gin.Context) {}

func GetAttendanceLogs(c *gin.Context) {}
