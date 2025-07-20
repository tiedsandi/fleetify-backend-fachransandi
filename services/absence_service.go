package services

import (
	"errors"
	"time"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/helpers"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

type AttendanceResult struct {
	EmployeeID   string
	AttendanceID string
	Timestamp    time.Time
}

type AttendanceLog struct {
	ID         uint   `json:"id"`
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Date       string `json:"date"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}

func CreateClockInService(employeeID string) (*AttendanceResult, error) {

	_, err := helpers.GetEmployeeByID(employeeID)
	if err != nil {
		return nil, errors.New("karyawan tidak ditemukan")
	}

	if _, err := helpers.HasClockedInToday(employeeID); err == nil {
		return nil, errors.New("anda sudah melakukan absensi masuk hari ini")
	}

	attendanceID, _ := helpers.GenerateAttendanceID(employeeID)
	now := time.Now()

	attendance := models.Attendance{
		EmployeeID:   employeeID,
		AttendanceID: attendanceID,
		ClockIn:      now,
	}
	config.DB.Create(&attendance)

	desc := helpers.DescriptionClockIn(employeeID)
	history := models.AttendanceHistory{
		EmployeeID:     employeeID,
		AttendanceID:   attendanceID,
		DateAttendance: now,
		AttendanceType: 1,
		Description:    desc,
	}
	config.DB.Create(&history)

	return &AttendanceResult{
		EmployeeID:   employeeID,
		AttendanceID: attendanceID,
		Timestamp:    now,
	}, nil
}

func UpdateClockOutService(employeeID string) (*AttendanceResult, error) {
	_, err := helpers.GetEmployeeByID(employeeID)
	if err != nil {
		return nil, errors.New("karyawan tidak ditemukan")
	}

	attendance, err := helpers.HasClockedInToday(employeeID)
	if err != nil {
		return nil, errors.New("belum ada absensi masuk untuk hari ini")
	}

	if !attendance.ClockOut.IsZero() {
		return nil, errors.New("anda sudah melakukan absensi keluar hari ini")
	}

	now := time.Now()
	attendance.ClockOut = now

	if err := config.DB.Save(&attendance).Error; err != nil {
		return nil, errors.New("gagal menyimpan absensi keluar")
	}

	desc := helpers.DescriptionClockOut(employeeID)
	history := models.AttendanceHistory{
		EmployeeID:     employeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: now,
		AttendanceType: 2,
		Description:    desc,
	}
	config.DB.Create(&history)

	return &AttendanceResult{
		EmployeeID:   employeeID,
		AttendanceID: attendance.AttendanceID,
		Timestamp:    now,
	}, nil
}

func GetAttendanceLogsService(dateFilter time.Time, departmentID string) ([]AttendanceLog, error) {
	var tempLogs []struct {
		ID             uint
		EmployeeID     string
		Name           string
		Department     string
		DateAttendance time.Time
		Type           string
		Status         string
	}

	db := config.DB.Table("attendance_history").
		Select(`
			attendance_history.id,
			employees.employee_id,
			employees.name,
			departments.department_name AS department,
			attendance_history.date_attendance,
			CASE attendance_history.attendance_type
				WHEN 1 THEN 'Masuk'
				WHEN 2 THEN 'Keluar'
				ELSE 'Tidak diketahui'
			END AS type,
			attendance_history.description AS status
		`).
		Joins("JOIN employees ON employees.employee_id = attendance_history.employee_id").
		Joins("JOIN departments ON departments.id = employees.department_id")

	if !dateFilter.IsZero() {
		db = db.Where("DATE(attendance_history.date_attendance) = ?", dateFilter.Format("2006-01-02"))
	}

	if departmentID != "" {
		db = db.Where("departments.id = ?", departmentID)
	}

	if err := db.Order("attendance_history.date_attendance DESC").Scan(&tempLogs).Error; err != nil {
		return nil, err
	}

	var logs []AttendanceLog
	for _, item := range tempLogs {
		logs = append(logs, AttendanceLog{
			ID:         item.ID,
			EmployeeID: item.EmployeeID,
			Name:       item.Name,
			Department: item.Department,
			Date:       item.DateAttendance.Format("2006-01-02 15:04:05"),
			Type:       item.Type,
			Status:     item.Status,
		})
	}

	return logs, nil
}
