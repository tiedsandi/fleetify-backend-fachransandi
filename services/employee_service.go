package services

import (
	"errors"

	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/models"
)

type EmployeeRequest struct {
	EmployeeID   string `json:"employee_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	DepartmentID uint   `json:"department_id" binding:"required"`
}

func CreateEmployeeService(input EmployeeRequest) (*models.Employee, error) {
	var dept models.Department
	if err := config.DB.First(&dept, input.DepartmentID).Error; err != nil {
		return nil, errors.New("ID departemen tidak valid")
	}

	var existing models.Employee
	if err := config.DB.Where("employee_id = ?", input.EmployeeID).First(&existing).Error; err == nil {
		return nil, errors.New("ID karyawan sudah digunakan")
	}

	employee := models.Employee{
		EmployeeID:   input.EmployeeID,
		Name:         input.Name,
		Address:      input.Address,
		DepartmentID: input.DepartmentID,
	}

	if err := config.DB.Create(&employee).Error; err != nil {
		return nil, errors.New("gagal menyimpan data karyawan")
	}

	config.DB.Preload("Department").First(&employee, employee.ID)
	return &employee, nil
}

func GetAllEmployeesService() ([]models.Employee, error) {
	var employees []models.Employee
	if err := config.DB.Preload("Department").Find(&employees).Error; err != nil {
		return nil, errors.New("gagal mengambil data karyawan")
	}
	return employees, nil
}

func GetEmployeeByIDService(id string) (*models.Employee, error) {
	var employee models.Employee
	if err := config.DB.Preload("Department").First(&employee, id).Error; err != nil {
		return nil, errors.New("karyawan tidak ditemukan")
	}
	return &employee, nil
}

func UpdateEmployeeService(id string, input EmployeeRequest) (*models.Employee, error) {
	var employee models.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		return nil, errors.New("karyawan tidak ditemukan")
	}

	if input.EmployeeID != employee.EmployeeID {
		var existing models.Employee
		if err := config.DB.Where("employee_id = ?", input.EmployeeID).First(&existing).Error; err == nil {
			return nil, errors.New("ID karyawan sudah digunakan")
		}
	}

	var dept models.Department
	if err := config.DB.First(&dept, input.DepartmentID).Error; err != nil {
		return nil, errors.New("ID departemen tidak valid")
	}

	employee.EmployeeID = input.EmployeeID
	employee.Name = input.Name
	employee.Address = input.Address
	employee.DepartmentID = input.DepartmentID

	if err := config.DB.Save(&employee).Error; err != nil {
		return nil, errors.New("gagal memperbarui data karyawan")
	}

	config.DB.Preload("Department").First(&employee, employee.ID)
	return &employee, nil
}

func DeleteEmployeeService(id string) error {
	var employee models.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		return errors.New("karyawan tidak ditemukan")
	}

	if err := config.DB.Delete(&employee).Error; err != nil {
		return errors.New("gagal menghapus karyawan")
	}

	return nil
}
