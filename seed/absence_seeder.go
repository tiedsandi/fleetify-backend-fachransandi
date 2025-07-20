package seeders

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/helpers"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func SeedAbsences() error {
	var employees []models.Employee
	if err := config.DB.Preload("Department").Find(&employees).Error; err != nil {
		return fmt.Errorf("gagal mengambil data karyawan: %v", err)
	}

	for _, emp := range employees {
		usedDates := map[string]bool{}

		for dayOffset := 0; dayOffset < 7; dayOffset++ {
			date := time.Date(2025, 7, 13+dayOffset, 0, 0, 0, 0, time.Local)
			dateKey := date.Format("2006-01-02")
			if usedDates[dateKey] {
				continue
			}
			usedDates[dateKey] = true

			clockIn := randomTimeForDate(date, 7, 11)
			clockOut := randomTimeForDate(date, 14, 20)

			attendanceID, err := helpers.GenerateAttendanceID(emp.EmployeeID)
			if err != nil {
				fmt.Printf("❌ Gagal generate ID absensi untuk %s: %v\n", emp.EmployeeID, err)
				continue
			}

			attendance := models.Attendance{
				EmployeeID:   emp.EmployeeID,
				AttendanceID: attendanceID,
				ClockIn:      clockIn,
				ClockOut:     clockOut,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			if err := config.DB.Create(&attendance).Error; err != nil {
				fmt.Printf("❌ Gagal insert attendance untuk %s: %v\n", emp.EmployeeID, err)
				continue
			}

			historyIn := models.AttendanceHistory{
				EmployeeID:     emp.EmployeeID,
				AttendanceID:   attendanceID,
				DateAttendance: clockIn,
				AttendanceType: 1,
				Description:    helpers.DescriptionClockIn(emp.EmployeeID),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			historyOut := models.AttendanceHistory{
				EmployeeID:     emp.EmployeeID,
				AttendanceID:   attendanceID,
				DateAttendance: clockOut,
				AttendanceType: 2,
				Description:    helpers.DescriptionClockOut(emp.EmployeeID),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			_ = config.DB.Create(&historyIn)
			_ = config.DB.Create(&historyOut)
		}
	}

	fmt.Println("✅ Seeder absensi selesai untuk 13–19 Juli 2025")
	return nil
}

func randomTimeForDate(date time.Time, startHour, endHour int) time.Time {
	hour := rand.Intn(endHour-startHour+1) + startHour
	minute := rand.Intn(60)
	second := rand.Intn(60)

	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, second, 0, time.Local)
}
