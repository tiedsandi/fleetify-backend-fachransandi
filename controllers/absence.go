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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid: " + err.Error()})
		return
	}

	_, err := helpers.GetEmployeeByID(req.EmployeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	if _, err := helpers.HasClockedInToday(req.EmployeeID); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah melakukan absensi masuk hari ini"})
		return
	}

	attendanceID, _ := helpers.GenerateAttendanceID(req.EmployeeID)
	now := time.Now()

	attendance := models.Attendance{
		EmployeeID:   req.EmployeeID,
		AttendanceID: attendanceID,
		ClockIn:      now,
	}
	config.DB.Create(&attendance)

	desc := helpers.DescriptionClockIn(req.EmployeeID)

	history := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendanceID,
		DateAttendance: now,
		AttendanceType: 1,
		Description:    desc,
	}
	config.DB.Create(&history)

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Absensi masuk berhasil",
		"attendance_id": attendanceID,
		"clock_in_time": now,
		"employee_id":   req.EmployeeID,
	})
}

func UpdateAttendanceOut(c *gin.Context) {
	var req AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid: " + err.Error()})
		return
	}

	if _, err := helpers.GetEmployeeByID(req.EmployeeID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	attendance, err := helpers.HasClockedInToday(req.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Belum ada absensi masuk untuk hari ini"})
		return
	}

	if !attendance.ClockOut.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah melakukan absensi keluar hari ini"})
		return
	}

	now := time.Now()
	attendance.ClockOut = now
	if err := config.DB.Save(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan absensi keluar"})
		return
	}

	desc := helpers.DescriptionClockOut(req.EmployeeID)

	history := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: now,
		AttendanceType: 2,
		Description:    desc,
	}
	config.DB.Create(&history)

	c.JSON(http.StatusOK, gin.H{
		"message":        "Absensi keluar berhasil",
		"attendance_id":  attendance.AttendanceID,
		"clock_out_time": now,
		"employee_id":    req.EmployeeID,
	})
}

func GetAttendanceLogs(c *gin.Context) {
	dateFilter, err := helpers.ParseDateQueryParam(c, "tanggal")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal tidak valid. Gunakan format YYYY-MM-DD"})
		return
	}
	deptID := c.Query("department_id")

	var logs []struct {
		AttendanceHistoryID uint      `json:"id"`
		EmployeeID          string    `json:"employee_id"`
		Name                string    `json:"name"`
		DepartmentName      string    `json:"department"`
		DateAttendance      time.Time `json:"date_attendance"`
		AttendanceType      string    `json:"type"`
		Description         string    `json:"description"`
	}

	query := config.DB.
		Table("attendance_history").
		Select(`
			attendance_history.id as attendance_history_id,
			employees.employee_id,
			employees.name,
			departments.department_name,
			attendance_history.date_attendance,
			CASE attendance_history.attendance_type
				WHEN 1 THEN 'Masuk'
				WHEN 2 THEN 'Keluar'
				ELSE 'Tidak diketahui'
			END as attendance_type,
			attendance_history.description`).
		Joins("JOIN employees ON employees.employee_id = attendance_history.employee_id").
		Joins("JOIN departments ON departments.id = employees.department_id")

	if !dateFilter.IsZero() {
		query = query.Where("DATE(attendance_history.date_attendance) = ?", dateFilter.Format("2006-01-02"))
	}
	if deptID != "" {
		query = query.Where("departments.id = ?", deptID)
	}

	if err := query.Order("attendance_history.date_attendance desc").Scan(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil riwayat absensi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": logs})
}
