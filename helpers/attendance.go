package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
	"github.com/tiedsandi/fleetify-backend-fachransandi/utils"
)

func GenerateAttendanceID(employeeID string) (string, error) {
	var emp models.Employee
	err := config.DB.Preload("Department").Where("employee_id = ?", employeeID).First(&emp).Error
	if err != nil {
		return "", fmt.Errorf("karyawan tidak ditemukan")
	}

	if emp.Department.ID == 0 {
		return "", fmt.Errorf("karyawan tidak memiliki departemen yang valid")
	}

	deptName := strings.ToLower(emp.Department.DepartmentName)
	codeMap := map[string]string{
		"hr":          "HR",
		"developer":   "DEV",
		"engineering": "ENG",
		"marketing":   "MKT",
		"finance":     "FIN",
	}
	code, ok := codeMap[deptName]
	if !ok {
		code = "GEN"
	}

	today := time.Now().Format("20060102")

	var last models.Attendance
	_ = config.DB.Select("id").Order("id desc").First(&last).Error

	return fmt.Sprintf("%s-%s-%d", code, today, last.ID+1), nil
}

func DescriptionClockIn(employeeID string) string {
	var employee models.Employee
	if err := config.DB.Preload("Department").Where("employee_id = ?", employeeID).First(&employee).Error; err != nil {
		return "Karyawan atau departemen tidak ditemukan"
	}

	maxClockIn, err := time.Parse("15:04:05", employee.Department.MaxClockInTime)
	if err != nil {
		return "Format jam masuk departemen tidak valid"
	}

	now := time.Now()
	maxTimeToday := time.Date(now.Year(), now.Month(), now.Day(),
		maxClockIn.Hour(), maxClockIn.Minute(), maxClockIn.Second(), 0, now.Location())

	if now.Before(maxTimeToday) || now.Equal(maxTimeToday) {
		return "Tepat waktu"
	}

	lateDuration := now.Sub(maxTimeToday)
	return fmt.Sprintf("Terlambat %s", utils.FormatDuration(lateDuration))
}

func DescriptionClockOut(employeeID string) string {
	var employee models.Employee
	if err := config.DB.Preload("Department").Where("employee_id = ?", employeeID).First(&employee).Error; err != nil {
		return "Karyawan atau departemen tidak ditemukan"
	}

	maxClockOut, err := time.Parse("15:04:05", employee.Department.MaxClockOutTime)
	if err != nil {
		return "Format jam pulang departemen tidak valid"
	}

	now := time.Now()
	maxTimeToday := time.Date(now.Year(), now.Month(), now.Day(),
		maxClockOut.Hour(), maxClockOut.Minute(), maxClockOut.Second(), 0, now.Location())

	if now.Before(maxTimeToday) {
		return "Pulang lebih awal"
	}

	overtime := now.Sub(maxTimeToday)
	return fmt.Sprintf("Lembur %s", utils.FormatDuration(overtime))
}

func GetEmployeeByID(employeeID string) (models.Employee, error) {
	var emp models.Employee
	err := config.DB.Where("employee_id = ?", employeeID).First(&emp).Error
	return emp, err
}

func HasClockedInToday(employeeID string) (models.Attendance, error) {
	today := time.Now()
	start := today.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	var attendance models.Attendance
	err := config.DB.
		Where("employee_id = ? AND clock_in >= ? AND clock_in < ?", employeeID, start, end).
		First(&attendance).Error

	return attendance, err
}

func ParseDateQueryParam(c *gin.Context, param string) (time.Time, error) {
	value := c.Query(param)
	if value == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", value)
}
