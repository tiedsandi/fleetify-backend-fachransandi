package seeds

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func SeedDepartments() {
	codeMap := map[string]string{
		"hr":          "HR",
		"developer":   "DEV",
		"engineering": "ENG",
		"marketing":   "MKT",
		"finance":     "FIN",
	}

	rand.Seed(time.Now().UnixNano())

	for name := range codeMap {
		clockIn := randomTimeString(7, 10)
		clockOut := randomTimeString(15, 18)

		dept := models.Department{
			DepartmentName:  name,
			MaxClockInTime:  clockIn,
			MaxClockOutTime: clockOut,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if err := config.DB.Create(&dept).Error; err != nil {
			fmt.Printf("Failed to seed department %s: %v\n", name, err)
		} else {
			fmt.Printf("Seeded department: %s (%s - %s)\n", name, clockIn, clockOut)
		}
	}
}

func randomTimeString(startHour, endHour int) string {
	hour := rand.Intn(endHour-startHour+1) + startHour
	minute := []int{0, 30}[rand.Intn(2)]
	return fmt.Sprintf("%02d:%02d:00.000", hour, minute)
}
