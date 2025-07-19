package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func GenerateAttendanceID(employeeID string) (string, error) {
	var emp models.Employee
	err := config.DB.Preload("Department").Where("employee_id = ?", employeeID).First(&emp).Error
	if err != nil {
		return "", fmt.Errorf("employee not found")
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

	// 	var code string
	// switch strings.ToLower(emp.Department.DepartmentName) {
	// case "hr", "human resource":
	// 	code = "HR"
	// case "developer":
	// 	code = "DEV"
	// case "marketing":
	// 	code = "MKT"
	// case "finance":
	// 	code = "FIN"
	// default:
	// 	code = "GEN"
	// }

	today := time.Now().Format("20250720")

	var last models.Attendance
	_ = config.DB.Select("id").Order("id desc").First(&last).Error

	return fmt.Sprintf("%s-%s-%d", code, today, last.ID+1), nil
}
