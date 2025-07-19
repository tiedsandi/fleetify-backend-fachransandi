package models

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	EmployeeID   string     `gorm:"type:varchar(50);unique;not null" json:"employee_id"`
	DepartmentID uint       `gorm:"not null" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	Address      string     `gorm:"type:text;not null" json:"address"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
