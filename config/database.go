package config

import (
	"log"

	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/fleetify?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}

	DB = db
	log.Println("✅ Koneksi ke database berhasil")
}

func Migration() {
	DB.AutoMigrate(
		&models.Department{},
		&models.Employee{},
		&models.Attendance{},
		&models.AttendanceHistory{},
	)

	DB.Exec(`
		ALTER TABLE attendance
		ADD CONSTRAINT fk_attendance_employee
		FOREIGN KEY (employee_id) REFERENCES employees(employee_id)
		 ON DELETE RESTRICT ON UPDATE RESTRICT
	`)
	DB.Exec(`
		ALTER TABLE attendance_history
		ADD CONSTRAINT fk_attendance_history_attendance
		FOREIGN KEY (attendance_id) REFERENCES attendance(attendance_id)
		 ON DELETE RESTRICT ON UPDATE RESTRICT
	`)
	DB.Exec(`
		ALTER TABLE attendance_history
		ADD CONSTRAINT fk_attendance_history_employee
		FOREIGN KEY (employee_id) REFERENCES employees(employee_id)
		 ON DELETE RESTRICT ON UPDATE RESTRICT
	`)

}

func ResetDB() {
	DB.Migrator().DropTable(
		&models.AttendanceHistory{},
		&models.Attendance{},
		&models.Employee{},
		&models.Department{},
	)
}
