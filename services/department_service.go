package services

import (
	"errors"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

type DepartmentRequest struct {
	DepartmentName  string `json:"department_name" binding:"required"`
	MaxClockInTime  string `json:"max_clock_in_time" binding:"required"`
	MaxClockOutTime string `json:"max_clock_out_time" binding:"required"`
}

func CreateDepartmentService(input DepartmentRequest) (*models.Department, error) {
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
	var dept models.Department
	if err := config.DB.First(&dept, id).Error; err != nil {
		return errors.New("departemen tidak ditemukan")
	}
	if err := config.DB.Delete(&dept).Error; err != nil {
		return errors.New("gagal menghapus departemen")
	}
	return nil
}
