package services

import (
	"errors"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/helpers"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

type DepartmentRequest struct {
	DepartmentName  string `json:"department_name" binding:"required"`
	MaxClockInTime  string `json:"max_clock_in_time" binding:"required"`
	MaxClockOutTime string `json:"max_clock_out_time" binding:"required"`
}

func CreateDepartmentService(input DepartmentRequest) (*models.Department, error) {
	if err := helpers.ValidateClockTimes(input.MaxClockInTime, input.MaxClockOutTime); err != nil {
		return nil, err
	}

	dept := models.Department{
		DepartmentName:  input.DepartmentName,
		MaxClockInTime:  input.MaxClockInTime,
		MaxClockOutTime: input.MaxClockOutTime,
	}
	if err := config.DB.Create(&dept).Error; err != nil {
		return nil, errors.New("gagal menyimpan data departemen")
	}
	return &dept, nil
}

func GetDepartmentsService() ([]models.Department, error) {
	var depts []models.Department
	if err := config.DB.Find(&depts).Error; err != nil {
		return nil, errors.New("gagal mengambil data departemen")
	}
	return depts, nil
}

func GetDepartmentByIDService(id string) (*models.Department, error) {
	var dept models.Department
	if err := config.DB.First(&dept, id).Error; err != nil {
		return nil, errors.New("departemen tidak ditemukan")
	}
	return &dept, nil
}

func UpdateDepartmentService(id string, input DepartmentRequest) (*models.Department, error) {
	if err := helpers.ValidateClockTimes(input.MaxClockInTime, input.MaxClockOutTime); err != nil {
		return nil, err
	}

	var dept models.Department
	if err := config.DB.First(&dept, id).Error; err != nil {
		return nil, errors.New("departemen tidak ditemukan")
	}

	dept.DepartmentName = input.DepartmentName
	dept.MaxClockInTime = input.MaxClockInTime
	dept.MaxClockOutTime = input.MaxClockOutTime

	if err := config.DB.Save(&dept).Error; err != nil {
		return nil, errors.New("gagal memperbarui departemen")
	}

	return &dept, nil
}

func DeleteDepartmentService(id string) error {
	tx := config.DB.Begin()

	var dept models.Department
	if err := tx.First(&dept, id).Error; err != nil {
		tx.Rollback()
		return errors.New("departemen tidak ditemukan")
	}

	var employees []models.Employee
	if err := tx.Where("department_id = ?", id).Find(&employees).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, emp := range employees {
		if err := tx.Where("employee_id = ?", emp.EmployeeID).Delete(&models.AttendanceHistory{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("employee_id = ?", emp.EmployeeID).Delete(&models.Attendance{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Where("department_id = ?", id).Delete(&models.Employee{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&dept).Error; err != nil {
		tx.Rollback()
		return errors.New("gagal menghapus departemen")
	}

	return tx.Commit().Error
}
