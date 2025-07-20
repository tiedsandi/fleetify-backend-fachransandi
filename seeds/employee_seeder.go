package seeds

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

func SeedEmployees() {
	rand.Seed(time.Now().UnixNano())

	names := []string{
		"Fachran Sandi", "John Doe", "Jane Doe", "Fachran Doe", "John Sandi",
		"Jane Sandi", "Sandi Doe", "Doe Fachran", "Sandi Fachran", "John Fachran",
	}

	locations := []string{
		"Cilangkap Jakarta", "Setiabudi Jakarta", "Cibaduyut Bandung", "Klojen Malang",
		"Sukajadi Bandung", "Ujungberung Bandung", "Gubeng Surabaya", "Mulyorejo Surabaya",
		"Panakkukang Makassar", "Medan Kota Medan",
	}

	var departments []models.Department
	if err := config.DB.Find(&departments).Error; err != nil {
		fmt.Println("Gagal mengambil departemen:", err)
		return
	}

	if len(departments) == 0 {
		fmt.Println("Seeder gagal: tidak ada departemen ditemukan.")
		return
	}

	codeMap := map[string]string{
		"hr":          "HR",
		"developer":   "DEV",
		"engineering": "ENG",
		"marketing":   "MKT",
		"finance":     "FIN",
	}
	counterMap := map[string]int{}

	for i := 0; i < 10; i++ {
		name := names[i%len(names)]
		address := fmt.Sprintf("no.%d %s", rand.Intn(50)+1, locations[rand.Intn(len(locations))])

		dept := departments[rand.Intn(len(departments))]

		prefix := codeMap[strings.ToLower(dept.DepartmentName)]
		counterMap[prefix]++
		empID := fmt.Sprintf("%s-%03d", prefix, counterMap[prefix])

		employee := models.Employee{
			EmployeeID:   empID,
			Name:         name,
			Address:      address,
			DepartmentID: dept.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := config.DB.Create(&employee).Error; err != nil {
			fmt.Printf("❌ Gagal membuat employee %s: %v\n", name, err)
		} else {
			fmt.Printf("✅ Employee created: %s (%s)\n", name, empID)
		}
	}
}
